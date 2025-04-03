package db

import (
	"SnipSnap/model"
	"database/sql"
	"strings"
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
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func (s *SnippetStore) InsertSnippet(snippet model.Snippet) (int64, error) {
	result, err := s.db.Exec(`
		INSERT INTO snippets (title, tags, content)
		VALUES (?, ?, ?)
	`, snippet.Title, joinTags(snippet.Tags), snippet.Content)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
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

func joinTags(tags []string) string {
	return strings.Join(tags, ",")
}
