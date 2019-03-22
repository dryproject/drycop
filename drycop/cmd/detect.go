/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
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
			detectProject(arg)
		}
	},
}

func init() {
	RootCmd.AddCommand(DetectCmd)
}

func detectProject(dir string) {
	logger := log.WithField("project", dir)
	logger.Info("Detecting project")

	language := detectProjectLanguage(dir)
	builder := detectProjectBuilder(dir, language)
	framework := detectProjectFramework(dir, language)

	fmt.Printf("%s\tlanguage=%s\tbuilder=%s\tframework=%s\n", dir,
		strings.ToLower(language.String()),
		strings.ToLower(builder.String()),
		strings.ToLower(framework.String()))
}
