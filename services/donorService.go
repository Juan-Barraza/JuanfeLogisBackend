package services

import (
	"errors"
	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"
	"juanfeLogis/models"
	"juanfeLogis/repositories"
	"juanfeLogis/utils"
)

type DonorService struct {
	donorRepo      *repositories.DonorRepository
	paginationRepo *repositories.PaginationRepository
}

func NewDonorService(donorRepo *repositories.DonorRepository, paginationRepo *repositories.PaginationRepository) *DonorService {
	return &DonorService{donorRepo: donorRepo, paginationRepo: paginationRepo}
}

func (s *DonorService) CreateDonor(req request.DonorRequest) (*response.DonorResponse, error) {
	if req.Name == "" || req.Type == "" {
		return nil, errors.New("nombre y tipo de donante son requeridos")
	}

	existing, _ := s.donorRepo.GetByName(req.Name)
	if existing != nil {
		return nil, errors.New("ya existe un donante con este nombre")
	}

	donor := &models.Donor{
		Name: req.Name,
		Type: req.Type,
	}

	if err := s.donorRepo.Create(donor); err != nil {
		return nil, errors.New("error al crear el donante")
	}

	return &response.DonorResponse{
		ID: donor.ID.String(), Name: donor.Name, Type: donor.Type, CreatedAt: donor.CreatedAt,
	}, nil
}

func (s *DonorService) GetAllDonors(pagination *utils.Pagination) (*utils.Pagination, error) {
	query, donors, err := s.donorRepo.FindAll()
	if err != nil {
		return nil, errors.New("error al obtener donantes")
	}

	paginationResult, err := s.paginationRepo.GetPaginatedResults(query, pagination, &donors)
	if err != nil {
		return nil, err
	}

	res := make([]response.DonorResponse, len(donors))
	for i, d := range donors {
		res[i] = response.DonorResponse{
			ID: d.ID.String(), Name: d.Name, Type: d.Type, CreatedAt: d.CreatedAt,
		}
	}
	paginationResult.Data = res
	return paginationResult, nil
}

func (s *DonorService) UpdateDonor(id string, req request.DonorRequest) (*response.DonorResponse, error) {
	donor, err := s.donorRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("donante no encontrado")
	}

	existing, _ := s.donorRepo.GetByName(req.Name)
	if existing != nil && existing.ID.String() != id {
		return nil, errors.New("ya existe otro donante con ese nombre")
	}

	donor.Name = req.Name
	donor.Type = req.Type

	if err := s.donorRepo.Update(donor); err != nil {
		return nil, errors.New("error al actualizar donante")
	}

	return &response.DonorResponse{
		ID: donor.ID.String(), Name: donor.Name, Type: donor.Type, CreatedAt: donor.CreatedAt,
	}, nil
}

func (s *DonorService) GetByName(name string) (*response.DonorResponse, error) {
	donor, err := s.donorRepo.GetByName(name)
	if err != nil {
		return nil, errors.New("donante no encontrado")
	}
	return &response.DonorResponse{
		ID: donor.ID.String(), Name: donor.Name, Type: donor.Type, CreatedAt: donor.CreatedAt,
	}, nil
}
