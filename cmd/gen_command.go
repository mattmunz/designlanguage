package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	klog "github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/mattmunz/designlanguage/appkit/cmd"
	"github.com/mattmunz/designlanguage/appkit/misc"
	"github.com/mattmunz/designlanguage/gen"
	"github.com/mattmunz/designlanguage/model/dl"
)

func newGenCmd(logger klog.Logger) *cobra.Command {
	// TODO Add dry-run flag here. Refactor up from subcommands which already have this flag.
	genCmd := &cobra.Command{
		Use:   "gen",
		Short: "Code generation commands.",
	}

	genCmd.AddCommand(newDLMCmd(logger))

	return genCmd
}

// TODO Refactor this up into gen command.
/* Design Language Model command */
func newDLMCmd(logger klog.Logger) *cobra.Command {
	dlmCmd := &cobra.Command{
		Args:  validateGenDLMArgs,
		Use:   "dlm",
		Short: "Generate model source code from Design Language Model files.",
	}

	// TODO Share these flags with all gen subcommands.
	dryRunFlag := dlmCmd.Flags().BoolP("dryRun", "d", true, "True to read but not make changes on disk, true by default.")
	projectDirFlag := dlmCmd.Flags().StringP("projectDir", "p", "",
		"Path to the project directory.")

	doDLM1 := func(log klog.Logger, cmd *cobra.Command, args []string) error {
		return doDesignLanguageModel(log, cmd, args, *dryRunFlag, *projectDirFlag)
	}

	dlmCmd.Run = cmd.WrapRunner(doDLM1, logger)

	return dlmCmd
}

func validateGenDLMArgs(_ *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.New("No args expected")
	}
	return nil
}

func doDesignLanguageModel(logger klog.Logger, cmd *cobra.Command, args []string, dryRun bool, projectDir string) error {
	err := validateGenDLMArgs(cmd, args)
	if err != nil {
		return err
	}

	// TODO Move all of the following into a new fn in gen package

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
		return errors.Errorf("File is not dir: %s", designPath)
	}

	parsedDesigns := gen.NewDesignList()

	fmt.Printf("TODO walking the path %q\n", designPath)

	err = filepath.Walk(designPath, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".nzsd.txt") {
			return nil
		}

		return gen.HandleDLMFile(logger, parsedDesigns, designPath, path, info, dryRun, projectDirPath, err)
	})

	if err != nil {
		fmt.Printf("TODO error walking the path %q: %v\n", designPath, err)
		return err
	}

	// TODO Generate code based on the parsed models

	misc.LogMessage(logger, "TODO Handling designs...")

	for _, design := range parsedDesigns.All() {
		fmt.Printf("TODO Design: %s, Namespace: %s\n", "?", design.Namespace())

		designSummary := dl.RenderDesignSummary(design)

		// TODO Now walk design and output each component, entity, and object as an interface. For now just print out the design source.

		fmt.Printf("Design summary:\n%s\n", designSummary)

		designSource, err := dl.RenderDesignSource(design)

		if err != nil {
			fmt.Printf("Error rendering design source: %v\n", err)
			continue
		}

		fmt.Printf("Design source:\n%s\n", designSource)

		if dryRun {
			continue
		}

		fileName := fmt.Sprintf("%s.go", design.Namespace())
		dirPath := filepath.Join(projectDirPath, "model", "gen", design.Namespace())
		filePath := filepath.Join(dirPath, fileName)

		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}

		if err = os.WriteFile(filePath, []byte(designSource), 0644); err != nil {
			return err
		}

		fmt.Printf("Wrote file: %s\n", filePath)
	}

	return nil
}
