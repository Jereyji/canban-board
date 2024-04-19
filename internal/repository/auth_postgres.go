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

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := "INSERT INTO " + usersTable + " (name, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"

	row := r.db.QueryRow(query, user.Name, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}


