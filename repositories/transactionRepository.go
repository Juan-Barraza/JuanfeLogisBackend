package repositories

import (
	"juanfeLogis/models"
	"strings"

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

func (r *TransactionRepository) FindAllQuery(txType, startDate, endDate string) (*gorm.DB, []models.Transaction, error) {
	var transactions []models.Transaction
	query := r.db.Model(&models.Transaction{}).Preload("User").Preload("Items")

	if txType != "" {
		query = query.Where("LOWER(type) = LOWER(?)", strings.ToLower(txType))
	}

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}

	if endDate != "" {
		query = query.Where("created_at <= ?", endDate+" 23:59:59")
	}

	result := query.Order("created_at DESC").Find(&transactions)
	return result, transactions, result.Error
}

func (r *TransactionRepository) GetByID(id string) (*models.Transaction, error) {
	var tx models.Transaction
	err := r.db.Preload("User").
		Preload("Items.Product.ProductType").
		Preload("Items.Box").
		First(&tx, "id = ?", id).Error
	return &tx, err
}
