package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Board interface {

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
	}
}