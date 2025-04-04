package search

import (
	"SnipSnap/db"
	"SnipSnap/embed"
	"SnipSnap/model"
)

func Search(query string, k int) ([]model.Snippet, error) {
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

	var scoredSnippets []model.ScoredSnippet

	for _, indexedSnippet := range indexedSnippets {
		score := CosineSimilarity(queryEmbedding, indexedSnippet.Vector)

		scoredSnippets = append(scoredSnippets, model.ScoredSnippet{
			Snippet: indexedSnippet.Snippet,
			Score:   score,
		})
	}

	topKSnippets := TopKElements(scoredSnippets, k)

	var topSnippets []model.Snippet
	for _, scored := range topKSnippets {
		topSnippets = append(topSnippets, scored.Snippet)
	}
	return topSnippets, nil
}
