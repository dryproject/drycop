/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CheckCmd describes and implements the `drycop check` command
var CheckCmd = &cobra.Command{
	Use:   "check [dir...]",
	Short: "Check project structure",
	Long:  "Check project structure",
	Args: func(cmd *cobra.Command, args []string) error {
		// Validate all input arguments:
		for _, arg := range args {
			if _, err := validateInputDirectory(arg); err != nil {
				return err
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = append(args, ".")
		}

		// Process the input arguments:
		for _, arg := range args {
			checkProject(arg)
		}
	},
}

func init() {
	RootCmd.AddCommand(CheckCmd)
}

func validateInputDirectory(arg string) (int, error) {
	info, err := os.Stat(arg)
	if err != nil {
		if os.IsNotExist(err) {
			return 66, fmt.Errorf("%s does not exist", arg) // EX_NOINPUT
		}
		panic(err)
	}
	if !info.IsDir() {
		return 66, fmt.Errorf("%s is not a directory", arg) // EX_NOINPUT
	}
	return 0, nil
}

func _checkFileExists(dir string, file string) bool {
	info, err := os.Stat(fmt.Sprintf("%s/%s", dir, file))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	if info.IsDir() {
		return false
	}
	return true
}

func checkFileExists(logger *log.Entry, dir string, file string) bool {
	logger = logger.WithField("file", file)
	logger.Trace("Checking for file")

	exists := _checkFileExists(dir, file)
	if !exists {
		logger.Warnf("Missing file: %s", file)
	}
	return exists
}

func checkProject(dir string) {
	logger := log.WithField("project", dir)
	logger.Info("Checking project")

	checkFileExists(logger, dir, ".gitignore")
	checkFileExists(logger, dir, "AUTHORS")
	checkFileExists(logger, dir, "Makefile")
	checkFileExists(logger, dir, "UNLICENSE")
	checkFileExists(logger, dir, "VERSION")
}
