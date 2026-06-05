// TODO Move this to its own project.
package appkit

import (
	"fmt"
	"os"

	"github.com/mattmunz/designlanguage/appkit/cmd"
	"github.com/mattmunz/designlanguage/appkit/misc"
	"github.com/mattmunz/designlanguage/appkit/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	klog "github.com/go-kit/kit/log"
)

// TODO All the package-level state is gross. Make an object or something.
var (
	app     model.App
	cli     model.CLI
	cfgFile string
	logger  klog.Logger
)

func Logger() klog.Logger {
	return logger
}

func DoInit(app1 model.App, cli1 model.CLI) {
	app = app1
	cli = cli1
	cobra.OnInitialize(initConfig)

	logger = klog.NewLogfmtLogger(klog.NewSyncWriter(os.Stdout))

	misc.LogMessage(logger, "Logger initialized.")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(app.ConfigName())
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// TODO Log
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func NewCommandBase(logger klog.Logger) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   cli.Name(),
		Short: cli.ShortDescription(),
	}
	newCmd.AddCommand(newVersionCommand(logger))
	newCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		fmt.Sprintf("config file (default is $HOME/%s.yaml)", app.ConfigName()))

	return newCmd
}

// TODO Maybe make method on CLI.
func Execute(cli model.CLI) {
	rootCmd, err := cli.NewRootCommand(logger)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func newVersionCommand(logger klog.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version number of this command.",
		Run:   cmd.WrapRunner(doVersion, logger),
	}
}

func doVersion(_ klog.Logger, _ *cobra.Command, _ []string) error {
	fmt.Println(app.Version())
	return nil
}
