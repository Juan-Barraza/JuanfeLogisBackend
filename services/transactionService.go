package services

import (
	"errors"
	"juanfeLogis/dtos/response"
	"juanfeLogis/models"
	"juanfeLogis/repositories"
	"juanfeLogis/utils"
)

type TransactionService struct {
	txRepo         *repositories.TransactionRepository
	paginationRepo *repositories.PaginationRepository
}

func NewTransactionService(txRepo *repositories.TransactionRepository, paginationRepo *repositories.PaginationRepository) *TransactionService {
	return &TransactionService{
		txRepo:         txRepo,
		paginationRepo: paginationRepo,
	}
}

func (s *TransactionService) GetAllTransactions(page, pageSize int, txType, startDate, endDate string) (*utils.Pagination, error) {
	query, _, err := s.txRepo.FindAllQuery(txType, startDate, endDate)
	if err != nil {
		return nil, errors.New("error al consultar las transacciones")
	}

	var transactions []models.Transaction
	pagination := &utils.Pagination{Page: page, PageSize: pageSize}
	paginationResult, err := s.paginationRepo.GetPaginatedResults(query, pagination, &transactions)
	if err != nil {
		return nil, errors.New("error al paginar los resultados")
	}

	res := make([]response.TransactionResponse, len(transactions))
	for i, tx := range transactions {
		res[i] = response.TransactionResponse{
			ID:        tx.ID.String(),
			Type:      tx.Type,
			UserName:  tx.User.Name,
			CreatedAt: tx.CreatedAt,
			ItemCount: len(tx.Items),
		}
	}

	paginationResult.Data = res
	return paginationResult, nil
}

func (s *TransactionService) GetTransactionByID(id string) (*response.TransactionDetailResponse, error) {
	tx, err := s.txRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("transacción no encontrada")
	}

	items := make([]response.TransactionItemResponse, len(tx.Items))
	for i, item := range tx.Items {
		items[i] = response.TransactionItemResponse{
			ID:           item.ID.String(),
			ProductID:    item.ProductID.String(),
			ProductName:  item.Product.ProductType.Name,
			BoxID:        item.BoxID.String(),
			BoxName:      item.Box.Name,
			Quantity:     item.Quantity,
			AppliedPrice: item.AppliedPrice,
		}
	}

	return &response.TransactionDetailResponse{
		ID:        tx.ID.String(),
		Type:      tx.Type,
		UserName:  tx.User.Name,
		CreatedAt: tx.CreatedAt,
		Items:     items,
	}, nil
}
