package services

import (
	"errors"
	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"
	"juanfeLogis/models"
	"juanfeLogis/repositories"
)

type ProductTypeService struct {
	ptRepo *repositories.ProductTypeRepository
}

func NewProductTypeService(ptRepo *repositories.ProductTypeRepository) *ProductTypeService {
	return &ProductTypeService{ptRepo: ptRepo}
}

func (s *ProductTypeService) CreateProductType(req request.ProductTypeRequest) (*response.ProductTypeResponse, error) {

	existing, _ := s.ptRepo.GetByName(req.Name)
	if existing != nil {
		return nil, errors.New("ya existe un tipo de producto con ese nombre")
	}

	pt := &models.ProductType{
		Name: req.Name,
	}

	if err := s.ptRepo.Create(pt); err != nil {
		return nil, errors.New("error al crear el tipo de producto")
	}

	return &response.ProductTypeResponse{ID: pt.ID, Name: pt.Name}, nil
}

func (s *ProductTypeService) GetAllProductTypes() ([]response.ProductTypeResponse, error) {
	pts, err := s.ptRepo.FindAll()
	if err != nil {
		return nil, errors.New("error al obtener los tipos de producto")
	}

	var res []response.ProductTypeResponse
	for _, pt := range pts {
		res = append(res, response.ProductTypeResponse{
			ID:   pt.ID,
			Name: pt.Name,
		})
	}

	return res, nil
}

func (s *ProductTypeService) UpdateProductType(id uint, req request.ProductTypeRequest) (*response.ProductTypeResponse, error) {
	pt, err := s.ptRepo.GetByID(id)
	if err != nil || pt == nil {
		return nil, errors.New("tipo de producto no encontrado")
	}

	// Verificar si el nuevo nombre ya está ocupado por OTRO tipo de producto
	existing, _ := s.ptRepo.GetByName(req.Name)
	if existing != nil && existing.ID != id {
		return nil, errors.New("ya existe otro tipo de producto con ese nombre")
	}

	pt.Name = req.Name

	if err := s.ptRepo.Update(pt); err != nil {
		return nil, errors.New("error al actualizar el tipo de producto")
	}

	return &response.ProductTypeResponse{ID: pt.ID, Name: pt.Name}, nil
}
