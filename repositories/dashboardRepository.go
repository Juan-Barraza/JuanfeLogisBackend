package repositories

import (
	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}

// Scope privado para reutilizar los filtros en varias consultas
func applyProductFilters(req request.DashboardFilterRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.Size != "" {
			db = db.Where("products.size = ?", req.Size)
		}
		if req.Condition != "" {
			db = db.Where("products.physical_condition = ?", req.Condition)
		}
		if req.Disposition != "" {
			db = db.Where("products.disposition = ?", req.Disposition)
		}
		if req.ProductType != 0 {
			db = db.Where("products.product_type_id = ?", req.ProductType)
		}
		return db
	}
}

func (r *DashboardRepository) GetLogisticsKPIs(req request.DashboardFilterRequest) (response.LogisticsKPIResponse, error) {
	var kpis response.LogisticsKPIResponse

	// 1. Total Físico
	err := r.db.Table("box_stocks").
		Joins("JOIN products ON products.id = box_stocks.product_id").
		Scopes(applyProductFilters(req)).
		Select("COALESCE(SUM(box_stocks.quantity), 0)").
		Scan(&kpis.TotalPhysicalItems).Error
	if err != nil {
		return kpis, err
	}

	// 2. Cajas Activas
	err = r.db.Table("box_stocks").
		Joins("JOIN products ON products.id = box_stocks.product_id").
		Scopes(applyProductFilters(req)).
		Where("box_stocks.quantity > 0").
		Distinct("box_stocks.box_id").
		Count(&kpis.TotalActiveBoxes).Error
	if err != nil {
		return kpis, err
	}

	// 3. Donantes Únicos
	err = r.db.Table("products").
		Scopes(applyProductFilters(req)).
		Distinct("products.donor_id").
		Count(&kpis.TotalUniqueDonors).Error

	return kpis, err
}

// GetLogisticsDistribution obtiene las cantidades agrupadas por categoría
func (r *DashboardRepository) GetLogisticsDistribution(req request.DashboardFilterRequest) (response.LogisticsDistributionResponse, error) {
	var result response.LogisticsDistributionResponse

	// Base query reutilizable (JOIN box_stocks con products)
	baseQuery := func() *gorm.DB {
		return r.db.Table("box_stocks").
			Joins("JOIN products ON products.id = box_stocks.product_id").
			Scopes(applyProductFilters(req)).
			Where("box_stocks.quantity > 0")
	}

	// 1. Agrupado por Talla (Size)
	baseQuery().Select("products.size as label, SUM(box_stocks.quantity) as value").
		Group("products.size").
		Scan(&result.BySize)

	// 2. Agrupado por Condición Física
	baseQuery().Select("products.physical_condition as label, SUM(box_stocks.quantity) as value").
		Group("products.physical_condition").
		Scan(&result.ByCondition)

	// 3. Agrupado por Disposición
	baseQuery().Select("products.disposition as label, SUM(box_stocks.quantity) as value").
		Group("products.disposition").
		Scan(&result.ByDisposition)

	// 4. Agrupado por Tipo de Producto (Hacemos JOIN extra para obtener el nombre del tipo)
	baseQuery().Joins("JOIN product_types ON product_types.id = products.product_type_id").
		Select("product_types.name as label, SUM(box_stocks.quantity) as value").
		Group("product_types.name").
		Scan(&result.ByType)

	return result, nil
}

func (r *DashboardRepository) GetLocationDistribution(req request.DashboardFilterRequest) ([]response.DistributionItem, error) {
	var result []response.DistributionItem

	query := r.db.Table("box_stocks").
		Joins("JOIN products ON products.id = box_stocks.product_id").
		Joins("JOIN boxes ON boxes.id = box_stocks.box_id").
		Joins("JOIN locations ON locations.id = boxes.location_id").
		Scopes(applyProductFilters(req)).
		Where("box_stocks.quantity > 0")

	// Si envían una sede específica en el filtro
	if req.LocationID != 0 {
		query = query.Where("boxes.location_id = ?", req.LocationID)
	}

	err := query.Select("locations.name as label, COUNT(DISTINCT boxes.id) as value").
		Group("locations.id, locations.name").
		Scan(&result).Error

	return result, err
}

func (r *DashboardRepository) GetTopDonorsLogistics(req request.DashboardFilterRequest) ([]response.DistributionItem, error) {
	var result []response.DistributionItem

	err := r.db.Table("box_stocks").
		Joins("JOIN products ON products.id = box_stocks.product_id").
		Joins("JOIN donors ON donors.id = products.donor_id").
		Scopes(applyProductFilters(req)).
		Where("box_stocks.quantity > 0").
		Select("donors.name as label, SUM(box_stocks.quantity) as value"). // Sumamos cantidad física
		Group("donors.id, donors.name").
		Order("value DESC"). // Los que más tienen de primero
		Limit(10).           // Solo el Top 10
		Scan(&result).Error

	return result, err
}

// Nuevo Scope para filtrar por fechas en las transacciones
func applyDateFilters(req request.DashboardFilterRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.StartDate != "" {
			// Asumiendo formato YYYY-MM-DD
			db = db.Where("transactions.created_at >= ?", req.StartDate+" 00:00:00")
		}
		if req.EndDate != "" {
			db = db.Where("transactions.created_at <= ?", req.EndDate+" 23:59:59")
		}
		return db
	}
}

// Método Principal
func (r *DashboardRepository) GetFinancialKPIs(req request.DashboardFilterRequest) (response.FinancialKPIResponse, error) {
	var kpis response.FinancialKPIResponse

	// 1. Total Vendido (Suma de Precio Aplicado * Cantidad en transacciones tipo 'salida')
	err := r.db.Table("transaction_items").
		Joins("JOIN transactions ON transactions.id = transaction_items.transaction_id").
		Scopes(applyDateFilters(req)).
		Where("transactions.type = ?", "salida").
		Select("COALESCE(SUM(transaction_items.applied_price * transaction_items.quantity), 0)").
		Scan(&kpis.TotalSold).Error
	if err != nil {
		return kpis, err
	}

	// 2. Total Donado (Valor de entrada, recordamos que pusimos el DonationPrice como AppliedPrice en el AddStock)
	err = r.db.Table("transaction_items").
		Joins("JOIN transactions ON transactions.id = transaction_items.transaction_id").
		Scopes(applyDateFilters(req)).
		Where("transactions.type = ?", "entrada").
		Select("COALESCE(SUM(transaction_items.applied_price * transaction_items.quantity), 0)").
		Scan(&kpis.TotalDonatedValue).Error
	if err != nil {
		return kpis, err
	}

	// 3. Valorización del Inventario (Stock actual * SalePrice del producto)
	// Nota: Esto no usa filtro de fechas, usa el filtro de productos del dashboard logístico
	err = r.db.Table("box_stocks").
		Joins("JOIN products ON products.id = box_stocks.product_id").
		Scopes(applyProductFilters(req)). // Reutilizamos el scope logístico!
		Where("box_stocks.quantity > 0").
		Select("COALESCE(SUM(box_stocks.quantity * products.sale_price), 0)").
		Scan(&kpis.InventoryValue).Error

	return kpis, err
}

func (r *DashboardRepository) GetFinancialTrends(req request.DashboardFilterRequest) (response.FinancialTrendResponse, error) {
	var result response.FinancialTrendResponse

	// 1. Tendencia de Ventas (Salidas) agrupadas por mes
	err := r.db.Table("transaction_items").
		Joins("JOIN transactions ON transactions.id = transaction_items.transaction_id").
		Scopes(applyDateFilters(req)).
		Where("transactions.type = ?", "salida").
		// Usamos TO_CHAR para extraer el Año-Mes en PostgreSQL
		Select("TO_CHAR(transactions.created_at, 'YYYY-MM') as period, COALESCE(SUM(transaction_items.applied_price * transaction_items.quantity), 0) as value").
		Group("TO_CHAR(transactions.created_at, 'YYYY-MM')").
		Order("period ASC"). // Ordenamos cronológicamente
		Scan(&result.SalesTrends).Error
	if err != nil {
		return result, err
	}

	// 2. Tendencia de Donaciones Recibidas (Entradas) agrupadas por mes
	err = r.db.Table("transaction_items").
		Joins("JOIN transactions ON transactions.id = transaction_items.transaction_id").
		Scopes(applyDateFilters(req)).
		Where("transactions.type = ?", "entrada").
		Select("TO_CHAR(transactions.created_at, 'YYYY-MM') as period, COALESCE(SUM(transaction_items.applied_price * transaction_items.quantity), 0) as value").
		Group("TO_CHAR(transactions.created_at, 'YYYY-MM')").
		Order("period ASC").
		Scan(&result.DonationTrends).Error

	return result, err
}

func (r *DashboardRepository) GetTopDonorsFinancial(req request.DashboardFilterRequest) ([]response.FinancialDistributionItem, error) {
	var result []response.FinancialDistributionItem

	err := r.db.Table("transaction_items").
		Joins("JOIN transactions ON transactions.id = transaction_items.transaction_id").
		Joins("JOIN products ON products.id = transaction_items.product_id").
		Joins("JOIN donors ON donors.id = products.donor_id").
		Scopes(applyDateFilters(req)).
		Where("transactions.type = ?", "entrada"). // Solo entradas (donaciones reales)
		Select("donors.name as label, COALESCE(SUM(transaction_items.applied_price * transaction_items.quantity), 0) as value").
		Group("donors.id, donors.name").
		Order("value DESC").
		Limit(10).
		Scan(&result).Error

	return result, err
}

// Rentabilidad (Valor Generado)
func (r *DashboardRepository) GetProfitability(req request.DashboardFilterRequest) (response.ProfitabilityResponse, error) {
	var result response.ProfitabilityResponse

	// Calculamos los ingresos (Revenue) y el costo (Cost) en una sola consulta
	err := r.db.Table("transaction_items").
		Joins("JOIN transactions ON transactions.id = transaction_items.transaction_id").
		Joins("JOIN products ON products.id = transaction_items.product_id").
		Scopes(applyDateFilters(req)).
		Where("transactions.type = ?", "salida"). // Analizamos solo lo que ya se vendió/salió
		Select(`
			COALESCE(SUM(transaction_items.applied_price * transaction_items.quantity), 0) as total_revenue,
			COALESCE(SUM(products.donation_price * transaction_items.quantity), 0) as total_cost
		`).
		Scan(&result).Error

	if err != nil {
		return result, err
	}

	// Calculamos la ganancia neta
	result.NetProfit = result.TotalRevenue - result.TotalCost

	return result, nil
}
