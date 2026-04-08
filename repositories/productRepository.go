package repositories

import (
	"juanfeLogis/models"
	"strings"

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

func (r *ProductRepository) FindAllQuery(itemType, donor, size, disposition string) (*gorm.DB, []models.Product, error) {
	var products []models.Product
	query := r.db.Model(&models.Product{}).
		Preload("ProductType").
		Preload("Donor")

	if itemType != "" {
		query = query.Joins("Join product_types On product_types.id = products.product_type_id").
			Where("LOWER(product_types.name) LIKE ?", "%"+strings.ToLower(itemType)+"%")
	}

	if donor != "" {
		query = query.Joins("Join donors On donors.id = products.donor_id").
			Where("LOWER(donors.name) LIKE ?", "%"+strings.ToLower(donor)+"%")
	}

	if size != "" {
		query = query.Where("LOWER(products.size) LIKE ?", "%"+strings.ToLower(size)+"%")
	}

	if disposition != "" {
		query = query.Where("LOWER(products.disposition) = LOWER(?)", disposition)
	}

	result := query.Find(&products)
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
	return r.db.Delete(&models.Product{}, "id = ?", id).Error
}
