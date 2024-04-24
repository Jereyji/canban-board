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
	CheckPermissionToBoard(userId, boardId int, accessLevel string) error
}

type Card interface {
	Create(boardId int, card todo.Card) (int, error)
	CheckPermissionToCard(userId, boardId int) error
	GetAll(userId, boardId int) ([]todo.Card, error)
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
		Card:          NewCardPostgres(db),
	}
}
