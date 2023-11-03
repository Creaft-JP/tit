package mb

import "github.com/google/uuid"

type Section struct {
	Slug        string      `json:"slug"`
	Title       string      `json:"title"`
	ResourceIds []uuid.UUID `json:"resourceIds"`
	Commits     []Commit    `json:"commits"`
}
