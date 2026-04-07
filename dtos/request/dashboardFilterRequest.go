package request

type DashboardFilterRequest struct {
	StartDate   string `query:"startDate"` // Formato YYYY-MM-DD
	EndDate     string `query:"endDate"`   // Formato YYYY-MM-DD
	Size        string `query:"size"`
	Condition   string `query:"condition"`
	Disposition string `query:"disposition"`
	ProductType uint   `query:"productType"`
	LocationID  uint   `query:"locationId"`
}
