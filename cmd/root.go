package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snipsnap",
	Short: "SnipSnap is a personal code snippet manager.",
	Long: `SnipSnap is a fast, minimal CLI tool to store, search, and manage
your code snippets from the command line. You can add, list, remove, 
and fuzzy search snippets using powerful subcommands.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
