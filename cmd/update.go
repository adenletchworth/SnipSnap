package cmd

import (
	"SnipSnap/db"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a saved code snippet",
	Long: `Update the title, tags, or content of an existing code snippet by its ID.

You can selectively update one or more fields by passing the appropriate flags.

Examples:
  snipsnap update 4 --title "Updated Title"
  snipsnap update 7 --tags "go,refactor" --code "fmt.Println(\"Updated\")"

To view snippet IDs, run:
  snipsnap list

By default, fields you don't specify will remain unchanged.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a snippet ID.")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			fmt.Printf("Invalid ID: %v\n", args[0])
			return
		}

		updates := make(map[string]interface{})

		title, _ := cmd.Flags().GetString("title")
		if cmd.Flags().Changed("title") {
			updates["title"] = title
		}

		tags, _ := cmd.Flags().GetString("tags")
		if cmd.Flags().Changed("tags") {
			updates["tags"] = tags
		}

		code, _ := cmd.Flags().GetString("code")
		if cmd.Flags().Changed("code") {
			updates["content"] = code // if "content" is the DB field
		}

		if len(updates) == 0 {
			fmt.Println("No fields provided to update. Use --title, --tags, or --code.")
			return
		}

		store, err := db.NewSnippetStore("./snippets.db")
		if err != nil {
			log.Fatalf("Failed to open database: %v\n", err)
		}
		defer store.Close()

		err = store.UpdateByID(uint(id), updates)
		if err != nil {
			log.Fatalf("Failed to update snippet: %v\n", err)
		}

		fmt.Printf("✅ Snippet with ID %d updated successfully.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Flags
	updateCmd.Flags().String("title", "", "Title of the snippet")
	updateCmd.Flags().String("tags", "", "Comma-separated list of tags")
	updateCmd.Flags().String("code", "", "Code content of the snippet")

}
