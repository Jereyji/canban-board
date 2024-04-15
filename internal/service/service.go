package service

import "github.com/Jereyji/canban-board/internal/repository"

type Authorization interface {
	
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
	return &Service{}
}