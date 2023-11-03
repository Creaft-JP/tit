package mb

import "github.com/google/uuid"

type Commit struct {
	Message     string      `json:"message"`
	Files       []File      `json:"files"`
	ResourceIds []uuid.UUID `json:"resourceIds"`
}
