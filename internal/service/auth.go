package service

import (
	"os"
	"time"

	todo "github.com/Jereyji/canban-board"
	"github.com/Jereyji/canban-board/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json: "user_id"`
}

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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	tokenTTL := 12 * time.Hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	return token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
}

func generatePasswordHash(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
