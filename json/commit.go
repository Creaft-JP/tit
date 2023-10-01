package json

type Commit struct {
	Message string `json:"message"`
	Files   []File `json:"files"`
}
