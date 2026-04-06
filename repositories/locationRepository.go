package repositories

import (
	"juanfeLogis/models"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) Create(location *models.Location) error {
	return r.db.Create(location).Error
}

func (r *LocationRepository) FindAll() ([]models.Location, error) {
	var locations []models.Location
	result := r.db.Find(&locations)
	return locations, result.Error
}

func (r *LocationRepository) GetByName(name string) (*models.Location, error) {
	var location models.Location

	result := r.db.Where("name = ?", name).First(&location)
	if result.Error != nil {
		return nil, result.Error
	}

	return &location, nil

}
func (r *LocationRepository) GetByID(id uint) (*models.Location, error) {
	var location models.Location
	result := r.db.First(&location, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &location, nil
}

func (r *LocationRepository) Update(location *models.Location) error {
	return r.db.Save(location).Error
}
