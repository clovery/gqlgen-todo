// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}
