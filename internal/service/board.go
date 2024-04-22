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

func (s *BoardService) AddPermission(boardId, userId int, access string) error {
	return s.repo.AddPermission(boardId, userId, access)
}

func (s *BoardService) GetAll(userId int) ([]todo.Board, error) {
	return s.repo.GetAll(userId)
}

func (s *BoardService) GetById(userId, boardId int) (todo.Board, error) {
	return s.repo.GetById(userId, boardId)
}

func (s *BoardService) Delete(userId, boardId int) error {
	return s.repo.Delete(userId, boardId)
}

func (s *BoardService) Update(userId, boardId int, input todo.UpdateBoardInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, boardId, input)
}

func (s *BoardService) CheckPermissionToBoard(userId, boardId int, accessLevel string) error {
	return s.repo.CheckPermissionToBoard(userId, boardId, accessLevel)
}