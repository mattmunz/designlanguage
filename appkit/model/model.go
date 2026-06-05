package model

import (
	"errors"

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
	AppID() string
	Name() string
	ShortDescription() string
	NewRootCommand(logger klog.Logger) (*cobra.Command, error)
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

type cliImpl struct {
	appID            string
	name             string
	shortDescription string
}

func (c *cliImpl) AppID() string {
	return c.appID
}

func (c *cliImpl) Name() string {
	return c.name
}

func (c *cliImpl) ShortDescription() string {
	return c.shortDescription
}

func (c *cliImpl) NewRootCommand(_ klog.Logger) (*cobra.Command, error) {
	return nil, errors.New("Not implemented")
}

func NewCLI(appID, name, shortDescription string) CLI {
	return &cliImpl{
		appID:            appID,
		name:             name,
		shortDescription: shortDescription,
	}
}
