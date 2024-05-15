package repository

import (
	"fmt"
	"strconv"
	"strings"

	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *AuthPostgres) GetUser(email string) (todo.User, error) {
	var user todo.User
	query := "SELECT id, password_hash FROM " + usersTable + " WHERE email=$1"
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *AuthPostgres) CheckUser(email string) (string, error) {
	var userId string
	query := "SELECT id FROM " + usersTable + " WHERE email = $1"
	err := r.db.QueryRow(query, email).Scan(&userId)
	return userId, err
}

func (r *AuthPostgres) UpdateUser(userId string, input todo.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, input.Username)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, input.Email)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := "UPDATE " + usersTable + " SET " + setQuery + " WHERE id = $" + strconv.Itoa(argId)
	args = append(args, userId)

	logrus.Debugf("updateUserQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *AuthPostgres) GetAllUsers(boardId string) ([]todo.BoardUsers, error) {
	var Users []todo.BoardUsers

	query := `
        SELECT u.username, bp.access_level
        FROM users u
        JOIN board_permissions bp ON u.id = bp.user_id
        WHERE bp.board_id = $1`

	err := r.db.Select(&Users, query, boardId)
	if err != nil {
		return nil, err
	}

	return Users, nil
}

func (r *AuthPostgres) GetById(userId string) (todo.User, error) {
	var user todo.User

	query := "SELECT name, username, email, created_at FROM " + usersTable + " WHERE id = $1"
	err := r.db.Get(&user, query, userId)
	return user, err
}

func (r *AuthPostgres) ExcludeUser(username, boardId string) error {
	query := "DELETE FROM " + boardPermissionsTable +
		" WHERE user_id IN (SELECT id FROM " + usersTable + " WHERE username = $1) AND board_id = $2"
	_, err := r.db.Exec(query, username, boardId)

	return err
}

func (r *AuthPostgres) SetCode(email, code string) error {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM confirmation_codes WHERE email = $1;", email).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		query := "UPDATE confirmation_codes SET code = $1 WHERE email = $2;"
		_, err = r.db.Exec(query, code, email)
	} else {
		query := "INSERT INTO confirmation_codes (email, code) VALUES ($1, $2);"
		_, err = r.db.Exec(query, email, code)
	}
	return err
}

func (r *AuthPostgres) GetCode(email string) (string, error) {
	var code string

	err := r.db.QueryRow("SELECT code FROM confirmation_codes WHERE email = $1 LIMIT 1;", email).Scan(&code)
	if err != nil {
		return code, err
	}

	return code, nil
}

