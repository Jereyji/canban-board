package todo

import "time"

type Board struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type BoardPermission struct {
	Id          int    `json:"-" db:"id"`
	BoardId     int    `json:"board_id" db:"board_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	AccessLevel string `json:"access_level" db:"access_level"`
}

type Card struct {
	Id          int       `json:"-" db:"id"`
	BoardId     int       `json:"board_id" db:"board_id"`
	UserId      int       `json:"user_id" db:"user_id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	StatusCard  string    `json:"status_card" db:"status_card"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
