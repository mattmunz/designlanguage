package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mattmunz/appkit"
	"github.com/mattmunz/appkit/model"
)

var (
	cli = &dlCLI{
		appkit.NewCLI(
			model.NewApp("designlanguage", "1.0.0", ".sundries"),
			"designlanguage",
			"Tools for the Nonzero Sum Design Language.",
		)}
)

type dlCLI struct {
	model.CLI
}

func (c *dlCLI) NewRootCommand() (*cobra.Command, error) {
	cmd := appkit.NewCommandBase(c)
	cmd.AddCommand(newGenCmd(c))
	return cmd, nil
}

func init() {
	appkit.DoInit(cli)
}

func CLI() model.CLI {
	return cli
}
