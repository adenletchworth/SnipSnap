package search

import (
	"SnipSnap/model"
)

type SnippetMinHeap []model.ScoredSnippet

func (h SnippetMinHeap) Len() int           { return len(h) }
func (h SnippetMinHeap) Less(i, j int) bool { return h[i].Score < h[j].Score } // Min-heap
func (h SnippetMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SnippetMinHeap) Push(x interface{}) {
	*h = append(*h, x.(model.ScoredSnippet))
}

func (h *SnippetMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
