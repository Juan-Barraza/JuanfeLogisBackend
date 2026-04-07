package request

type BoxStockRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
