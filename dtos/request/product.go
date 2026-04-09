package request

type ProductRequest struct {
	ProductTypeID     uint    `json:"product_type_id"`
	DonorID           string  `json:"donor_id"`
	Size              string  `json:"size"`
	DonationPrice     float64 `json:"donation_price"`
	SalePrice         float64 `json:"sale_price"`
	PhysicalCondition string  `json:"physical_condition"`
	Disposition       string  `json:"disposition"`
	Description       string  `json:"description"`
}
