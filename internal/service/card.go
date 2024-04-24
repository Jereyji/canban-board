package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
)

type CardService struct {
	repo      repository.Card
	boardRepo repository.Board
}

func NewCardService(repo repository.Card, boardRepo repository.Board) *CardService {
	return &CardService{repo: repo, boardRepo: boardRepo}
}

func (s *CardService) Create(userId, boardId string, card todo.Card) (string, error) {
	return s.repo.Create(boardId, card)
}

func (s *CardService) CheckPermissionToCard(userId, boardId string) error {
	return s.repo.CheckPermissionToCard(userId, boardId)
}

func (s *CardService) GetAll(userId, boardId string) ([]todo.Card, error) {
	return s.repo.GetAll(userId, boardId)
}

func (s *CardService) GetById(userId, cardId string) (todo.Card, error) {
	return s.repo.GetById(userId, cardId)
}

func (s *CardService) Delete(userId, cardId string) error {
	return s.repo.Delete(userId, cardId)
}

func (s *CardService) Update(userId, cardId string, input todo.UpdateCardInput) error {
	return s.repo.Update(userId, cardId, input)
}