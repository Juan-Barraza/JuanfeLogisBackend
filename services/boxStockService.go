package services

import (
	"errors"
	"juanfeLogis/dtos/request"
	"juanfeLogis/models"
	"juanfeLogis/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoxStockService struct {
	db              *gorm.DB
	stockRepo       *repositories.BoxStockRepository
	transactionRepo *repositories.TransactionRepository
	productRepo     *repositories.ProductRepository
}

func NewBoxStockService(
	db *gorm.DB,
	stockRepo *repositories.BoxStockRepository,
	transactionRepo *repositories.TransactionRepository,
	productRepo *repositories.ProductRepository,
) *BoxStockService {
	return &BoxStockService{
		db:              db,
		stockRepo:       stockRepo,
		transactionRepo: transactionRepo,
		productRepo:     productRepo,
	}
}

func (s *BoxStockService) AddStock(boxID uuid.UUID, req request.BoxStockRequest, userID uuid.UUID) error {
	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return errors.New("product_id inválido")
	}

	// Iniciamos la transacción segura
	return s.db.Transaction(func(tx *gorm.DB) error {
		// A. Obtener el producto para saber su Precio de Donación
		var product models.Product
		if err := tx.First(&product, productID).Error; err != nil {
			return errors.New("producto no encontrado")
		}

		// B. Lógica de BoxStock
		var stock models.BoxStock
		err := tx.Where("box_id = ? AND product_id = ?", boxID, productID).First(&stock).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// No existe, creamos
				newStock := &models.BoxStock{
					BoxID:     boxID,
					ProductID: productID,
					Quantity:  req.Quantity,
				}
				if err := tx.Create(newStock).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// Ya existe, sumamos
			stock.Quantity += req.Quantity
			if err := tx.Save(&stock).Error; err != nil {
				return err
			}
		}

		// C. Registrar en el Historial usando DonationPrice
		return s.registerTransactionTx(tx, "entrada", userID, productID, boxID, req.Quantity, product.DonationPrice, nil)
	})
}

func (s *BoxStockService) RemoveStock(boxID uuid.UUID, req request.BoxStockRequest, userID uuid.UUID) error {
	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return errors.New("product_id inválido")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// A. Obtener el producto para saber su Precio de Venta
		var product models.Product
		if err := tx.First(&product, productID).Error; err != nil {
			return errors.New("producto no encontrado")
		}

		// B. Lógica de BoxStock
		var stock models.BoxStock
		err := tx.Where("box_id = ? AND product_id = ?", boxID, productID).First(&stock).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("el producto no existe en esta caja")
			}
			return err
		}

		if stock.Quantity == 0 {
			return errors.New("el stock de este producto en esta caja ya está en 0")
		}
		if stock.Quantity < req.Quantity {
			return errors.New("cantidad a remover supera el stock disponible")
		}

		// Restamos
		stock.Quantity -= req.Quantity
		if err := tx.Save(&stock).Error; err != nil {
			return err
		}

		// C. Registrar en el Historial usando SalePrice
		return s.registerTransactionTx(tx, "salida", userID, productID, boxID, req.Quantity, product.SalePrice, req.Destination)
	})
}

func (s *BoxStockService) ReturnStock(boxID uuid.UUID, req request.BoxStockRequest, userID uuid.UUID) error {
	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return errors.New("product_id inválido")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// A. Obtener el producto para saber su Precio de Venta (se devuelve al mismo valor que se vendió)
		var product models.Product
		if err := tx.First(&product, productID).Error; err != nil {
			return errors.New("producto no encontrado")
		}

		// B. Lógica de BoxStock
		var stock models.BoxStock
		err := tx.Where("box_id = ? AND product_id = ?", boxID, productID).First(&stock).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("el producto no existe en esta caja para hacer una devolución")
			}
			return err
		}

		// Sumamos de vuelta
		stock.Quantity += req.Quantity
		if err := tx.Save(&stock).Error; err != nil {
			return err
		}

		// C. Registrar en el Historial usando SalePrice
		return s.registerTransactionTx(tx, "devolucion", userID, productID, boxID, req.Quantity, product.SalePrice, nil)
	})
}

func (s *BoxStockService) registerTransactionTx(
	tx *gorm.DB,
	txType string,
	userID, productID, boxID uuid.UUID,
	quantity int,
	appliedPrice float64,
	destination *string,
) error {

	transaction := &models.Transaction{
		ID:     uuid.New(),
		Type:   txType,
		UserID: userID,
	}
	if err := tx.Create(transaction).Error; err != nil {
		return err
	}

	item := &models.TransactionItem{
		ID:            uuid.New(),
		TransactionID: transaction.ID,
		ProductID:     productID,
		BoxID:         boxID,
		Quantity:      quantity,
		Destination:   destination,
		AppliedPrice:  appliedPrice,
	}
	return tx.Create(item).Error
}
