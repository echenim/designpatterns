package main

import (
	"fmt"

	builders "github.com/echenim/patterns/buildar/builders"
)

func main() {
	bidr := builders.NewBooksBuilder()
	bidr.ID = 1234567867543
	bidr.SetAuthorID(10000)
	bidr.SetPublisherID(32399293)
	bidr.SetGenreD(3023)
	bidr.SetTitle("Lord Of The Rings")
	bidr.SetISBN("A1234-GRA23")
	bidr.SetType("Adults")
	bidr.SetCondition("Used")
	bidr.SetPrice(304.60)
	bidr.SetPublicationYear(2004)

	book, er := bidr.Build()
	if er.Error() != "" {
		fmt.Println("Error creating the books:", er)
	} else {
		fmt.Printf("Books: %+v\n", book)
	}
}
