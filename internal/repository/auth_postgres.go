package repository

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (string, error) {
	var id string
	query := "INSERT INTO " + usersTable + " (name, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"

	row := r.db.QueryRow(query, user.Name, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (todo.User, error) {
	var user todo.User
	query := "SELECT id, password_hash FROM " + usersTable + " WHERE username=$1"
	err := r.db.Get(&user, query, username)
	return user, err
}

func (r *AuthPostgres) CheckUser(email string) (string, error) {
	var userId string
	query := "SELECT id FROM " + usersTable + " WHERE email = $1"
	err := r.db.QueryRow(query, email).Scan(&userId)
	return userId, err
}

func (r *AuthPostgres) GetById(userId string) (todo.User, error) {
	var user todo.User

	query := "SELECT name, username, email, created_at FROM " + usersTable + " WHERE id = $1"
	err := r.db.Get(&user, query, userId)
	return user, err
}