package service

import (
	"errors"
	"os"
	"time"

	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (string, error) {
	enc, err := generatePasswordHash(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = enc
	return s.repo.CreateUser(user)
}

func (s *AuthService) CheckUser(email string) (string, error) {
	id, err := s.repo.CheckUser(email)
	if id == "" || err != nil {
		return "", errors.New("user with given email does not exist")
	}
	return id, nil
}

func (s *AuthService) ComparePassword(email, password string) error {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GenerateToken(email string) (string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return "", err
	}

	tokenTTL := 2 * time.Hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGNINGKEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *AuthService) GetById(userId string) (todo.User, error) {
	return s.repo.GetById(userId)
}

func (s *AuthService) UpdateUser(userId string, input todo.UpdateUserInput) error {
	return s.repo.UpdateUser(userId, input)
}

func (s *AuthService) GetAllUsers(boardId string) ([]todo.BoardUsers, error) {
	return s.repo.GetAllUsers(boardId)
}

func (s *AuthService) ExcludeUser(username, boardId string) error {
	return s.repo.ExcludeUser(username, boardId)
}

func (s *AuthService) SetCode(email, code string) error {
	return s.repo.SetCode(email, code)
}

func (s *AuthService) GetCode(email string) (string, error) {
	return s.repo.GetCode(email)
}