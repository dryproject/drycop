/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// InitCmd describes and implements the `drycop init` command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes configuration",
	Long:  "Initializes configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
