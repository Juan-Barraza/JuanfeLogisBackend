package services

import (
	"errors"
	"fmt"
	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"
	"juanfeLogis/models"
	"juanfeLogis/repositories"
	"juanfeLogis/utils"
)

type BoxService struct {
	boxRepo        *repositories.BoxRepository
	paginationRepo *repositories.PaginationRepository
}

func NewBoxService(boxRepo *repositories.BoxRepository, paginationRepo *repositories.PaginationRepository) *BoxService {
	return &BoxService{boxRepo: boxRepo, paginationRepo: paginationRepo}
}

func (s *BoxService) CreateBox(req request.BoxRequest) (*response.BoxResponse, error) {
	box := &models.Box{
		Name:       req.Name,
		LocationID: req.LocationID,
	}

	// 1. Guardamos la caja para que la DB le asigne su UUID
	if err := s.boxRepo.Create(box); err != nil {
		return nil, errors.New("error al crear la caja")
	}
	if len(req.LabelIDs) > 0 {
		labels, err := s.boxRepo.SetLabels(box.ID.String(), req.LabelIDs)
		if err != nil {
			return nil, errors.New("error al asignar etiquetas")
		}
		box.Labels = labels
	}

	// 2. Generamos el QR pasando el nuevo ID
	box.QRCodeURL = utils.GenerateBoxQR(box.ID.String())
	if err := s.boxRepo.Update(box); err != nil {
		return nil, errors.New("error al guardar el QR")
	}
	createdBox, _ := s.boxRepo.GetByID(box.ID.String())

	return s.toBoxResponse(createdBox), nil
}

func (s *BoxService) GetBoxDetail(boxID string) (*response.BoxDetailResponse, error) {
	box, stockList, err := s.boxRepo.GetBoxWithStock(boxID)
	if err != nil {
		return nil, errors.New("caja no encontrada")
	}

	// Mapeamos los productos del stock al DTO
	products := []response.BoxStockItem{}
	for _, item := range stockList {
		products = append(products, response.BoxStockItem{
			ProductID:         item.ProductID.String(),
			ProductTypeName:   item.Product.ProductType.Name,
			DonorName:         item.Product.Donor.Name,
			Quantity:          item.Quantity,
			DonationPrice:     item.Product.DonationPrice,
			SalePrice:         item.Product.SalePrice,
			PhysicalCondition: item.Product.PhysicalCondition,
		})
	}

	return &response.BoxDetailResponse{
		Box:      *s.toBoxResponse(box),
		Products: products,
	}, nil
}

func (s *BoxService) UpdateBox(id string, req request.BoxRequest) (*response.BoxResponse, error) {
	box, err := s.boxRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("caja no encontrada")
	}

	// Actualizamos los campos
	if req.Name != "" {
		box.Name = req.Name
	}
	if req.LocationID != 0 {
		box.LocationID = req.LocationID
	}
	if req.LabelIDs != nil {
		labels, err := s.boxRepo.SetLabels(box.ID.String(), req.LabelIDs)
		if err != nil {
			return nil, errors.New("error al asignar etiquetas")
		}
		// Sincronizamos el objeto en memoria para que el posterior Update no ruidoso
		box.Labels = labels
	}

	if err := s.boxRepo.Update(box); err != nil {
		return nil, errors.New("error al actualizar la caja")
	}
	updatedBox, _ := s.boxRepo.GetByID(box.ID.String())

	return s.toBoxResponse(updatedBox), nil
}

func (s *BoxService) DeleteBox(id string) error {
	_, err := s.boxRepo.GetByID(id)
	if err != nil {
		return errors.New("la caja no existe o ya fue eliminada")
	}

	// 1. Verificamos si tiene productos
	_, stock, err := s.boxRepo.GetBoxWithStock(id)
	if err != nil {
		return errors.New("error al verificar el stock")
	}

	if len(stock) > 0 {
		return errors.New("no se puede eliminar la caja porque contiene productos")
	}

	// 2. Si está vacía, la eliminamos
	if err := s.boxRepo.Delete(id); err != nil {
		return errors.New("error al eliminar la caja")
	}

	return nil
}

func (s *BoxService) GetAllBoxes(pagination *utils.Pagination, name string, location string) (*utils.Pagination, error) {
	// 1. Ejecutamos la query que ahora maneja los filtros en el repo
	result, boxes, err := s.boxRepo.FindAllQuery(name, location)
	if err != nil {
		return nil, errors.New("error al obtener las cajas")
	}

	paginationResult, err := s.paginationRepo.GetPaginatedResults(result, pagination, &boxes)
	if err != nil {
		return nil, errors.New("error al paginar las cajas")
	}

	// 4. Mapeamos las cajas al DTO
	res := make([]response.BoxResponse, len(boxes))
	for i, box := range boxes {
		res[i] = *s.toBoxResponse(&box)
	}
	paginationResult.Data = res
	return paginationResult, nil
}

func (s *BoxService) GetByID(id string) (*response.BoxResponse, error) {
	box, err := s.boxRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("caja no encontrada")
	}
	return s.toBoxResponse(box), nil
}

func (s *BoxService) toBoxResponse(box *models.Box) *response.BoxResponse {

	labels := make([]string, len(box.Labels))
	for i, label := range box.Labels {
		labels[i] = label.Name
	}

	return &response.BoxResponse{
		ID:           box.ID.String(),
		Name:         box.Name,
		LocationID:   box.LocationID,
		LocationName: box.Location.Name,
		Labels:       labels,
		QRCodeURL:    box.QRCodeURL,
		CreatedAt:    box.CreatedAt,
	}
}
