package todo

import "time"

type User struct {
	Id        string       `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
