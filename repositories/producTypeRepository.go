package repositories

import (
	"juanfeLogis/models"

	"gorm.io/gorm"
)

type ProductTypeRepository struct {
	db *gorm.DB
}

func NewProductTypeRepository(db *gorm.DB) *ProductTypeRepository {
	return &ProductTypeRepository{db: db}
}

func (r *ProductTypeRepository) Create(pt *models.ProductType) error {
	return r.db.Create(pt).Error
}

func (r *ProductTypeRepository) FindAll() ([]models.ProductType, error) {
	var productTypes []models.ProductType
	result := r.db.Find(&productTypes)
	return productTypes, result.Error
}

func (r *ProductTypeRepository) GetByID(id uint) (*models.ProductType, error) {
	var pt models.ProductType
	result := r.db.First(&pt, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pt, nil
}

func (r *ProductTypeRepository) GetByName(name string) (*models.ProductType, error) {
	var pt models.ProductType
	result := r.db.Where("name = ?", name).First(&pt)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pt, nil
}

func (r *ProductTypeRepository) Update(pt *models.ProductType) error {
	return r.db.Save(pt).Error
}
