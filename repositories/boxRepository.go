package repositories

import (
	"juanfeLogis/models"
	"strings"

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

func (r *BoxRepository) FindAllQuery(name string, location string, productId string) (*gorm.DB, []models.Box, error) {
	var boxes []models.Box
	query := r.db.Model(&models.Box{}).Preload("Location").Preload("Labels")

	name = strings.ToLower(name)
	location = strings.ToLower(location)

	if name != "" {
		query = query.Where("LOWER(boxes.name) LIKE ?", "%"+name+"%")
	}
	if location != "" {
		// Buscamos en la tabla relacionada de ubicaciones
		query = query.Joins("Join locations On locations.id = boxes.location_id").
			Where("LOWER(locations.name) LIKE ?", "%"+location+"%")
	}
	if productId != "" {
		// Subquery: solo cajas donde ese producto tiene stock > 0
		// Usamos IN en lugar de JOIN+Distinct para que Preload("Location") funcione correctamente
		subQuery := r.db.Model(&models.BoxStock{}).
			Select("box_id").
			Where("product_id = ? AND quantity > 0", productId)
		query = query.Where("boxes.id IN (?)", subQuery)
	}

	result := query.Find(&boxes)
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
	// Parseamos el id para usarlo en la limpieza de asociaciones
	boxID, err := uuid.Parse(id)
	if err == nil {
		// Eliminamos las relaciones en la tabla intermedia box_labels
		r.db.Model(&models.Box{ID: boxID}).Association("Labels").Clear()
	}

	return r.db.Delete(&models.Box{}, "id = ?", id).Error
}

func (r *BoxRepository) SetLabels(boxID string, labelIDs []uint) ([]models.ProductType, error) {
	var productTypes []models.ProductType
	if err := r.db.Where("id IN ?", labelIDs).Find(&productTypes).Error; err != nil {
		return nil, err
	}

	box := &models.Box{ID: uuid.MustParse(boxID)}
	err := r.db.Model(box).Association("Labels").Replace(productTypes)
	return productTypes, err
}
