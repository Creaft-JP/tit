package json

type Page struct {
	Pathname string    `json:"pathname"`
	Title    string    `json:"title"`
	Sections []Section `json:"sections"`
}
