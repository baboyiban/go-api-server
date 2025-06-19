package service

import (
	"context"

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

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	var emp models.Employee
	if err := s.db.WithContext(ctx).Where("employee_id = ?", req.EmployeeID).First(&emp).Error; err != nil {
		return "", err
	}
	// 비밀번호 해시 검증 (예: bcrypt)
	if !utils.CheckPasswordHash(req.Password, emp.Password) {
		return "", gorm.ErrRecordNotFound
	}
	// JWT 토큰 생성
	token, err := utils.GenerateJWT(emp.EmployeeID, emp.Position)
	if err != nil {
		return "", err
	}
	return token, nil
}
