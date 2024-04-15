package repository

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

func NewRepository() *Repository {
	return &Repository{}
}