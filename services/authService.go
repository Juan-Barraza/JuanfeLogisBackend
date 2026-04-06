package services

import (
	"errors"
	"fmt"

	"juanfeLogis/dtos/request"
	"juanfeLogis/dtos/response"
	"juanfeLogis/repositories"
	"juanfeLogis/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(req request.LoginRequest) (*response.AuthResponse, error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("error al generar token")
	}

	response := &response.AuthResponse{
		Token: token,
		User: response.User{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}

	return response, nil
}
