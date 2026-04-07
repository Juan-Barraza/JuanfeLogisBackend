package repositories

import (
	"juanfeLogis/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(tx *models.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *TransactionRepository) CreateTransactionItem(item *models.TransactionItem) error {
	return r.db.Create(item).Error
}

func (r *TransactionRepository) GetByProduct(productID uuid.UUID) ([]models.TransactionItem, error) {
	var items []models.TransactionItem
	err := r.db.Where("product_id = ?", productID).Find(&items).Error
	return items, err
}

func (r *TransactionRepository) GetByBox(boxID uuid.UUID) ([]models.TransactionItem, error) {
	var items []models.TransactionItem
	err := r.db.Where("box_id = ?", boxID).Find(&items).Error
	return items, err
}
