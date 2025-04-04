package model

import (
	"time"
)

type Snippet struct {
	ID        int
	Title     string
	Tags      string
	Content   string
	CreatedAt time.Time
}

type IndexedSnippet struct {
	Snippet Snippet
	Vector  []float32
}

type ScoredSnippet struct {
	Snippet Snippet
	Score   float32
}
