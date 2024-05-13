package todo

import (
	"errors"
	"time"
)

const (
	ToDo = "to do"
)

type Board struct {
	Id          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type BoardPermission struct {
	UserId      int    `json:"user_id" db:"user_id"`
	BoardId     string `json:"board_id" db:"board_id"`
	AccessLevel string `json:"access_level" db:"access_level"`
}

type Card struct {
	Id          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	DueDate     string    `json:"due_date" db:"due_date"`
	StatusCard  string    `json:"status_card" db:"status_card"`
	UserId      string    `json:"user_id" db:"user_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type BoardCards struct {
	BoardId string `json:"board_id" db:"board_id"`
	CardId  string `json:"card_id" db:"card_id"`
}

type UpdateBoardInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateBoardInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateCardInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DueDate     *string `json:"due_date"`
	UserId      *string `json:"user_id"`
	StatusCard  *string `json:"status_card"`
}

func (i UpdateCardInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.DueDate == nil && i.UserId == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
