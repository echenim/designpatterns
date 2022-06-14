package builders

import (
	"github.com/echenim/designpatterns/common/models/bookstores"
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

func (b *BooksBuilder) Build() (Books, error) {
	return Books{}, nil
}
