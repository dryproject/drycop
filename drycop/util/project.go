/* This is free and unencumbered software released into the public domain. */

package util

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dryproject/drycop/drycop/enum"
)

type Project struct {
	Logger       *log.Entry
	Dir          string
	CheckedPaths map[string]bool
	Builder      enum.Builder
	Framework    enum.Framework
	Language     enum.Language
	Markup       string
}

func (project Project) CheckDirExists(dir string) bool {
	if value, ok := project.CheckedPaths[dir]; ok {
		return value
	}

	logger := project.Logger.WithField("dir", dir)
	logger.Trace("Checking for directory")

	ok := project.DirExists(dir)
	project.CheckedPaths[dir] = ok
	if !ok {
		logger.Warnf("Missing directory: %s", dir)
	}
	return ok
}

func (project Project) CheckFileExists(file string) bool {
	if value, ok := project.CheckedPaths[file]; ok {
		return value
	}

	logger := project.Logger.WithField("file", file)
	logger.Trace("Checking for file")

	ok := project.FileExists(file)
	project.CheckedPaths[file] = ok
	if !ok {
		logger.Warnf("Missing file: %s", file)
	}
	return ok
}

func (project Project) DirExists(subdir string) bool {
	dir := project.Dir
	info, err := os.Stat(fmt.Sprintf("%s/%s", dir, subdir))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return info.IsDir()
}

func (project Project) FileExists(file string) bool {
	dir := project.Dir
	info, err := os.Stat(fmt.Sprintf("%s/%s", dir, file))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return !info.IsDir()
}

func (project Project) FilesExist(files string) bool {
	return false // TODO
}

func DetectProject(projectDir string) Project {
	var project Project
	project.Dir = projectDir
	project.CheckedPaths = make(map[string]bool)
	project.Builder = DetectProjectBuilder(project)
	project.Language = DetectProjectLanguage(project)
	if project.Language == enum.UnknownLanguage {
		project.Language = enum.Go // FIXME
	}
	project.Framework = DetectProjectFramework(project)
	project.Markup = DetectProjectMarkup(project)
	return project
}

func DetectProjectBuilder(project Project) enum.Builder {
	// TODO: this should be a YAML configuration file.
	if project.FileExists("configure.ac") {
		return enum.Autoconf
	}
	if project.FileExists("Makefile.am") { // must follow the Autoconf check
		return enum.Automake
	}
	if project.FileExists("CMakeLists.txt") {
		return enum.CMake
	}
	if project.FileExists("pubspec.yaml") {
		return enum.DartPub
	}
	if project.FileExists("mix.exs") {
		return enum.ElixirHex
	}
	if project.FileExists("go.mod") { // TODO: improve this
		return enum.GoBuild
	}
	if project.FileExists("build.gradle") {
		return enum.Gradle
	}
	if project.FileExists("pom.xml") {
		return enum.Maven
	}
	if project.FileExists("dune") {
		return enum.OCamlDune
	}
	if project.FileExists("setup.py") {
		return enum.PythonPIP
	}
	if project.FileExists("Gemfile") || project.FilesExist("*.gemspec") {
		return enum.RubyGems
	}
	if project.FileExists("Package.swift") {
		return enum.SwiftPackageManager
	}
	if project.FilesExist("*.go") { // must remain at the end
		return enum.GoBuild
	}
	if project.FileExists("Makefile") { // must remain the last check
		return enum.Make
	}
	return enum.UnknownBuilder
}

func DetectProjectLanguage(project Project) enum.Language {
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

func DetectProjectFramework(project Project) enum.Framework {
	// TODO: this should be a YAML configuration file.
	if project.FileExists("library.properties") {
		return enum.Arduino
	}
	if project.Language == enum.Dart && (project.FileExists(".flutter-plugins") || project.DirExists("android")) {
		return enum.Flutter
	}
	if project.Language == enum.Java && project.FileExists("app/build.gradle") {
		return enum.Android
	}
	return enum.UnknownFramework
}

func DetectProjectMarkup(project Project) string {
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
