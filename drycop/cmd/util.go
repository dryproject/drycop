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

func fileExists(dir string, file string) bool {
	info, err := os.Stat(fmt.Sprintf("%s/%s", dir, file))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	if info.IsDir() {
		return false
	}
	return true
}

func detectProjectBuilder(projectDir string) enum.Builder {
	if fileExists(projectDir, "configure.ac") {
		return enum.Autotools
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
	if fileExists(projectDir, "go.mod") {
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
	if fileExists(projectDir, "Gemfile") { // TODO: support for *.gemspec
		return enum.RubyGems
	}
	if fileExists(projectDir, "Package.swift") {
		return enum.SwiftPackageManager
	}
	if fileExists(projectDir, "Makefile") { // must remain the last check
		return enum.Make
	}
	return enum.UnknownBuilder
}

func detectProjectLanguage(projectDir string, builder enum.Builder) enum.Language {
	// TODO
	return enum.UnknownLanguage
}

func detectProjectFramework(projectDir string, builder enum.Builder, language enum.Language) enum.Framework {
	// TODO
	return enum.UnknownFramework
}
