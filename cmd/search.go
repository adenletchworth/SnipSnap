package cmd

import (
	"SnipSnap/internal/search"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [text]",
	Short: "Search for a code snippet",
	Long: `Search your saved snippets using semantic similarity.
This will compare your query to snippet content, titles, and tags using vector embeddings.`,
	Run: func(cmd *cobra.Command, args []string) {
		k, _ := cmd.Flags().GetInt("k")

		if len(args) == 0 {
			fmt.Println("Please provide search text. Example: snipsnap search \"http server in go\"")
			return
		}

		text := strings.Join(args, " ")

		results, err := search.SearchSnippets(text, k)
		if err != nil {
			log.Fatalf("Search failed: %v", err)
		}

		fmt.Printf("Top %d Results:\n\n", len(results))
		for i, result := range results {
			fmt.Printf("%d. %s (Score: %.4f)\n", i+1, result.Snippet.Title, result.Score)
			fmt.Printf("   Tags: %s\n", result.Snippet.Tags)
			fmt.Printf("   Content: %s\n\n", result.Snippet.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().Int("k", 5, "Number of snippets to return")
}
