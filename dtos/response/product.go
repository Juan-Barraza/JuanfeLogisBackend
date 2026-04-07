package response

import "time"

type ProductResponse struct {
	ID                string    `json:"id"`
	ProductTypeID     uint      `json:"product_type_id"`
	DonorID           string    `json:"donor_id"`
	ProductTypeName   string    `json:"product_type_name"`
	DonorName         string    `json:"donor_name"`
	Size              string    `json:"size"`
	DonationPrice     float64   `json:"donation_price"`
	SalePrice         float64   `json:"sale_price"`
	PhysicalCondition string    `json:"physical_condition"`
	Disposition       string    `json:"disposition"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
