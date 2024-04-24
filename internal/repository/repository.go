package repository

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (string, error)
	GetUser(username string) (todo.User, error)
	CheckUser(email string) (string, error)
}

type Board interface {
	Create(userId string, board todo.Board) (string, error)
	AddPermission(boardId, userId, access string) error
	GetAll(userId string) ([]todo.Board, error)
	GetById(userId, boardId string) (todo.Board, error)
	Delete(userId, boardId string) error
	Update(userId, boardId string, input todo.UpdateBoardInput) error
	CheckPermissionToBoard(userId, boardId, accessLevel string) error
}

type Card interface {
	Create(boardId string, card todo.Card) (string, error)
	CheckPermissionToCard(userId, boardId string) error
	GetAll(userId, boardId string) ([]todo.Card, error)
	GetById(userId, cardId string) (todo.Card, error)
	Delete(userId, cardId string) error
	Update(userId, cardId string, input todo.UpdateCardInput) error
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
