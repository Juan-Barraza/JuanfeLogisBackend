package services

import (
	"errors"
	"fmt"
	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"
	"juanfeLogis/models"
	"juanfeLogis/repositories"
	"juanfeLogis/utils"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepo    *repositories.ProductRepository
	paginationRepo *repositories.PaginationRepository
	boxStockRepo   *repositories.BoxStockRepository
}

func NewProductService(
	productRepo *repositories.ProductRepository,
	paginationRepo *repositories.PaginationRepository,
	boxStockRepo *repositories.BoxStockRepository,
) *ProductService {
	return &ProductService{
		productRepo:    productRepo,
		paginationRepo: paginationRepo,
		boxStockRepo:   boxStockRepo,
	}
}

func (s *ProductService) CreateProduct(req request.ProductRequest) (*response.ProductResponse, error) {
	donorUUID, err := uuid.Parse(req.DonorID)
	if err != nil {
		return nil, fmt.Errorf("Invalid donor Id, should be a valid UUID")
	}

	product := &models.Product{
		ProductTypeID:     req.ProductTypeID,
		DonorID:           donorUUID,
		Size:              req.Size,
		DonationPrice:     req.DonationPrice,
		SalePrice:         req.SalePrice,
		PhysicalCondition: req.PhysicalCondition,
		Disposition:       req.Disposition,
		Description:       req.Description,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, fmt.Errorf("Failed to create product: %v", err)
	}
	createdProduct, _ := s.productRepo.GetByID(product.ID.String())

	return s.toProductResponse(createdProduct), nil
}

func (s *ProductService) UpdateProduct(id string, req request.ProductRequest) (*response.ProductResponse, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("producto no encontrado")
	}

	donorUUID, err := uuid.Parse(req.DonorID)
	if err != nil {
		return nil, errors.New("donor_id inválido")
	}

	if req.ProductTypeID != 0 {
		product.ProductTypeID = req.ProductTypeID
	}
	if req.DonorID != "" {
		product.DonorID = donorUUID
	}
	if req.Size != "" {
		product.Size = req.Size
	}
	if req.DonationPrice != 0 {
		product.DonationPrice = req.DonationPrice
	}
	if req.SalePrice != 0 {
		product.SalePrice = req.SalePrice
	}
	if req.PhysicalCondition != "" {
		product.PhysicalCondition = req.PhysicalCondition
	}
	if req.Disposition != "" {
		product.Disposition = req.Disposition
	}
	if req.Description != "" {
		product.Description = req.Description
	}

	if err := s.productRepo.Update(product); err != nil {
		return nil, errors.New("error al actualizar el producto")
	}

	updatedProduct, _ := s.productRepo.GetByID(product.ID.String())

	return s.toProductResponse(updatedProduct), nil
}

func (s *ProductService) GetAllProducts(pagination *utils.Pagination, itemType, donor, size, disposition string) (*utils.Pagination, error) {
	result, products, err := s.productRepo.FindAllQuery(itemType, donor, size, disposition)
	if err != nil {
		return nil, errors.New("error al obtener los productos")
	}

	paginationResult, err := s.paginationRepo.GetPaginatedResults(result, pagination, &products)
	if err != nil {
		return nil, errors.New("error al paginar los productos")
	}

	res := make([]response.ProductResponse, len(products))
	for i, product := range products {
		res[i] = *s.toProductResponse(&product)
	}
	paginationResult.Data = res
	return paginationResult, nil
}

func (s *ProductService) GetByID(id string) (*response.ProductResponse, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("producto no encontrado")
	}
	return s.toProductResponse(product), nil
}

func (s *ProductService) DeleteProduct(id string) error {
	productID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("uuid inválido")
	}

	_, err = s.productRepo.GetByID(id)
	if err != nil {
		return errors.New("el producto no existe")
	}

	// Regla de Negocio: No borrar si hay stock
	totalStock, err := s.boxStockRepo.GetTotalQuantityByProductID(productID)
	if err != nil {
		return errors.New("error al verificar el stock del producto")
	}

	if totalStock > 0 {
		return fmt.Errorf("no se puede eliminar: el producto tiene aún %d unidades en cajas", totalStock)
	}

	if err := s.productRepo.Delete(id); err != nil {
		return errors.New("error al eliminar el producto")
	}
	return nil
}

func (s *ProductService) toProductResponse(product *models.Product) *response.ProductResponse {
	return &response.ProductResponse{
		ID:                product.ID.String(),
		ProductTypeID:     product.ProductTypeID,
		DonorID:           product.DonorID.String(),
		ProductTypeName:   product.ProductType.Name,
		DonorName:         product.Donor.Name,
		Size:              product.Size,
		DonationPrice:     product.DonationPrice,
		SalePrice:         product.SalePrice,
		PhysicalCondition: product.PhysicalCondition,
		Disposition:       product.Disposition,
		Description:       product.Description,
		CreatedAt:         product.CreatedAt.Format("2006-01-02"),
		UpdatedAt:         product.UpdatedAt.Format("2006-01-02"),
	}
}
