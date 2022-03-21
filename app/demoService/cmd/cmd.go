package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/AIYEE/go-demo/pkg/logging"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	optionNameLoggerFile      = "logger-file"
	optionNameVerbosity       = "verbosity"
	optionNameDbDriver        = "db-driver"
	optionNameDbFile          = "db-file"
	optionNameEthSinger       = "eth-signer"
	optionNameRPCEndPoint     = "endpoint"
	optionNameContractAddress = "contract-address"
	optionNameSendGas         = "send-gas"
)

type command struct {
	root    *cobra.Command
	config  *viper.Viper
	cfgFile string
	homeDir string
}

type option func(*command)

func new(opts ...option) (c *command, err error) {
	c = &command{
		root: &cobra.Command{
			Use:           "demo",
			Short:         "This is demo command.",
			SilenceErrors: true,
			SilenceUsage:  true,
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return c.initConfig()
			},
		},
	}

	for _, o := range opts {
		o(c)
	}

	if err := c.setHomeDir(); err != nil {
		return nil, err
	}

	c.initGlobalFlags()

	c.initVersionCmd()

	c.initStartCmd()

	return c, nil
}

func (c *command) Execute() (err error) {
	return c.root.Execute()
}

func Execute() (err error) {
	c, err := new()
	if err != nil {
		return err
	}

	return c.Execute()
}

func (c *command) initConfig() (err error) {
	config := viper.New()
	configName := ".demoConfig"
	if c.cfgFile != "" {
		config.SetConfigFile(c.cfgFile)
	} else {
		config.AddConfigPath(c.homeDir)
		config.SetConfigName(configName)
	}

	// Environment
	config.SetEnvPrefix("settlement")
	config.AutomaticEnv() // read in environment variables that match
	config.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if c.homeDir != "" && c.cfgFile == "" {
		c.cfgFile = filepath.Join(c.homeDir, configName+".yaml")
	}

	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) {
			return err
		}
	}
	c.config = config
	return nil
}

func (c *command) setHomeDir() (err error) {
	if c.homeDir != "" {
		return
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	c.homeDir = dir
	return nil
}

func (c *command) initGlobalFlags() {
	globalFlags := c.root.PersistentFlags()
	globalFlags.StringVar(&c.cfgFile, "config", "", "config file (default is $HOME/.demoConfig.yaml)")
}

func (c *command) setAllFlags(cmd *cobra.Command) {
	cmd.Flags().String(optionNameVerbosity, "debug", "log verbosity level 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=trace")
	cmd.Flags().String(optionNameLoggerFile, "./running.log", "log file")
	cmd.Flags().String(optionNameDbDriver, "sqlite3", "db driver name")
	cmd.Flags().String(optionNameDbFile, "./MinerNodes.db", "db file name")
	cmd.Flags().String(optionNameEthSinger, "668db522795b07ba80e178fbb1972f4487d39ebaaadc4bd50416ef3022e61e84", "BNB signer")
	cmd.Flags().String(optionNameRPCEndPoint, "https://data-seed-prebsc-1-s1.binance.org:8545", "rpc endpoint")
	cmd.Flags().String(optionNameContractAddress, "0xf4768352a50ccfe26490501cd8dc6716af80c843", "confirm contract address")
	cmd.Flags().Bool(optionNameSendGas, false, "need send gas")
}

func (c *command) NewLogger(verbosity string) (logging.Logger, error) {
	var logger logging.Logger
	writer := c.root.OutOrStdout()
	file := c.config.GetString(optionNameLoggerFile)
	if file != "" {
		writer = logging.CreateFileWriter(file)
	}
	switch verbosity {
	case "0", "silent":
		logger = logging.New(io.Discard, 0)
	case "1", "error":
		logger = logging.New(writer, logrus.ErrorLevel)
	case "2", "warn":
		logger = logging.New(writer, logrus.WarnLevel)
	case "3", "info":
		logger = logging.New(writer, logrus.InfoLevel)
	case "4", "debug":
		logger = logging.New(writer, logrus.DebugLevel)
	case "5", "trace":
		logger = logging.New(writer, logrus.TraceLevel)
	default:
		return nil, fmt.Errorf("unknown verbosity level %q", verbosity)
	}
	return logger, nil

}
