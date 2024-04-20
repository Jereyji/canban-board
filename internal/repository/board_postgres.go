package repository

import (
	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
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