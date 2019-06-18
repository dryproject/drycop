/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/enum"
	"github.com/karrick/godirwalk"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type File struct {
	File    string `mapstructure:"file"`
	Markup  string `mapstructure:"markup"`
	Builder string `mapstructure:"builder"`
}

type Config struct {
	Dirs  []string `mapstructure:"dirs"`
	Files []File   `mapstructure:"files"`
}

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

func checkAgainstTemplate(logger *log.Entry, projectDir string, language enum.Language) bool {
	ok := true

	templatePath, err := homedir.Expand(fmt.Sprintf("~/.drycop/templates/%s", language.String()))
	if err != nil {
		logger.WithError(err).Errorf("Unable to find template directory")
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
				ok = checkDirExists(logger, projectDir, expectedPath) && ok
			} else {
				ok = checkFileExists(logger, projectDir, expectedPath) && ok
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
	logger.Info("Checking project")

	builder := detectProjectBuilder(projectDir)
	language := detectProjectLanguage(projectDir, builder)
	if language == enum.UnknownLanguage {
		language = enum.Go // FIXME
	}
	framework := detectProjectFramework(projectDir, builder, language)
	markup := detectProjectMarkup(projectDir, language)
	logger.WithFields(log.Fields{
		"builder":   builder,
		"language":  language,
		"framework": framework,
		"markup":    markup,
	}).Info("Detected project")

	ok := true

	var config Config
	err := viper.UnmarshalKey("check", &config)
	if err != nil {
		logger.WithError(err).Errorf("Invalid configuration")
		return false
	}

	if language != enum.UnknownLanguage {
		ok = checkAgainstTemplate(logger, projectDir, language) && ok
	}

	// TODO: don't check dirs/files from template again, below.

	for _, expectedDir := range config.Dirs {
		ok = checkDirExists(logger, projectDir, expectedDir) && ok
	}

	for _, expectedFile := range config.Files {
		switch expectedFile.Markup {
		case "", markup:
			switch expectedFile.Builder {
			case "", builder.String(): // TODO: support negation filters
				ok = checkFileExists(logger, projectDir, expectedFile.File) && ok
			default:
				// ignore the file
			}
		default:
			// ignore the file
		}
	}

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

	if ok {
		logger.Infof("Checked project: OK")
	} else {
		logger.Warnf("Checked project: some issues")
	}
	return ok
}
