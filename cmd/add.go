package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [code]",
	Short: "Add a new code snippet",
	Long: `Add a new code snippet to your personal snippet store.
You can provide a snippet inline and use flags to add metadata
like programming language or description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add Command")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
