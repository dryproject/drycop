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

func detectProjectLanguage(dir string) enum.Language {
	// TODO
	return enum.UnknownLanguage
}

func detectProjectFramework(dir string, lang enum.Language) enum.Framework {
	// TODO
	return enum.UnknownFramework
}

func detectProjectBuilder(dir string, lang enum.Language) enum.Builder {
	// TODO
	return enum.UnknownBuilder
}
