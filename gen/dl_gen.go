package gen

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	klog "github.com/go-kit/kit/log"

	"github.com/mattmunz/appkit/misc"
	"github.com/mattmunz/designlanguage/model"
)

type DesignList struct {
	designs []model.Design
}

func NewDesignList() *DesignList {
	return &DesignList{[]model.Design{}}
}

func (d *DesignList) Add(design model.Design) {
	d.designs = append(d.designs, design)
}

func (d *DesignList) All() []model.Design {
	return d.designs
}

// Collect all design objects from design files.
// designPath is root dir of designs.
// path is the path to the specific design file being parsed.
func HandleDLMFile(logger klog.Logger, parsedDesigns *DesignList, designPath, path string, info fs.FileInfo, dryRun bool, outputPath string,
	err error) error {
	if err != nil {
		return err
	}

	namespace, _, err := ParseDLMFilePath(designPath, path)
	if err != nil {
		return err
	}

	misc.LogMessage(logger, fmt.Sprintf("Parsing design (%s) / (%s)...", designPath, path))

	design, err := model.Parse(path, namespace)
	if err != nil {
		return err
	}
	parsedDesigns.Add(design)

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
