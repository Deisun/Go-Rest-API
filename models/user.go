package models

type User struct {
	id        uint32 `json:"id"`
	name      string `json:"name"`
	email     string `json:"email"`
	expertise string `json:"expertise"`
}
