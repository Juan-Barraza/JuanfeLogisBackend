package repositories

import (
	"juanfeLogis/models"
	"strings"

	"gorm.io/gorm"
)

type DonorRepository struct {
	db *gorm.DB
}

func NewDonorRepository(db *gorm.DB) *DonorRepository {
	return &DonorRepository{db: db}
}

func (r *DonorRepository) Create(donor *models.Donor) error {
	return r.db.Create(donor).Error
}

func (r *DonorRepository) FindAll(name string) (*gorm.DB, []models.Donor, error) {
	var donors []models.Donor
	query := r.db.Model(&models.Donor{})
	if name != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%")
	}
	result := query.Find(&donors)
	return result, donors, result.Error
}

func (r *DonorRepository) GetByID(id string) (*models.Donor, error) {
	var donor models.Donor
	result := r.db.Where("id = ?", id).First(&donor)
	return &donor, result.Error
}

func (r *DonorRepository) GetByName(name string) (*models.Donor, error) {
	var donor models.Donor
	result := r.db.Where("lower(name) = lower(?)", name).First(&donor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &donor, nil
}

func (r *DonorRepository) Update(donor *models.Donor) error {
	return r.db.Save(donor).Error
}
