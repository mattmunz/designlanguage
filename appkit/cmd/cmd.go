package cmd

import (
	"fmt"
	"os"

	"github.com/mattmunz/designlanguage/appkit/model"
	"github.com/spf13/cobra"
)

type commandRunner func(cmd *cobra.Command, args []string)

func WrapRunner(run func(_ model.CLI, _ *cobra.Command, _ []string) error, cli model.CLI) commandRunner {
	return func(command *cobra.Command, args []string) {
		if err := run(cli, command, args); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
			os.Exit(1)
		}
	}
}
