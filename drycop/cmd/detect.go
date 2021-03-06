/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/util"
	"github.com/spf13/cobra"
)

// DetectCmd describes and implements the `drycop detect` command
var DetectCmd = &cobra.Command{
	Use:   "detect [dir...]",
	Short: "Detect project characteristics",
	Long:  "Detect project characteristics",
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
			detectProjectAtPath(arg)
		}
	},
}

func init() {
	RootCmd.AddCommand(DetectCmd)
}

func detectProjectAtPath(projectDir string) {
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

	fmt.Printf("%s\tbuilder=%s\tlanguage=%s\tframework=%s\n",
		projectDir,
		strings.ToLower(project.Builder.String()),
		strings.ToLower(project.Language.String()),
		strings.ToLower(project.Framework.String()))
}
