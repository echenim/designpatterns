package bookstore

type Publishers struct {
	ID      int    `json:"id"`
	BookID  int    `json:"book"`
	Name    string `json:"name"`
	Country string `json:"country"`
}
