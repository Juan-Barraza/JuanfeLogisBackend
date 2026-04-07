package repositories

import (
	"juanfeLogis/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) FindAllQuery() (*gorm.DB, []models.Product, error) {
	var products []models.Product
	result := r.db.Model(&models.Product{}).Preload("ProductType").Preload("Donor").Find(&products)
	return result, products, result.Error
}

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("ProductType").Preload("Donor").First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Delete(id string) error {
	return r.db.Unscoped().Delete(&models.Product{}, "id = ?", id).Error
}
