package models

type Post struct {
	id     uint32 `json:"id"`
	userId uint32 `json:"userId"`
	title  string `json:"title"`
	body   string `json:"body"`
}
