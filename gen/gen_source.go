package gen

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"go/format"

	"github.com/dave/jennifer/jen"
	klog "github.com/go-kit/kit/log"
	akstrings "github.com/mattmunz/designlanguage/appkit/strings"

	"github.com/mattmunz/designlanguage/appkit/misc"
	"github.com/mattmunz/designlanguage/model"
)

var goTypeByDLType = map[string]string{
	"String": "string",
	"Int":    "int",
	// TODO This type is platform-specific and should be abstracted further if possible.
	"CobraCommand": "*cobra.Command",
}

func RenderDesignSummary(design model.Design) string {
	return fmt.Sprintf("Namespace: %s, Components: %d, Entities: %d, Objects: %d", design.Namespace(), len(design.Components()), len(design.Entities()), len(design.Objects()))
}

func RenderDesignSource(design model.Design) (string, error) {
	packageName, err := getPackageName(design.Namespace())
	if err != nil {
		return "", err
	}

	outFile := jen.NewFile(packageName)
	stmt := outFile.Empty()

	for _, component := range design.Components() {
		if slices.Contains(entitiesToComponents(design.Entities()), component) ||
			slices.Contains(objectsToComponents(design.Objects()), component) {
			continue
		}

		stmt = addComponent(stmt, component)
	}

	for _, entity := range design.Entities() {
		stmt = addEntity(stmt, entity)
	}

	for _, obj := range design.Objects() {
		stmt = addObject(stmt, obj)
	}

	return renderFile(outFile)
}

// TODO There must be a nicer way to do this.
func entitiesToComponents(entity []model.Entity) []model.Component {
	components := make([]model.Component, len(entity))
	for i, e := range entity {
		components[i] = e
	}
	return components
}

// TODO There must be a nicer way to do this.
func objectsToComponents(obj []model.Object) []model.Component {
	components := make([]model.Component, len(obj))
	for i, o := range obj {
		components[i] = o
	}
	return components
}

func getPackageName(namespace string) (string, error) {
	tokens := strings.Split(namespace, "/")

	if len(tokens) < 1 {
		return "", errors.New("Invalid namespace: " + namespace)
	}

	return tokens[len(tokens)-1], nil
}

func RenderComponentSource(component model.Component) (string, error) {
	stmt := jen.Empty()
	addComponent(stmt, component)
	return render(stmt)
}

func addComponent(stmt *jen.Statement, component model.Component) *jen.Statement {
	return stmt.Type().Id(component.Name()).Interface().Line().Line()
}

func RenderEntitySource(entity model.Entity) (string, error) {
	stmt := jen.Empty()
	addEntity(stmt, entity)
	return render(stmt)
}

func addEntity(stmt *jen.Statement, entity model.Entity) *jen.Statement {
	return stmt.Type().Id(entity.Name()).InterfaceFunc(func(g *jen.Group) {
		for _, attr := range entity.Attributes() {
			addAttribute(g, attr)
		}
	}).Line().Line()
}

func addObject(stmt *jen.Statement, obj model.Object) *jen.Statement {
	return stmt.Type().Id(obj.Name()).InterfaceFunc(func(g *jen.Group) {
		for _, attr := range obj.Attributes() {
			addAttribute(g, attr)
		}

		for _, method := range obj.Methods() {
			addMethod(g, method)
		}
	}).Line().Line()
}

func addAttribute(g *jen.Group, attr model.Attribute) {
	g.Id(attr.Name()).Params().Id(getGoTypeName(attr.Type().Name()))
}

func getGoTypeName(dlTypeName string) string {
	if goType, exists := goTypeByDLType[dlTypeName]; exists {
		return goType
	}

	return dlTypeName
}

func RenderObjectSource(obj model.Object) (string, error) {
	stmt := jen.Empty()
	addObject(stmt, obj)
	return render(stmt)
}

func addMethod(g *jen.Group, method model.Method) error {
	next := g.Id(method.Name()).ParamsFunc(func(g *jen.Group) {
		for _, param := range method.Params() {
			addParamToGroup(g, param)
		}
	})

	if len(method.ReturnVals()) < 1 {
		return errors.New("No return val specified for method")
	}

	next.ParamsFunc(func(g *jen.Group) {
		for _, param := range method.ReturnVals() {
			addParamToGroup(g, param)
		}
	})

	return nil
}

func addParamToGroup(g *jen.Group, param model.Param) {
	g.Id(akstrings.LowercaseFirst(param.Name())).Id(getGoTypeName(param.Type().Name()))
}

func renderFile(f *jen.File) (string, error) {
	src, err := format.Source([]byte(f.GoString()))
	return string(src), err
}

func render(statement *jen.Statement) (string, error) {
	src, err := format.Source([]byte(strings.TrimSpace(statement.GoString())))
	return string(src), err
}

type ComponentImpl struct {
	name string
}

func (c *ComponentImpl) Name() string {
	return c.name
}

func NewComponent(name string) model.Component {
	return &ComponentImpl{name}
}

func GenerateSourceForDL(projectDir string, logger klog.Logger, dryRun bool) error {
	projectDirPath, err := filepath.Abs(projectDir)
	if err != nil {
		return err
	}
	designPath := filepath.Join(projectDirPath, "documentation", "design")

	designDirInfo, err := os.Stat(designPath)
	if err != nil {
		return err
	}

	if !designDirInfo.IsDir() {
		return fmt.Errorf("File is not dir: %s", designPath)
	}

	parsedDesigns := NewDesignList()

	misc.LogMessage(logger, fmt.Sprintf("Walking the path %q", designPath))

	err = filepath.Walk(designPath, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".nzsd.txt") {
			return nil
		}

		return HandleDLMFile(logger, parsedDesigns, designPath, path, info, dryRun, projectDirPath, err)
	})

	if err != nil {

		return fmt.Errorf("Error walking the path %q: %w", designPath, err)
	}

	misc.LogMessage(logger, "Generating code...")

	for _, design := range parsedDesigns.All() {
		misc.LogMessage(logger, fmt.Sprintf("Design summary:\n%s", RenderDesignSummary(design)))

		designSource, err := RenderDesignSource(design)

		if err != nil {
			misc.LogMessage(logger, fmt.Sprintf("Error rendering design source: %v", err))
			continue
		}

		misc.LogMessage(logger, fmt.Sprintf("Design source: %s", designSource))

		if dryRun {
			continue
		}

		dirPath := filepath.Join(projectDirPath, "model", "gen", design.Namespace())
		filePath := filepath.Join(dirPath, fmt.Sprintf("%s.go", design.Namespace()))

		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}

		if err = os.WriteFile(filePath, []byte(designSource), 0644); err != nil {
			return err
		}

		misc.LogMessage(logger, fmt.Sprintf("Wrote file: %s", filePath))
	}
	return nil
}
