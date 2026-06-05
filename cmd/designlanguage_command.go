package cmd

import (
	"github.com/spf13/cobra"

	// TODO Upgrade.
	klog "github.com/go-kit/kit/log"

	"github.com/mattmunz/designlanguage/appkit"
	"github.com/mattmunz/designlanguage/appkit/model"
)

// TODO Move most of this to AppKit

// TODO Refactor back with sundries
var (
	app = model.NewApp("designlanguage", "1.0.0", ".sundries")
	// TODO Clean up this multiple cli object business.
	dlCLI1 = &dlCLI{model.NewCLI(app.ID(), "designlanguage", "Tools for the Nonzero Sum Design Language.")}
)

type dlCLI struct {
	model.CLI
}

func (c *dlCLI) NewRootCommand(logger klog.Logger) (*cobra.Command, error) {
	return newDesignLanguageCommand(logger), nil
}

func init() {
	appkit.DoInit(app, dlCLI1)
}

func Execute() {
	appkit.Execute(dlCLI1)
}

func newDesignLanguageCommand(logger klog.Logger) *cobra.Command {
	designLanguageCmd := appkit.NewCommandBase(logger)
	designLanguageCmd.AddCommand(newGenCmd(logger))
	return designLanguageCmd
}
