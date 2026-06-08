// TODO Move this to its own project.
package appkit

import (
	"errors"
	"fmt"
	"os"

	"github.com/mattmunz/designlanguage/appkit/cmd"
	"github.com/mattmunz/designlanguage/appkit/misc"
	"github.com/mattmunz/designlanguage/appkit/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	klog "github.com/go-kit/kit/log"
)

var (
	cfgFile string
)

func DoInit(cli1 model.CLI) {
	logger := klog.NewLogfmtLogger(klog.NewSyncWriter(os.Stdout))

	misc.LogMessage(logger, "Logger initialized.")

	cli1.SetLogger(logger)

	misc.LogMessage(cli1.Logger(), "CLI initialized.")

	initConfig2 := func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		}

		viper.SetConfigName(cli1.App().ConfigName())
		viper.AddConfigPath("$HOME")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			misc.LogMessage(cli1.Logger(), fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
		}
	}

	cobra.OnInitialize(initConfig2)
}

type cliImpl struct {
	app              model.App
	name             string
	shortDescription string
	logger           klog.Logger
}

func (c *cliImpl) App() model.App {
	return c.app
}

func (c *cliImpl) Name() string {
	return c.name
}

func (c *cliImpl) ShortDescription() string {
	return c.shortDescription
}

func (c *cliImpl) Logger() klog.Logger {
	return c.logger
}

func (c *cliImpl) SetLogger(logger klog.Logger) {
	c.logger = logger
}

func (c *cliImpl) NewRootCommand() (*cobra.Command, error) {
	return nil, errors.New("Not implemented")
}

func (c *cliImpl) Execute() {
	rootCmd, err := c.NewRootCommand()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func NewCLI(app model.App, name, shortDescription string) model.CLI {
	return &cliImpl{
		app:              app,
		name:             name,
		shortDescription: shortDescription,
	}
}
func NewCommandBase(cli model.CLI) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   cli.Name(),
		Short: cli.ShortDescription(),
	}
	newCmd.AddCommand(newVersionCommand(cli))
	newCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		fmt.Sprintf("config file (default is $HOME/%s.yaml)", cli.App().ConfigName()))

	return newCmd
}

// TODO Maybe make method on CLI.
func Execute(cli model.CLI) {
	rootCmd, err := cli.NewRootCommand()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func newVersionCommand(cli model.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version number of this command.",
		Run:   cmd.WrapRunner(doVersion, cli),
	}
}

func doVersion(cli model.CLI, _ *cobra.Command, _ []string) error {
	fmt.Println(cli.App().Version())
	return nil
}
