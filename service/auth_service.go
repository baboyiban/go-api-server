package service

import (
	"context"
	"errors"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (string, *models.Employee, error) {
	var emp models.Employee
	if err := s.db.WithContext(ctx).Where("employee_id = ?", req.EmployeeID).First(&emp).Error; err != nil {
		return "", nil, err
	}
	if !utils.CheckPasswordHash(req.Password, emp.Password) {
		return "", nil, errors.New("invalid password")
	}
	token, err := utils.GenerateJWT(emp.EmployeeID, emp.Position)
	if err != nil {
		return "", nil, err
	}
	return token, &emp, nil
}
