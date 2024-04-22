package repository

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BoardPostgres struct {
	db *sqlx.DB
}

func NewBoardPostgres(db *sqlx.DB) *BoardPostgres {
	return &BoardPostgres{db: db}
}

func (r *BoardPostgres) Create(userId int, board todo.Board) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBoardQuery := "INSERT INTO " + boardsTable + " (title, description) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRow(createBoardQuery, board.Title, board.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserBoardQuery := "INSERT INTO " + boardPermissionsTable + " (user_id, board_id, access_level) VALUES ($1, $2, 'admin')"
	_, err = tx.Exec(createUserBoardQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BoardPostgres) AddPermission(boardId, userId int, access string) error {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM "+boardPermissionsTable+" WHERE board_id = $1 AND user_id = $2", boardId, userId).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("this user already has access to the board")
	}

	query := "INSERT INTO " + boardPermissionsTable + " (board_id, user_id, access_level) VALUES ($1, $2, $3)"
	_, err = r.db.Exec(query, boardId, userId, access)
	return err
}

func (r *BoardPostgres) GetAll(userId int) ([]todo.Board, error) {
	var boards []todo.Board
	query := "SELECT bt.id, bt.title, bt.description, bt.created_at FROM " + boardsTable + " bt INNER JOIN " +
		boardPermissionsTable + " bu on bt.id = bu.board_id WHERE bu.user_id = $1"

	err := r.db.Select(&boards, query, userId)
	return boards, err
}

func (r *BoardPostgres) GetById(userId, boardId int) (todo.Board, error) {
	var board todo.Board
	query := "SELECT bt.id, bt.title, bt.description, bt.created_at FROM " + boardsTable + " bt INNER JOIN " +
		boardPermissionsTable + " bu on bt.id = bu.board_id WHERE bu.user_id = $1 AND bu.board_id = $2"

	err := r.db.Get(&board, query, userId, boardId)
	return board, err
}

func (r *BoardPostgres) Delete(userId, boardId int) error {
	query := "DELETE FROM " + boardsTable + " bt USING " + boardPermissionsTable +
		" bu WHERE bt.id = bu.board_id AND bu.user_id = $1 AND bu.board_id = $2"
	_, err := r.db.Exec(query, userId, boardId)

	return err
}

func (r *BoardPostgres) Update(userId, boardId int, input todo.UpdateBoardInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	argBoardIdStr := strconv.Itoa(argId)
	argUserIdStr := strconv.Itoa(argId + 1)
	setQuery := strings.Join(setValues, ", ")
	query := "UPDATE " + boardsTable + " bt SET " + setQuery + " FROM " + boardPermissionsTable +
		" bu WHERE bt.id = bu.board_id AND bu.board_id = $" + argBoardIdStr + " AND bu.user_id = $" + argUserIdStr
	args = append(args, boardId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BoardPostgres) CheckPermission(ownerId, boardId int, accessLevel string) error {
	var exists bool
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM "+boardPermissionsTable+
		" WHERE user_id = $1 AND board_id = $2  AND access_level = $3", 
		ownerId,
		boardId, 
		accessLevel,
	).Scan(&exists)
	
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("you do not have the right to make changes to this board")
	}
	return nil
}
