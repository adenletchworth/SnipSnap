package cmd

import (
	"SnipSnap/db"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved snippets",
	Long: `List all code snippets currently stored locally.
Each snippet will show its ID, title, tags, and creation timestamp.`,
	Run: func(cmd *cobra.Command, args []string) {
		store, err := db.NewSnippetStore("./snippets.db")
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		defer store.Close()

		snippets, err := store.ListSnippets()
		if err != nil {
			log.Fatalf("Failed to retrieve snippets: %v", err)
		}

		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
			return
		}

		fmt.Println("Saved Snippets:")
		for _, snip := range snippets {
			fmt.Printf("\nID: %d\nTitle: %s\nTags: %s\nCreated At: %s\nContent:\n%s\n",
				snip.ID, snip.Title, snip.Tags, snip.CreatedAt.Format(time.RFC3339), snip.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
