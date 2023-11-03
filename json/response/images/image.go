package images

import "github.com/google/uuid"

type Image struct {
	Id        uuid.UUID `json:"id"`
	UploadUrl string    `json:"uploadURL"`
}
