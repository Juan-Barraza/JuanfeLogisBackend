package services

import (
	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"
	"juanfeLogis/repositories"
)

type DashboardService struct {
	repo *repositories.DashboardRepository
}

func NewDashboardService(repo *repositories.DashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetLogisticsKPIs(req request.DashboardFilterRequest) (response.LogisticsKPIResponse, error) {
	return s.repo.GetLogisticsKPIs(req)
}

func (s *DashboardService) GetLogisticsDistribution(req request.DashboardFilterRequest) (response.LogisticsDistributionResponse, error) {
	return s.repo.GetLogisticsDistribution(req)
}

func (s *DashboardService) GetLocationDistribution(req request.DashboardFilterRequest) ([]response.DistributionItem, error) {
	return s.repo.GetLocationDistribution(req)
}

func (s *DashboardService) GetTopDonorsLogistics(req request.DashboardFilterRequest) ([]response.DistributionItem, error) {
	return s.repo.GetTopDonorsLogistics(req)
}

func (s *DashboardService) GetFinancialKPIs(req request.DashboardFilterRequest) (response.FinancialKPIResponse, error) {
	return s.repo.GetFinancialKPIs(req)
}

func (s *DashboardService) GetFinancialTrends(req request.DashboardFilterRequest) (response.FinancialTrendResponse, error) {
	return s.repo.GetFinancialTrends(req)
}

func (s *DashboardService) GetTopDonorsFinancial(req request.DashboardFilterRequest) ([]response.FinancialDistributionItem, error) {
	return s.repo.GetTopDonorsFinancial(req)
}

func (s *DashboardService) GetProfitability(req request.DashboardFilterRequest) (response.ProfitabilityResponse, error) {
	return s.repo.GetProfitability(req)
}
