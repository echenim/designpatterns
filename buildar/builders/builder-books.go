package builders

import (
	models "github.com/echenim/patterns/common/models"
)

type BooksBuilder struct {
	ID              int
	AuthorID        int
	PublisherID     int
	Title           string
	ISBN            string
	GenreID         int
	Type            string
	PublicationYear int
	Price           float64
	Condition       string
}

func NewBooksBuilder() *BooksBuilder {
	return &BooksBuilder{}
}

func (b *BooksBuilder) Build() (models.Books, error) {
	return models.Books{}, nil
}
