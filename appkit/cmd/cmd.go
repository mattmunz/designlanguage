package cmd

import (
	"fmt"
	"os"

	klog "github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
)

type commandRunner func(cmd *cobra.Command, args []string)

func WrapRunner(run func(klog.Logger, *cobra.Command, []string) error, logger klog.Logger) commandRunner {
	return func(command *cobra.Command, args []string) {
		if err := run(logger, command, args); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
			os.Exit(1)
		}
	}
}
