/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// TestCmd describes and implements the `drycop test` command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Run the test suite, if any",
	Long:  "Run the test suite, if any",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	RootCmd.AddCommand(TestCmd)
}
