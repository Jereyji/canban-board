package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
	
}

type Board interface {
	
}

type Card interface {

}

type Repository struct {
	Authorization
	Board
	Card
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}