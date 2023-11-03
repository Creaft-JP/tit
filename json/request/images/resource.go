package images

import "github.com/google/uuid"

type Resource struct {
	Id          uuid.UUID `json:"id"`
	Extension   string    `json:"extension"`
	Description string    `json:"description"`
}
