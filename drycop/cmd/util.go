/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"
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

func checkPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
