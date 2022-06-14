package builders

import (
	"fmt"

	models "github.com/echenim/patterns/common/models/bookstore"
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

func (b *BooksBuilder) SetAuthorID(id int) {
	b.AuthorID = id
}

func (b *BooksBuilder) SetPublisherID(id int) {
	b.PublisherID = id
}

func (b *BooksBuilder) SetGenreD(id int) {
	b.GenreID = id
}

func (b *BooksBuilder) SetTitle(title string) {
	b.Title = title
}

func (b *BooksBuilder) SetISBN(isbn string) {
	b.ISBN = isbn
}

func (b *BooksBuilder) SetType(types string) {
	b.Type = types
}

func (b *BooksBuilder) SetPublicationYear(publication_year int) {
	b.PublicationYear = publication_year
}

func (b *BooksBuilder) SetPrice(price float64) {
	b.Price = price
}

func (b *BooksBuilder) SetCondition(condition string) {
	b.Condition = condition
}

func (b *BooksBuilder) Build() (models.Books, error) {
	if b.Title == "" {
		return models.Books{}, fmt.Errorf("")
	}
	return models.Books{
		ID:              b.ID,
		AuthorID:        b.AuthorID,
		PublisherID:     b.PublisherID,
		Title:           b.Title,
		ISBN:            b.ISBN,
		GenreID:         b.GenreID,
		Type:            b.Type,
		PublicationYear: b.PublicationYear,
		Price:           b.Price,
		Condition:       b.Condition,
	}, nil
}
