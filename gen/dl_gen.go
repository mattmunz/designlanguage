package gen

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	klog "github.com/go-kit/kit/log"

	"github.com/mattmunz/designlanguage/appkit/misc"
	"github.com/mattmunz/designlanguage/model/designlanguage"
)

type DesignList struct {
	designs []designlanguage.Design
}

func NewDesignList() *DesignList {
	return &DesignList{[]designlanguage.Design{}}
}

func (d *DesignList) Add(model designlanguage.Design) {
	d.designs = append(d.designs, model)
}

func (d *DesignList) All() []designlanguage.Design {
	return d.designs
}

// Collect all design objects from design files.
// designPath is root dir of designs
// path is the path to the specific design file being parsed.
// TODO Implement this to read in the file to a set of designlanguage.Design objects.
// They should share a namespace defined by their path.
func HandleDLMFile(logger klog.Logger, parsedModels *DesignList, designPath, path string, info fs.FileInfo, dryRun bool, outputPath string,
	err error) error {
	if err != nil {
		return err
	}

	// TODO Filter out files that don't end in ".Design.md"
	namespace, fileName, err := ParseDLMFilePath(designPath, path)
	if err != nil {
		// TODO if err type is wrongfiletype then skip.
		return err
	}

	misc.LogMessage(logger, fmt.Sprintf("TODO Does filename (%s) matter?", fileName))

	// TODO Somewhere past this a panic on read!!!

	misc.LogMessage(logger, fmt.Sprintf("TODO Parsing design (%s) / (%s)...", designPath, path))

	design := designlanguage.Parse(path, namespace)

	misc.LogMessage(logger, "TODO Adding design...")

	parsedModels.Add(design)

	return nil
}

func ParseDLMFilePath(designPath, path string) (namespace string, fileName string, err error) {
	relativePath, err := filepath.Rel(designPath, path)
	if err != nil {
		return "", "", fmt.Errorf("no relative: %s, %s", designPath, path)
	}

	parts := strings.Split(relativePath, string(filepath.Separator))

	if len(parts) < 1 {
		return "", "", fmt.Errorf("invalid path: %s", path)
	}

	if len(parts) == 1 {
		return "", parts[0], nil
	}

	return strings.Join(parts[:len(parts)-1], "/"), parts[len(parts)-1], nil
}
