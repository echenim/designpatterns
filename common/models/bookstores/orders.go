package bookstores

import "time"

type Orders struct {
	ID         int       `json:"id"`
	ItemID     int       `json:"item_id"`
	CustomerID int       `json:"customer_id"`
	OrderDate  time.Time `json:"order_date"`
	Shipping   time.Time `json:"shipping"`
	Total      int       `json:"total"`
}
