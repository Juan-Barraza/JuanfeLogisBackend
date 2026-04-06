package services

import (
	"errors"
	"fmt"
	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"
	"juanfeLogis/models"
	"juanfeLogis/repositories"
)

type LocationService struct {
	locationRepo *repositories.LocationRepository
}

func NewLocationService(locationRepo *repositories.LocationRepository) *LocationService {
	return &LocationService{locationRepo: locationRepo}
}

func (s *LocationService) CreateLocation(req request.LocationRequest) (*response.LocationResponse, error) {
	lct, _ := s.locationRepo.GetByName(req.Name)
	if lct != nil {
		return nil, fmt.Errorf("error already exist location with the same name")
	}

	location := &models.Location{
		Name: req.Name,
	}

	if err := s.locationRepo.Create(location); err != nil {
		return nil, errors.New("error al crear la ubicación")
	}

	res := &response.LocationResponse{
		ID:        location.ID,
		Name:      location.Name,
		CreatedAt: location.CreatedAt,
	}

	return res, nil
}

func (s *LocationService) GetAllLocations() ([]response.LocationResponse, error) {
	locations, err := s.locationRepo.FindAll()
	if err != nil {
		return nil, errors.New("error al obtener las ubicaciones")
	}

	var res []response.LocationResponse
	for _, loc := range locations {
		res = append(res, response.LocationResponse{
			ID:        loc.ID,
			Name:      loc.Name,
			CreatedAt: loc.CreatedAt,
		})
	}

	return res, nil
}

func (s *LocationService) UpdateLocation(id uint, req request.LocationRequest) (*response.LocationResponse, error) {
	location, err := s.locationRepo.GetByID(id)
	if err != nil || location == nil {
		return nil, errors.New("ubicación no encontrada")
	}

	lct, _ := s.locationRepo.GetByName(req.Name)
	if lct != nil {
		return nil, fmt.Errorf("error already exist location with the same name")
	}

	location.Name = req.Name

	if err := s.locationRepo.Update(location); err != nil {
		return nil, errors.New("error al actualizar la ubicación")
	}

	res := &response.LocationResponse{
		ID:        location.ID,
		Name:      location.Name,
		CreatedAt: location.CreatedAt,
	}

	return res, nil
}
