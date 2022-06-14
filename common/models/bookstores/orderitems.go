package bookstores

type OrderItem struct {
	ID       int     `json:"id"`
	BookID   int     `json:"book_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
