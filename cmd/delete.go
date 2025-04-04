package cmd

import (
	"SnipSnap/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a code snippet",
	Long: `Delete a saved code snippet from your local collection using its unique ID.

To find snippet IDs, use the 'list' command:
  snipsnap list

Then delete a snippet like so:
  snipsnap delete 3

This permanently removes the snippet from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a snippet ID to delete.")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			fmt.Printf("Invalid ID: %v\n", args[0])
			return
		}

		store, err := db.NewSnippetStore("./snippets.db")
		if err != nil {
			fmt.Printf("Failed to open database: %v\n", err)
			return
		}
		defer store.Close()

		err = store.DeleteSnippetWithID(uint(id))
		if err != nil {
			fmt.Printf("Failed to delete snippet with ID %d: %v\n", id, err)
			return
		}

		fmt.Printf("✅ Snippet with ID %d successfully deleted.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
