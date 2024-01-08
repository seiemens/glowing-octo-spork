package models

type status int

const (
	Published status = iota + 1 // EnumIndex = 1
	Hidden                      // EnumIndex = 2
	Deleted                     // EnumIndex = 3
)

type Snippet struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"userid"`
	Status  status `json:"status"`
}
