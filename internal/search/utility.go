package search

import (
	"SnipSnap/model"
	"container/heap"
	"math"
)

func CosineSimilarity(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("vectors must be the same length")
	}

	var dotProduct, normA, normB float32

	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dotProduct / float32(math.Sqrt(float64(normA*normB)))
}

func TopKElements(snippets []model.ScoredSnippet, k int) []model.ScoredSnippet {
	if k <= 0 {
		return []model.ScoredSnippet{}
	}

	minHeap := &SnippetMinHeap{}
	heap.Init(minHeap)

	for _, snippet := range snippets {
		if minHeap.Len() < k {
			heap.Push(minHeap, snippet)
		} else if snippet.Score > (*minHeap)[0].Score {
			heap.Pop(minHeap)
			heap.Push(minHeap, snippet)
		}
	}

	// Convert heap to sorted slice (descending by score)
	topK := make([]model.ScoredSnippet, 0, minHeap.Len())
	for minHeap.Len() > 0 {
		topK = append([]model.ScoredSnippet{heap.Pop(minHeap).(model.ScoredSnippet)}, topK...)
	}

	return topK
}
