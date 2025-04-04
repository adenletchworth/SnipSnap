package search

import (
	"SnipSnap/db"
	"SnipSnap/embed"
	"SnipSnap/model"
	"errors"
	"strings"
)

func SearchSnippets(query string, k int) ([]model.ScoredSnippet, error) {
	queryEmbedding, err := embed.GetEmbedding(query)
	if err != nil {
		return nil, err
	}

	store, err := db.NewSnippetStore("./snippets.db")
	if err != nil {
		return nil, err
	}
	defer store.Close()

	indexedSnippets, err := store.ListSnippetsWithEmbedding()
	if err != nil {
		return nil, err
	}
	if len(indexedSnippets) == 0 {
		return nil, errors.New("no embedded snippets found")
	}

	var scoredSnippets []model.ScoredSnippet
	for _, indexed := range indexedSnippets {
		score := CosineSimilarity(queryEmbedding, indexed.Vector)

		if strings.Contains(strings.ToLower(indexed.Snippet.Title), strings.ToLower(query)) {
			score += 0.1
		}
		if strings.Contains(strings.ToLower(indexed.Snippet.Tags), strings.ToLower(query)) {
			score += 0.05
		}

		scoredSnippets = append(scoredSnippets, model.ScoredSnippet{
			Snippet: indexed.Snippet,
			Score:   score,
		})
	}

	topK := TopKElements(scoredSnippets, k)

	return topK, nil
}
