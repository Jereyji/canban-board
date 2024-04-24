package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
)

type CardService struct {
	repo repository.Card
	boardRepo repository.Board
}

func NewCardService(repo repository.Card, boardRepo repository.Board) *CardService {
	return &CardService{repo: repo, boardRepo: boardRepo}
}

func (s *CardService) Create(userId, boardId int, card todo.Card) (int, error) {
	return s.repo.Create(boardId, card)
}

func (s *CardService) CheckPermissionToCard(userId, boardId int) error {
	return s.repo.CheckPermissionToCard(userId, boardId)
}

func (s *CardService) GetAll(userId, boardId int) ([]todo.Card, error) {
	return s.repo.GetAll(userId, boardId)
}

func (s *CardService) GetById(userId, cardId int) (todo.Card, error) {
	return s.repo.GetById(userId, cardId)
}