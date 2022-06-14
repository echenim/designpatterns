package bookstores

type Customers struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	State       string `json:"state"`
	PhoneNumber string `json:"phone_number"`
}
