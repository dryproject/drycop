/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/enum"
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

func checkDirExists(logger *log.Entry, dir string, subdir string) bool {
	logger = logger.WithField("dir", subdir)
	logger.Trace("Checking for directory")

	ok := dirExists(dir, subdir)
	if !ok {
		logger.Warnf("Missing directory: %s", subdir)
	}
	return ok
}

func checkProject(projectDir string) bool {
	logger := log.WithField("project", projectDir)
	logger.Info("Checking project")

	builder := detectProjectBuilder(projectDir)
	language := detectProjectLanguage(projectDir, builder)
	framework := detectProjectFramework(projectDir, builder, language)
	markup := detectProjectMarkup(projectDir, language)
	logger.WithFields(log.Fields{
		"builder":   builder,
		"language":  language,
		"framework": framework,
		"markup":    markup,
	}).Info("Detected project")

	ok := true

	ok = checkDirExists(logger, projectDir, ".git") && ok
	ok = checkFileExists(logger, projectDir, ".gitignore") && ok
	// TODO: .travis.yml?
	ok = checkFileExists(logger, projectDir, "AUTHORS") && ok
	switch markup {
	case "rst": // reStructuredText
		ok = checkFileExists(logger, projectDir, "CHANGES.rst") && ok
	case "md": // Markdown
		switch builder {
		case enum.DartPub:
			ok = checkFileExists(logger, projectDir, "CHANGELOG.md") && ok
		default:
			ok = checkFileExists(logger, projectDir, "CHANGES.md") && ok
		}
	}
	ok = checkFileExists(logger, projectDir, "CREDITS."+markup) && ok
	ok = checkFileExists(logger, projectDir, "Makefile") && ok
	// TODO: README symlink?
	ok = checkFileExists(logger, projectDir, "README."+markup) && ok
	ok = checkFileExists(logger, projectDir, "TODO."+markup) && ok
	ok = checkFileExists(logger, projectDir, "UNLICENSE") && ok
	ok = checkFileExists(logger, projectDir, "VERSION") && ok

	switch language {
	case enum.C:
	case enum.Csharp:
	case enum.Cxx:
	case enum.CommonLisp:
	case enum.D:
	case enum.Dart:
	case enum.DRY:
	case enum.Elixir:
	case enum.Erlang:
	case enum.Go:
	case enum.Java:
	case enum.JavaScript:
	case enum.Julia:
	case enum.Kotlin:
	case enum.Lua:
	case enum.Markdown:
	case enum.ObjectiveC:
	case enum.OCaml:
	case enum.PHP:
	case enum.Python:
	case enum.RestructuredText:
	case enum.Ruby:
	case enum.Rust:
	case enum.Swift:
	case enum.TypeScript:
	case enum.YAML:
	case enum.Zig:
	case enum.UnknownLanguage:
	default:
	}

	// TODO: check .github/
	// TODO: check author of last commit
	// TODO: check file checksums
	// TODO: check change log

	return ok
}
