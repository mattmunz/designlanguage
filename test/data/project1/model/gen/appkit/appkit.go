// Appkit is an application development kit, part of the Nonzero Sum Stack.
// Author: Nonzero Sum
package appkit

import (
	log "github.com/go-kit/kit/log"
	cobra "github.com/spf13/cobra"
)

// A software application.
type App interface {
	ID() string
	// Semver preferred.
	Version() string
	ConfigName() string
}

type Command interface {
	Execute()
}

// Command Line Interface.
type CLI interface {
	Command

	AppID() string
	Name() string
	ShortDescription() string
	NewRootCommand() (cmd *cobra.Command, err error)
	SetLogger(logger log.Logger)
}

type CommandFactory interface {
	New() (cmd *cobra.Command)
}
