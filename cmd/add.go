package cmd

import (
	"SnipSnap/db"
	"SnipSnap/model"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [code]",
	Short: "Add a new code snippet",
	Long: `Add a new code snippet to your personal snippet store.
You can provide a snippet inline and use flags to add metadata
like programming language or description.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		tags, _ := cmd.Flags().GetString("tags")
		codeContent, _ := cmd.Flags().GetString("code")

		snippet := model.Snippet{
			Title:   title,
			Tags:    strings.Split(tags, ","),
			Content: codeContent,
		}

		store, err := db.NewSnippetStore("./snippets.db")

		if err != nil {
			log.Fatal(err)
		}

		defer store.Close()

		id, err := store.InsertSnippet(snippet)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Snippet added with ID %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Flags
	addCmd.Flags().String("title", "", "Title of the snippet")
	addCmd.Flags().String("tags", "", "Comma-separated list of tags")
	addCmd.Flags().String("code", "", "Code content of the snippet")
}
