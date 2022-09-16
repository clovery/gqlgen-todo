package model

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	User      *User     `json:"user"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// type User struct {
// 	ID   int    `json:"id" gorm:"primarykey"`
// 	Name string `json:"name"`
// }
