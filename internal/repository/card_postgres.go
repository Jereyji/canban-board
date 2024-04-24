package repository

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func (r *CardPostgres) Create(boardId string, card todo.Card) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	var dueDate time.Time
	if card.DueDate == "" {
		dueDate = time.Now().AddDate(0, 0, 7)
	} else {
		dueDate, err = time.Parse(time.RFC3339, card.DueDate)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	var cardId string
	err = tx.QueryRow(
		"INSERT INTO cards (title, description, due_date, user_id) VALUES ($1, $2, $3, $4) RETURNING id",
		card.Title,
		card.Description,
		dueDate,
		card.UserId,
	).Scan(&cardId)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	_, err = tx.Exec(
		"INSERT INTO board_cards (board_id, card_id, status_card) VALUES ($1, $2, $3)",
		boardId,
		cardId,
		todo.ToDo,
	)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return cardId, tx.Commit()
}

func (r *CardPostgres) CheckPermissionToCard(userId, boardId string) error {
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

func (r *CardPostgres) GetAll(userId, boardId string) ([]todo.Card, error) {
	var cards []todo.Card

	query := "SELECT ct.id, ct.title, ct.description, ct.due_date, ct.user_id, ct.created_at" +
		" FROM " + cardsTable +
		" ct INNER JOIN " + boardCardsTable +
		" bc on bc.card_id = ct.id INNER JOIN " + boardPermissionsTable +
		" bp on bp.board_id = bc.board_id WHERE bc.board_id = $1 AND bp.user_id = $2"

	err := r.db.Select(&cards, query, boardId, userId)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *CardPostgres) GetById(userId, cardId string) (todo.Card, error) {
	var card todo.Card

	query := "SELECT ct.id, ct.title, ct.description, ct.due_date, ct.user_id, ct.created_at" +
		" FROM " + cardsTable +
		" ct INNER JOIN " + boardCardsTable +
		" bc on bc.card_id = ct.id INNER JOIN " + boardPermissionsTable +
		" bp on bp.board_id = bc.board_id WHERE bc.card_id = $1 AND bp.user_id = $2"

	err := r.db.Get(&card, query, cardId, userId)
	if err != nil {
		return card, err
	}

	return card, nil
}

func (r *CardPostgres) Delete(userId, cardId string) error {
	query := "DELETE FROM " + cardsTable + " ct USING " + boardCardsTable +
		" bc, " + boardPermissionsTable +
		" bp WHERE ct.id = bc.card_id AND bc.board_id = bp.board_id AND bp.user_id = $1 AND ct.id = $2"
	_, err := r.db.Exec(query, userId, cardId)

	return err
}

func (r *CardPostgres) Update(userId, cardId string, input todo.UpdateCardInput) error {
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

	if input.DueDate != nil {
		setValues = append(setValues, fmt.Sprintf("due_date=$%d", argId))
		args = append(args, *input.DueDate)
		argId++
	}

	if input.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *input.UserId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	argBoardIdStr := strconv.Itoa(argId)
	argUserIdStr := strconv.Itoa(argId + 1)
	query := "UPDATE " + cardsTable + " ct SET " + setQuery + " FROM " + boardCardsTable +
		" bc, " + boardPermissionsTable + " bp WHERE ct.id = bc.card_id AND bc.board_id = bp.board_id" + 
		" AND bp.user_id = $" + argUserIdStr + " AND ct.id = $" + argBoardIdStr
	args = append(args, cardId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}