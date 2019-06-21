/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/enum"
	"github.com/dryproject/drycop/drycop/util"
	"github.com/karrick/godirwalk"
	"github.com/mitchellh/go-homedir"
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

func checkAgainstTemplate(project util.Project) bool {
	ok := true

	templatePath, err := homedir.Expand(fmt.Sprintf("~/.drycop/templates/%s", project.Language.String()))
	if err != nil {
		project.Logger.WithError(err).Errorf("Unable to find template directory")
		return false
	}

	godirwalk.Walk(templatePath, &godirwalk.Options{
		Callback: func(osPathname string, dirent *godirwalk.Dirent) error {
			if osPathname == templatePath {
				return nil
			}
			expectedPath := osPathname[len(templatePath)+1:]
			if dirent.IsDir() {
				if dirent.Name() == ".git" {
					return filepath.SkipDir
				}
				ok = project.CheckDirExists(expectedPath) && ok
			} else {
				ok = project.CheckFileExists(expectedPath) && ok
			}
			//fmt.Printf("%s %s\n", dirent.ModeType(), expectedPath) // DEBUG
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})

	return ok
}

func checkProject(projectDir string) bool {
	logger := log.WithField("project", projectDir)
	project := util.Project{
		Logger:       logger,
		Dir:          projectDir,
		CheckedPaths: make(map[string]bool),
		Builder:      builderOverride,
		Language:     languageOverride,
		Framework:    frameworkOverride,
		Markup:       markupOverride,
	}

	logger.Info("Checking project")
	project.Detect()
	logger.WithFields(log.Fields{
		"builder":   project.Builder.String(),
		"language":  project.Language.String(),
		"framework": project.Framework.String(),
		"markup":    project.Markup,
	}).Info("Detected project")

	ok := true

	config, err := util.LoadCheckConfig()
	if err != nil {
		logger.WithError(err).Errorf("Invalid configuration")
		return false
	}

	if project.Language != enum.UnknownLanguage {
		ok = checkAgainstTemplate(project) && ok
	}

	for _, expectedDir := range config.Dirs {
		ok = project.CheckDirExists(expectedDir) && ok
	}

	for _, expectedFile := range config.Files {
		switch expectedFile.Markup {
		case "", project.Markup:
			switch expectedFile.Builder {
			case "", project.Builder.String(): // TODO: support negation filters
				ok = project.CheckFileExists(expectedFile.File) && ok
			default:
				// ignore the file
			}
		default:
			// ignore the file
		}
	}

	switch project.Language {
	case enum.C:
	case enum.Csharp:
	case enum.Cpp:
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

	if ok {
		logger.Infof("Checked project: OK")
	} else {
		logger.Warnf("Checked project: some issues")
	}
	return ok
}
