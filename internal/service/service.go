package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
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
	Create(userId, boardId int, card todo.Card) (int, error)
	CheckPermissionToCard(userId, boardId int) error
}

type Service struct {
	Authorization
	Board
	Card
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Board:         NewBoardService(repos.Board),
		Card:          NewCardService(repos.Card, repos.Board),
	}
}
