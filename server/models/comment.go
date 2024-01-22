package models

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	PostID  string `json:"postid"`
	Author  string `json:"author"`
}
