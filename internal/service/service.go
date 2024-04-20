package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Board interface {
	Create(userId int, board todo.Board) (int, error)
	GetAll(userId int) ([]todo.Board, error)
	GetById(userId, boardId int) (todo.Board, error)
	Delete(userId, boardId int) error
}

type Card interface {
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
	}
}
