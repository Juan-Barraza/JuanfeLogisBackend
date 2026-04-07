package utils

import (
	"errors"
	"strings"

	"juanfeLogis/dtos/request"
)

func ValidateProductRequest(req request.ProductRequest) error {
	var sb strings.Builder

	if req.ProductTypeID == 0 {
		sb.WriteString("ProductTypeID es requerido\n")
	}
	if req.DonorID == "" {
		sb.WriteString("DonorID es requerido\n")
	}
	if req.Size == "" {
		sb.WriteString("Size es requerido\n")
	}
	if req.DonationPrice == 0 {
		sb.WriteString("DonationPrice es requerido\n")
	}
	if req.SalePrice == 0 {
		sb.WriteString("SalePrice es requerido\n")
	}
	if req.PhysicalCondition == "" {
		sb.WriteString("PhysicalCondition es requerido\n")
	}
	if req.Disposition == "" {
		sb.WriteString("Disposition es requerido\n")
	}

	if sb.Len() > 0 {
		return errors.New(sb.String())
	}
	return nil
}

func ValidateBoxStockRequest(req request.BoxStockRequest) error {
	var sb strings.Builder

	if req.ProductID == "" {
		sb.WriteString("ProductID es requerido\n")
	}
	if req.Quantity == 0 {
		sb.WriteString("Quantity es requerido\n")
	}

	if sb.Len() > 0 {
		return errors.New(sb.String())
	}
	return nil
}
