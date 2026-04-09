package repositories

import (
	"juanfeLogis/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoxStockRepository struct {
	db *gorm.DB
}

func NewBoxStockRepository(db *gorm.DB) *BoxStockRepository {
	return &BoxStockRepository{db: db}
}

func (r *BoxStockRepository) FindByBoxAndProduct(boxID, productID uuid.UUID) (*models.BoxStock, error) {
	var stock models.BoxStock
	err := r.db.Where("box_id = ? AND product_id = ?", boxID, productID).First(&stock).Error
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *BoxStockRepository) Create(stock *models.BoxStock) error {
	return r.db.Create(stock).Error
}
func (r *BoxStockRepository) Update(stock *models.BoxStock) error {
	return r.db.Save(stock).Error
}

func (r *BoxStockRepository) Delete(stock *models.BoxStock) error {
	return r.db.Delete(stock).Error
}
func (r *BoxStockRepository) GetByBox(boxID uuid.UUID) ([]models.BoxStock, error) {
	var stocks []models.BoxStock
	err := r.db.Where("box_id = ?", boxID).Find(&stocks).Error
	return stocks, err
}

func (r *BoxStockRepository) GetByProduct(productID uuid.UUID) ([]models.BoxStock, error) {
	var stocks []models.BoxStock
	err := r.db.Where("product_id = ?", productID).Find(&stocks).Error
	return stocks, err
}

func (r *BoxStockRepository) GetTotalQuantityByProductID(productID uuid.UUID) (int, error) {
	var total int
	err := r.db.Model(&models.BoxStock{}).
		Where("product_id = ?", productID).
		Select("COALESCE(SUM(quantity), 0)").
		Row().Scan(&total)
	return total, err
}
