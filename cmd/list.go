package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved snippets",
	Long: `List all code snippets currently stored locally.
Each snippet will show its ID, title, tags, and creation timestamp.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List Command")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
