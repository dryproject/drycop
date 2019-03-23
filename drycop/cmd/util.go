/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/dryproject/drycop/drycop/enum"
)

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

func detectProjectBuilder(projectDir string) enum.Builder {
	// TODO: this should be a YAML configuration file.
	if fileExists(projectDir, "configure.ac") {
		return enum.Autoconf
	}
	if fileExists(projectDir, "Makefile.am") { // must follow the Autoconf check
		return enum.Automake
	}
	if fileExists(projectDir, "CMakeLists.txt") {
		return enum.CMake
	}
	if fileExists(projectDir, "pubspec.yaml") {
		return enum.DartPub
	}
	if fileExists(projectDir, "mix.exs") {
		return enum.ElixirHex
	}
	if fileExists(projectDir, "go.mod") { // TODO: improve this
		return enum.GoBuild
	}
	if fileExists(projectDir, "build.gradle") {
		return enum.Gradle
	}
	if fileExists(projectDir, "pom.xml") {
		return enum.Maven
	}
	if fileExists(projectDir, "dune") {
		return enum.OCamlDune
	}
	if fileExists(projectDir, "setup.py") {
		return enum.PythonPIP
	}
	if fileExists(projectDir, "Gemfile") || filesExist(projectDir, "*.gemspec") {
		return enum.RubyGems
	}
	if fileExists(projectDir, "Package.swift") {
		return enum.SwiftPackageManager
	}
	if filesExist(projectDir, "*.go") { // must remain at the end
		return enum.GoBuild
	}
	if fileExists(projectDir, "Makefile") { // must remain the last check
		return enum.Make
	}
	return enum.UnknownBuilder
}

func detectProjectLanguage(projectDir string, builder enum.Builder) enum.Language {
	// TODO: this should be a YAML configuration file.
	switch builder {
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

func detectProjectFramework(projectDir string, builder enum.Builder, language enum.Language) enum.Framework {
	// TODO: this should be a YAML configuration file.
	if fileExists(projectDir, "library.properties") {
		return enum.Arduino
	}
	if language == enum.Dart && (fileExists(projectDir, ".flutter-plugins") || dirExists(projectDir, "android")) {
		return enum.Flutter
	}
	if language == enum.Java && fileExists(projectDir, "app/build.gradle") {
		return enum.Android
	}
	return enum.UnknownFramework
}

func detectProjectMarkup(projectDir string, language enum.Language) string {
	switch language {
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
