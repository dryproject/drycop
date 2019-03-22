/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CheckCmd describes and implements the `drycop check` command
var CheckCmd = &cobra.Command{
	Use:   "check [dir...]",
	Short: "Check project conformance",
	Long:  "Check project conformance",
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
		ok := true
		for _, arg := range args {
			ok = checkProject(arg) && ok
		}

		if !ok {
			os.Exit(1) // TODO: use EX_* exit codes
		}
	},
}

func init() {
	RootCmd.AddCommand(CheckCmd)
}

func checkFileExists(logger *log.Entry, dir string, file string) bool {
	logger = logger.WithField("file", file)
	logger.Trace("Checking for file")

	ok := fileExists(dir, file)
	if !ok {
		logger.Warnf("Missing file: %s", file)
	}
	return ok
}

func checkProject(dir string) bool {
	logger := log.WithField("project", dir)
	logger.Info("Checking project")

	ok := true
	ok = checkFileExists(logger, dir, ".gitignore") && ok
	ok = checkFileExists(logger, dir, "AUTHORS") && ok
	ok = checkFileExists(logger, dir, "Makefile") && ok
	ok = checkFileExists(logger, dir, "UNLICENSE") && ok
	ok = checkFileExists(logger, dir, "VERSION") && ok

	return ok
}
