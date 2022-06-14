package bookstores

type Books struct {
	ID              int     `json:"id"`
	AuthorID        int     `json:"author_id"`
	PublisherID     int     `json:"publisher_id"`
	Title           string  `json:"title"`
	ISBN            string  `json:"isbn"`
	GenreID         int     `json:"genre_id"`
	Type            string  `json:"type"`
	PublicationYear int     `json:"publication_year"`
	Price           float64 `json:"price"`
	Condition       string  `json:"condition"`
}
