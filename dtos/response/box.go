package response

import "time"

type BoxResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	LocationID   uint      `json:"location_id"`
	LocationName string    `json:"location_name"`
	Labels       []string  `json:"labels"`
	QRCodeURL    string    `json:"qr_code_url"`
	CreatedAt    time.Time `json:"created_at"`
}

// DTO para el detalle de la caja con su inventario
type BoxStockItem struct {
	ProductID         string  `json:"product_id"`
	ProductTypeName   string  `json:"product_type_name"`
	DonorName         string  `json:"donor_name"`
	Quantity          int     `json:"quantity"`
	DonationPrice     float64 `json:"donation_price"`
	SalePrice         float64 `json:"sale_price"`
	PhysicalCondition string  `json:"physical_condition"`
}

type BoxDetailResponse struct {
	Box      BoxResponse    `json:"box"`
	Products []BoxStockItem `json:"products"`
}
