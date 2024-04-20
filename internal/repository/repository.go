package repository

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string) (todo.User, error)
	CheckUser(email string) (int, error)
}

type Board interface {
	Create(userId int, board todo.Board) (int, error)
	AddPermission(boardId, userId int, access string) error
	GetAll(userId int) ([]todo.Board, error)
	GetById(userId, boardId int) (todo.Board, error)
	Delete(userId, boardId int) error
	Update(userId, boardId int, input todo.UpdateBoardInput) error
	CheckPermission(ownerId, boardId int) error
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
