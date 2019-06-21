/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/util"
	"github.com/spf13/cobra"
)

// TestCmd describes and implements the `drycop test` command
var TestCmd = &cobra.Command{
	Use:   "test [dir...]",
	Short: "Run the project test suite, if any",
	Long:  "Run the project test suite, if any",
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
			runTestSuiteAt(arg)
		}
	},
}

func init() {
	RootCmd.AddCommand(TestCmd)
}

func runTestSuiteAt(projectDir string) {
	project := util.NewProject(projectDir, log.WithFields(log.Fields{}))
	project.Builder = builderOverride
	project.Language = languageOverride
	project.Framework = frameworkOverride
	project.Markup = markupOverride

	project.Logger.Info("Detecting project")
	project.Detect()
	project.Logger.WithFields(log.Fields{
		"builder":   project.Builder.String(),
		"language":  project.Language.String(),
		"framework": project.Framework.String(),
		"markup":    project.Markup,
	}).Info("Detected project")

	err := os.Chdir(project.Dir)
	if err != nil {
		project.Logger.WithError(err).Errorf("Failed to change the working directory")
		os.Exit(66) // EX_NOINPUT
	}

	command := project.Framework.TestCommand()
	if command == "" {
		command = project.Builder.TestCommand()
		if command == "" {
			command = project.Language.TestCommand()
		}
	}

	if command == "" {
		project.Logger.Errorf("Failed to detect a test suite runner")
		os.Exit(64) // EX_USAGE
	}

	logger := project.Logger.WithField("command", command)
	logger.Info("Running the test suite")

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		logger.WithError(err).Errorf("Test suite: failed")
		os.Exit(70) // EX_SOFTWARE
	} else {
		logger.Info("Test suite: OK")
	}
}
