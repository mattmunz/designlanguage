package dl

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dave/jennifer/jen"

	"go/format"

	"github.com/mattmunz/designlanguage/model/designlanguage"
)

func RenderDesignSummary(design designlanguage.Design) string {
	return fmt.Sprintf("Namespace: %s, Components: %d, Entities: %d, Objects: %d", design.Namespace(), len(design.Components()), len(design.Entities()), len(design.Objects()))
}

func RenderDesignSource(design designlanguage.Design) (string, error) {
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
func entitiesToComponents(entity []designlanguage.Entity) []designlanguage.Component {
	components := make([]designlanguage.Component, len(entity))
	for i, e := range entity {
		components[i] = e
	}
	return components
}

// TODO There must be a nicer way to do this.
func objectsToComponents(obj []designlanguage.Object) []designlanguage.Component {
	components := make([]designlanguage.Component, len(obj))
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

func RenderComponentSource(component designlanguage.Component) (string, error) {
	stmt := jen.Empty()
	addComponent(stmt, component)
	return render(stmt)
}

func addComponent(stmt *jen.Statement, component designlanguage.Component) *jen.Statement {
	return stmt.Type().Id(component.Name()).Interface().Line().Line()
}

func RenderEntitySource(entity designlanguage.Entity) (string, error) {
	stmt := jen.Empty()
	addEntity(stmt, entity)
	return render(stmt)
}

func addEntity(stmt *jen.Statement, entity designlanguage.Entity) *jen.Statement {
	return stmt.Type().Id(entity.Name()).InterfaceFunc(func(g *jen.Group) {
		for _, attr := range entity.Attributes() {
			addAttribute(g, attr)
		}
	}).Line().Line()
}

func addObject(stmt *jen.Statement, obj designlanguage.Object) *jen.Statement {
	return stmt.Type().Id(obj.Name()).InterfaceFunc(func(g *jen.Group) {
		for _, attr := range obj.Attributes() {
			addAttribute(g, attr)
		}

		for _, method := range obj.Methods() {
			addMethod(g, method)
		}
	}).Line().Line()
}

func addAttribute(g *jen.Group, attr designlanguage.Attribute) {
	g.Id(attr.Name()).Params().Id(getGoTypeName(attr.Type().Name()))
}

// TODO replace this with a map.
func getGoTypeName(dlTypeName string) string {
	switch dlTypeName {
	case "String":
		return "string"
	case "Int":
		return "int"
	default:
		return dlTypeName
	}
}

func RenderObjectSource(obj designlanguage.Object) (string, error) {
	stmt := jen.Empty()
	addObject(stmt, obj)
	return render(stmt)
}

func addMethod(g *jen.Group, method designlanguage.Method) error {
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

func addParamToGroup(g *jen.Group, param designlanguage.Param) {
	g.Id(LowercaseFirst(param.Name())).Id(getGoTypeName(param.Type().Name()))
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

func NewComponent(name string) designlanguage.Component {
	return &ComponentImpl{name}
}

// TODO Move to strings lib. Google says this is not in the stdlib.
func LowercaseFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Decode the first rune safely to support UTF-8 characters
	r, size := utf8.DecodeRuneInString(s)

	// Convert the first rune to lowercase and reconstruct the string
	return string(unicode.ToLower(r)) + s[size:]
}
