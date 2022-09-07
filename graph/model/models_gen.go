// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}