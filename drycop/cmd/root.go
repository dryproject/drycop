/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"go/build"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/enum"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var outputFormat string
var builderOverride enum.Builder
var languageOverride enum.Language
var frameworkOverride enum.Framework
var markupOverride string
var ignoredPath string
var debug bool
var verbose bool

// RootCmd describes the `drycop` command
var RootCmd = &cobra.Command{
	Use:     "drycop",
	Short:   "DRYcop",
	Long:    `DRYcop is a lint tool for DRY projects.`,
	Version: "0.0.0",
}

// Execute implements the `drycop` command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "C", "", "Set config file (default: $HOME/.drycop/config.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debugging")
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "Set output format to 'text' or 'json' (default: text)")
	RootCmd.PersistentFlags().VarP(&builderOverride, "builder", "B", "Set project builder")
	RootCmd.PersistentFlags().VarP(&languageOverride, "language", "L", "Set project language")
	RootCmd.PersistentFlags().VarP(&frameworkOverride, "framework", "F", "Set project framework")
	RootCmd.PersistentFlags().StringVarP(&markupOverride, "markup", "M", "", "Set project markup format")
	RootCmd.PersistentFlags().StringVarP(&ignoredPath, "ignore", "I", "", "Skip a file or directory")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	RootCmd.SetVersionTemplate(`DRYcop {{printf "%s" .Version}}
`)
}

// initConfig reads the configuration file and environment variables.
func initConfig() {
	if configFile != "" {
		// Use the specified config file:
		viper.SetConfigFile(configFile)
	} else {
		// Search for config file in the current directory and under the home directory:
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.drycop")
		viper.AddConfigPath(fmt.Sprintf("%s/src/%s", build.Default.GOPATH, "github.com/dryproject/drycop"))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in:
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}

	// Configure the logger output format:
	log.SetOutput(os.Stderr)
	if outputFormat == "" {
		outputFormat = viper.GetString("output")
	}
	switch outputFormat {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	// Configure the logger level:
	switch {
	case debug:
		log.SetLevel(log.TraceLevel)
	case verbose:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
}
