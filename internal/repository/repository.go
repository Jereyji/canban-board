package repository

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string) (todo.User, error)
}

type Board interface {
	Create(userId int, board todo.Board) (int, error)
	GetAll(userId int) ([]todo.Board, error)
	GetById(userId, boardId int) (todo.Board, error)
}

type Card interface {
}

type Repository struct {
	Authorization
	Board
	Card
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Board:         NewBoardPostgres(db),
	}
}
