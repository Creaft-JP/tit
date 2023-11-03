package mb

import "github.com/google/uuid"

type Page struct {
	Pathname    string      `json:"pathname"`
	Title       string      `json:"title"`
	ResourceIds []uuid.UUID `json:"resourceIds"`
	Sections    []Section   `json:"sections"`
}
