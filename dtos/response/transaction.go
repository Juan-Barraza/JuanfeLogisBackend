package response

import "time"

type TransactionResponse struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	ItemCount int       `json:"item_count"`
}

type TransactionItemResponse struct {
	ID           string  `json:"id"`
	ProductID    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	BoxID        string  `json:"box_id"`
	BoxName      string  `json:"box_name"`
	Quantity     int     `json:"quantity"`
	AppliedPrice float64 `json:"applied_price"`
}

type TransactionDetailResponse struct {
	ID        string                    `json:"id"`
	Type      string                    `json:"type"`
	UserName  string                    `json:"user_name"`
	CreatedAt time.Time                 `json:"created_at"`
	Items     []TransactionItemResponse `json:"items"`
}
