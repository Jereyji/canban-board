package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
)

type BoardService struct {
	repo repository.Board
}

func NewBoardService(repo repository.Board) *BoardService {
	return &BoardService{repo: repo}
}

func (s *BoardService) Create(userId int, board todo.Board) (int, error) {
	return s.repo.Create(userId, board)
}

func (s *BoardService) GetAll(userId int) ([]todo.Board, error) {
	return s.repo.GetAll(userId)
}

func (s *BoardService) GetById(userId, boardId int) (todo.Board, error) {
	return s.repo.GetById(userId, boardId)
}