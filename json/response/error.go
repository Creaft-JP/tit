package response

type Error struct {
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}
