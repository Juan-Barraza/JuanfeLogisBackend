package repositories

import (
	"juanfeLogis/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoxRepository struct {
	db *gorm.DB
}

func NewBoxRepository(db *gorm.DB) *BoxRepository {
	return &BoxRepository{db: db}
}

func (r *BoxRepository) Create(box *models.Box) error {
	return r.db.Create(box).Error
}

func (r *BoxRepository) Update(box *models.Box) error {
	return r.db.Save(box).Error
}

func (r *BoxRepository) FindAllQuery() (*gorm.DB, []models.Box, error) {
	var boxes []models.Box
	result := r.db.Model(&models.Box{}).Preload("Location").Preload("Labels").Find(&boxes)
	return result, boxes, result.Error
}

func (r *BoxRepository) GetBoxWithStock(boxID string) (*models.Box, []models.BoxStock, error) {
	var box models.Box
	if err := r.db.Preload("Location").Preload("Labels").First(&box, "id = ?", boxID).Error; err != nil {
		return nil, nil, err
	}

	var stock []models.BoxStock
	// Preloads anidados para traer toda la info del producto desde la tabla pivote
	err := r.db.Preload("Product.ProductType").
		Preload("Product.Donor").
		Where("box_id = ?", boxID).
		Find(&stock).Error

	return &box, stock, err
}

func (r *BoxRepository) GetByID(id string) (*models.Box, error) {
	var box models.Box
	err := r.db.Preload("Location").Preload("Labels").First(&box, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &box, nil
}

func (r *BoxRepository) Delete(id string) error {
	return r.db.Unscoped().Delete(&models.Box{}, "id = ?", id).Error
}

func (r *BoxRepository) SetLabels(boxID string, labelIDs []uint) error {
	var productTypes []models.ProductType
	if err := r.db.Where("id IN ?", labelIDs).Find(&productTypes).Error; err != nil {
		return err
	}

	box := &models.Box{ID: uuid.MustParse(boxID)}
	return r.db.Model(box).Association("Labels").Replace(productTypes)
}
