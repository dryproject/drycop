/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/enum"
)

type Project struct {
	Logger    *log.Entry
	Dir       string
	Builder   enum.Builder
	Framework enum.Framework
	Language  enum.Language
	Markup    string
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

func dirExists(dir string, subdir string) bool {
	info, err := os.Stat(fmt.Sprintf("%s/%s", dir, subdir))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return info.IsDir()
}

func fileExists(dir string, file string) bool {
	info, err := os.Stat(fmt.Sprintf("%s/%s", dir, file))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return !info.IsDir()
}

func filesExist(dir string, files string) bool {
	return false // TODO
}

func detectProject(projectDir string) Project {
	var project Project
	project.Dir = projectDir
	project.Builder = detectProjectBuilder(project)
	project.Language = detectProjectLanguage(project)
	if project.Language == enum.UnknownLanguage {
		project.Language = enum.Go // FIXME
	}
	project.Framework = detectProjectFramework(project)
	project.Markup = detectProjectMarkup(project)
	return project
}

func detectProjectBuilder(project Project) enum.Builder {
	// TODO: this should be a YAML configuration file.
	if fileExists(project.Dir, "configure.ac") {
		return enum.Autoconf
	}
	if fileExists(project.Dir, "Makefile.am") { // must follow the Autoconf check
		return enum.Automake
	}
	if fileExists(project.Dir, "CMakeLists.txt") {
		return enum.CMake
	}
	if fileExists(project.Dir, "pubspec.yaml") {
		return enum.DartPub
	}
	if fileExists(project.Dir, "mix.exs") {
		return enum.ElixirHex
	}
	if fileExists(project.Dir, "go.mod") { // TODO: improve this
		return enum.GoBuild
	}
	if fileExists(project.Dir, "build.gradle") {
		return enum.Gradle
	}
	if fileExists(project.Dir, "pom.xml") {
		return enum.Maven
	}
	if fileExists(project.Dir, "dune") {
		return enum.OCamlDune
	}
	if fileExists(project.Dir, "setup.py") {
		return enum.PythonPIP
	}
	if fileExists(project.Dir, "Gemfile") || filesExist(project.Dir, "*.gemspec") {
		return enum.RubyGems
	}
	if fileExists(project.Dir, "Package.swift") {
		return enum.SwiftPackageManager
	}
	if filesExist(project.Dir, "*.go") { // must remain at the end
		return enum.GoBuild
	}
	if fileExists(project.Dir, "Makefile") { // must remain the last check
		return enum.Make
	}
	return enum.UnknownBuilder
}

func detectProjectLanguage(project Project) enum.Language {
	// TODO: this should be a YAML configuration file.
	switch project.Builder {
	case enum.Autoconf:
		return enum.Cxx // TODO: also C, etc
	case enum.Automake:
		return enum.Cxx // TODO: also C, etc
	case enum.CMake:
		return enum.Cxx // TODO: also C, etc
	case enum.DartPub:
		return enum.Dart
	case enum.ElixirHex:
		return enum.Elixir
	case enum.GoBuild:
		return enum.Go
	case enum.Gradle:
		return enum.Java // TODO: also Kotlin, etc
	case enum.Make:
		return enum.UnknownLanguage // TODO: also C++, C, etc
	case enum.Maven:
		return enum.Java // TODO: also Kotlin, etc
	case enum.OCamlDune:
		return enum.OCaml
	case enum.PythonPIP:
		return enum.Python
	case enum.RubyGems:
		return enum.Ruby
	case enum.SwiftPackageManager:
		return enum.Swift
	}
	return enum.UnknownLanguage
}

func detectProjectFramework(project Project) enum.Framework {
	// TODO: this should be a YAML configuration file.
	if fileExists(project.Dir, "library.properties") {
		return enum.Arduino
	}
	if project.Language == enum.Dart && (fileExists(project.Dir, ".flutter-plugins") || dirExists(project.Dir, "android")) {
		return enum.Flutter
	}
	if project.Language == enum.Java && fileExists(project.Dir, "app/build.gradle") {
		return enum.Android
	}
	return enum.UnknownFramework
}

func detectProjectMarkup(project Project) string {
	switch project.Language {
	case enum.C:
	case enum.Csharp:
	case enum.Cxx:
	case enum.CommonLisp:
	case enum.D:
	case enum.Dart:
		return "md"
	case enum.DRY:
	case enum.Elixir:
		return "md"
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
		return "md"
	case enum.Rust:
	case enum.Swift:
	case enum.TypeScript:
	case enum.YAML:
	case enum.Zig:
	case enum.UnknownLanguage:
	default:
	}
	return "rst"
}
