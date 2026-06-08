package model

import (
	klog "github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
)

/*
TODO replace this with code generated from design language.
*/

type App interface {
	ID() string
	Version() string
	ConfigName() string
}

type CLI interface {
	Command
	App() App
	Name() string
	ShortDescription() string
	NewRootCommand() (*cobra.Command, error)
	Logger() klog.Logger
	SetLogger(logger klog.Logger)
}

type Command interface {
	Execute()
}

type CommandFactory interface {
	New() *cobra.Command
}

type appImpl struct {
	id         string
	version    string
	configName string
}

func (a *appImpl) ID() string {
	return a.id
}

func (a *appImpl) Version() string {
	return a.version
}

func (a *appImpl) ConfigName() string {
	return a.configName
}

func NewApp(id, version, configName string) App {
	return &appImpl{
		id:         id,
		version:    version,
		configName: configName,
	}
}
