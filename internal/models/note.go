package models

import "time"

type Note struct {
	Id        int       `json:"id" db:"id"`
	UserId    int       `json:"-" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Content   *string   `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type InputNote struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}
