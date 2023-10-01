package json

type Section struct {
	Slug    string   `json:"slug"`
	Title   string   `json:"title"`
	Commits []Commit `json:"commits"`
}
