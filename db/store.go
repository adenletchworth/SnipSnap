package db

import (
	"SnipSnap/embed"
	"SnipSnap/model"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SnippetStore struct {
	db *sql.DB
}

func NewSnippetStore(path string) (*SnippetStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	store := &SnippetStore{db: db}

	if err := store.initSchema(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *SnippetStore) Close() error {
	return s.db.Close()
}

func (s *SnippetStore) initSchema() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS snippets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			tags TEXT,
			content TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			embedding BLOB
		)
	`)
	return err
}

func (s *SnippetStore) InsertSnippet(snippet model.Snippet) (int64, error) {
	textForEmbedding := snippet.Title + snippet.Tags + snippet.Content

	embedding, err := embed.GetEmbedding(textForEmbedding)
	if err != nil {
		return 0, err
	}

	embeddingBytes := float32SliceToBytes(embedding)
	if embeddingBytes == nil {
		return 0, fmt.Errorf("failed to convert embedding to byte slice")
	}

	result, err := s.db.Exec(`
		INSERT INTO snippets (title, tags, content, embedding)
		VALUES (?, ?, ?, ?)
	`, snippet.Title, snippet.Tags, snippet.Content, embeddingBytes)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (s *SnippetStore) DeleteSnippetWithID(ID uint) error {
	res, err := s.db.Exec(`DELETE FROM snippets WHERE id = ?`, ID)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("no snippet with ID %d", ID)
	}
	return nil
}

func (s *SnippetStore) ListSnippets() ([]model.Snippet, error) {
	rows, err := s.db.Query(`SELECT id, title, tags, content, created_at FROM snippets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []model.Snippet
	for rows.Next() {
		var snip model.Snippet
		err := rows.Scan(&snip.ID, &snip.Title, &snip.Tags, &snip.Content, &snip.CreatedAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snip)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (s *SnippetStore) ListSnippetsWithEmbedding() ([]model.IndexedSnippet, error) {
	rows, err := s.db.Query(`SELECT id, title, tags, content, created_at, embedding FROM snippets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.IndexedSnippet

	for rows.Next() {
		var snip model.Snippet
		var embeddingBytes []byte

		err := rows.Scan(&snip.ID, &snip.Title, &snip.Tags, &snip.Content, &snip.CreatedAt, &embeddingBytes)
		if err != nil {
			return nil, err
		}

		vector, err := bytesToFloat32Slice(embeddingBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to decode embedding for snippet ID %d: %v", snip.ID, err)
		}

		results = append(results, model.IndexedSnippet{
			Snippet: snip,
			Vector:  vector,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// func joinTags(tags []string) string {
// 	return strings.Join(tags, ",")
// }
