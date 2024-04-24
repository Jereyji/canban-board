package repository

import (
	"errors"
	"time"

	todo "github.com/Jereyji/canban-board"
	"github.com/jmoiron/sqlx"
)

type CardPostgres struct {
	db *sqlx.DB
}

func NewCardPostgres(db *sqlx.DB) *CardPostgres {
	return &CardPostgres{db: db}
}

func (r *CardPostgres) Create(boardId int, card todo.Card) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var dueDate time.Time
	if card.DueDate == "" {
		dueDate = time.Now().AddDate(0, 0, 7)
	} else {
		dueDate, err = time.Parse(time.RFC3339, card.DueDate)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	var cardId int
	err = tx.QueryRow(
		"INSERT INTO cards (title, description, due_date, user_id) VALUES ($1, $2, $3, $4) RETURNING id",
		card.Title,
		card.Description,
		dueDate,
		card.UserId,
	).Scan(&cardId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(
		"INSERT INTO board_cards (board_id, card_id, status_card) VALUES ($1, $2, $3)",
		boardId,
		cardId,
		todo.ToDo,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return cardId, tx.Commit()
}

func (r *CardPostgres) CheckPermissionToCard(userId, boardId int) error {
	var exists bool
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM "+boardPermissionsTable+
			" WHERE user_id = $1 AND board_id = $2  AND access_level != 'viewer'",
		userId,
		boardId,
	).Scan(&exists)

	if err != nil {
		return err
	}
	if !exists {
		return errors.New("you do not have the right to make changes in this board")
	}
	return nil
}

func (r *CardPostgres) GetAll(userId, boardId int) ([]todo.Card, error) {
	var cards []todo.Card

	query := "SELECT ct.id, ct.title, ct.description, ct.due_date, ct.user_id, ct.created_at" +
		" FROM " + cardsTable +
		" ct INNER JOIN " + boardCardsTable +
		" bc on bc.card_id = ct.id INNER JOIN " + boardPermissionsTable +
		" bp on bp.board_id = bc.board_id WHERE bc.board_id = $1 AND bp.user_id = $2"

	err := r.db.Select(&cards, query, boardId, userId); 
	if err != nil {
		return nil, err
	}

	return cards, nil
}
