package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (string, error)
	ComparePassword(email, password string) error
	GenerateToken(email string) (string, error)
	ParseToken(accessToken string) (string, error)
	CheckUser(email string) (string, error)
	GetById(userId string) (todo.User, error)
	UpdateUser(userId string, input todo.UpdateUserInput) error
	GetAllUsers(boardId string) ([]todo.BoardUsers, error)
	ExcludeUser(userId, boardId string) error
	SetCode(email, code string) error
	GetCode(email string) (string, error)
}

type Board interface {
	Create(userId string, board todo.Board) (string, error)
	AddPermission(boardId, userId, access string) error
	GetAll(userId string) ([]todo.Board, error)
	GetById(userId, boardId string) (todo.Board, error)
	Delete(userId, boardId string) error
	Update(userId, boardId string, input todo.UpdateBoardInput) error
	CheckPermissionToBoard(username, boardId, accessLevel string) error
}

type Card interface {
	Create(userId, boardId string, card todo.Card) (string, error)
	CheckPermissionToCard(userId, boardId string) error
	GetAll(userId, boardId string) ([]todo.Card, error)
	GetById(userId, cardId string) (todo.Card, error)
	Delete(userId, cardId string) error
	Update(userId, cardId string, input todo.UpdateCardInput) error
	GetBoardIdByCard(cardId string) (string, error)
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
