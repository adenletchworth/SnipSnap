package cmd

import (
	"SnipSnap/db"
	"SnipSnap/model"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var searchCommand = &cobra.Command{
	Use:   "search [text]",
	Short: "Search for a code snippet",
	Long: `Search your saved snippets using semantic similarity.
This will compare your query to snippet content, titles, and tags using vector embeddings.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide search text. Example: snipsnap search \"http server in go\"")
			return
		}
		
		text := strings.Join(args, " ")  

		snippets, err := //search command

		if err != nil {
			...
		}

		fmt.Println("")
	},
}

