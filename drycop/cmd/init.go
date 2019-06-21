/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/enum"
	"github.com/dryproject/drycop/drycop/util"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// InitCmd describes and implements the `drycop init` command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes configuration",
	Long:  "Initializes configuration",
	Run: func(cmd *cobra.Command, args []string) {
		initTemplates()
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}

func initTemplates() {
	config, err := util.LoadTemplatesConfig()
	if err != nil {
		panic(err)
	}

	if languageOverride != enum.UnknownLanguage {
		initTemplateForLanguage(config, languageOverride)
	} else {
		initTemplateForLanguage(config, enum.C)
		initTemplateForLanguage(config, enum.Csharp)
		initTemplateForLanguage(config, enum.Cpp)
		initTemplateForLanguage(config, enum.CommonLisp)
		initTemplateForLanguage(config, enum.D)
		initTemplateForLanguage(config, enum.Dart)
		initTemplateForLanguage(config, enum.DRY)
		initTemplateForLanguage(config, enum.Elixir)
		initTemplateForLanguage(config, enum.Erlang)
		initTemplateForLanguage(config, enum.Go)
		initTemplateForLanguage(config, enum.Java)
		initTemplateForLanguage(config, enum.JavaScript)
		initTemplateForLanguage(config, enum.Julia)
		initTemplateForLanguage(config, enum.Kotlin)
		initTemplateForLanguage(config, enum.Lua)
		initTemplateForLanguage(config, enum.Markdown)
		initTemplateForLanguage(config, enum.ObjectiveC)
		initTemplateForLanguage(config, enum.OCaml)
		initTemplateForLanguage(config, enum.PHP)
		initTemplateForLanguage(config, enum.Python)
		initTemplateForLanguage(config, enum.RestructuredText)
		initTemplateForLanguage(config, enum.Ruby)
		initTemplateForLanguage(config, enum.Rust)
		initTemplateForLanguage(config, enum.Swift)
		initTemplateForLanguage(config, enum.TypeScript)
		initTemplateForLanguage(config, enum.YAML)
		initTemplateForLanguage(config, enum.Zig)
	}
}

func initTemplateForLanguage(config *util.TemplatesConfig, language enum.Language) {
	logger := log.WithField("language", language.String())

	gitURL := config.Templates[language.String()]
	if gitURL == "" {
		return // skip any language without a configured template
	}

	gitPath, err := homedir.Expand(fmt.Sprintf("~/.drycop/templates/%s", language.String()))
	if err != nil {
		panic(err)
	}

	exists, err := checkPathExists(gitPath)
	if err != nil {
		panic(err)
	}
	if exists {
		return // skip any already cloned templates
	}

	command := fmt.Sprintf("git clone %s %s", gitURL, gitPath)

	logger = logger.WithField("command", command)
	logger.Info("Cloning template")

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		logger.WithError(err).Errorf("Cloning template: failed")
		os.Exit(70) // EX_SOFTWARE
	} else {
		logger.Info("Cloning template: OK")
	}
}
