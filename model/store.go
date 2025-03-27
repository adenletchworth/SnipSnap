package model

import (
	"time"
)

type Snippet struct {
	ID        int
	Title     string
	Tags      []string
	Content   string
	CreatedAt time.Time
}
