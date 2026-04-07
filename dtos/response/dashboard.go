package response

// logistics kpis
type LogisticsKPIResponse struct {
	TotalPhysicalItems int64 `json:"totalPhysicalItems"`
	TotalActiveBoxes   int64 `json:"totalActiveBoxes"`
	TotalUniqueDonors  int64 `json:"totalUniqueDonors"`
}

// DistributionItem representa una porción del gráfico de torta
type DistributionItem struct {
	Label string `json:"label"`
	Value int64  `json:"value"` // Cantidad física
}

type LogisticsDistributionResponse struct {
	BySize        []DistributionItem `json:"bySize"`
	ByCondition   []DistributionItem `json:"byCondition"`
	ByDisposition []DistributionItem `json:"byDisposition"`
	ByType        []DistributionItem `json:"byType"`
}

// financial kpis
type FinancialKPIResponse struct {
	TotalSold         float64 `json:"totalSold"`
	TotalDonatedValue float64 `json:"totalDonatedValue"`
	InventoryValue    float64 `json:"inventoryValue"`
}

type FinancialTrendItem struct {
	Period string  `json:"period"` // Formato "YYYY-MM" (Ej: "2026-04")
	Value  float64 `json:"value"`  // Dinero en ese período
}

type FinancialTrendResponse struct {
	SalesTrends    []FinancialTrendItem `json:"salesTrends"`
	DonationTrends []FinancialTrendItem `json:"donationTrends"`
}


// FinancialDistributionItem sirve para rankings financieros (Nombre -> Dinero)
type FinancialDistributionItem struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

// ProfitabilityResponse muestra qué tanto valor real se está generando
type ProfitabilityResponse struct {
	TotalRevenue float64 `json:"totalRevenue"` // Total Vendido
	TotalCost    float64 `json:"totalCost"`    // Costo base (DonationPrice) de lo que se vendió
	NetProfit    float64 `json:"netProfit"`    // Ganancia Neta (Revenue - Cost)
}