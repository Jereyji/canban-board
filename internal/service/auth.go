package service

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	enc, err := generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = enc
	return s.repo.CreateUser(user)
}

func generatePasswordHash(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}