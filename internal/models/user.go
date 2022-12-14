package models

import "time"

type User struct {
	Id        int       `json:"-" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type InputUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
