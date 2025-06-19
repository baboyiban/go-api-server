package service

import (
	"context"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"gorm.io/gorm"
)

var allowedEmployeeSortFields = map[string]bool{
	"employee_id": true,
	"position":    true,
	"is_active":   true,
}

func applyEmployeeSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedEmployeeSortFields[field] {
		order := field
		if desc {
			order += " DESC"
		} else {
			order += " ASC"
		}
		return query.Order(order)
	}
	return query
}

type EmployeeService struct {
	db *gorm.DB
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	return &EmployeeService{db: db}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, req dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}
	emp := models.Employee{
		Password: req.Password,
		Position: req.Position,
		IsActive: isActive,
	}
	if err := s.db.WithContext(ctx).Create(&emp).Error; err != nil {
		return nil, err
	}
	return toEmployeeResponse(&emp), nil
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id int) (*dto.EmployeeResponse, error) {
	var emp models.Employee
	if err := s.db.WithContext(ctx).Where("employee_id = ?", id).First(&emp).Error; err != nil {
		return nil, err
	}
	return toEmployeeResponse(&emp), nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	result := s.db.WithContext(ctx).Where("employee_id = ?", id).Delete(&models.Employee{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, id int, req dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	var emp models.Employee
	if err := s.db.WithContext(ctx).Where("employee_id = ?", id).First(&emp).Error; err != nil {
		return nil, err
	}
	if req.Password != "" {
		emp.Password = req.Password
	}
	if req.Position != "" {
		emp.Position = req.Position
	}
	if req.IsActive != nil {
		emp.IsActive = *req.IsActive
	}
	if err := s.db.WithContext(ctx).Save(&emp).Error; err != nil {
		return nil, err
	}
	return toEmployeeResponse(&emp), nil
}

func (s *EmployeeService) ListEmployees(ctx context.Context, sort string) ([]dto.EmployeeResponse, error) {
	var emps []models.Employee
	query := s.db.WithContext(ctx).Model(&models.Employee{})
	query = applyEmployeeSort(query, sort)
	if err := query.Find(&emps).Error; err != nil {
		return nil, err
	}
	var res []dto.EmployeeResponse
	for _, e := range emps {
		res = append(res, *toEmployeeResponse(&e))
	}
	return res, nil
}

func (s *EmployeeService) SearchEmployees(ctx context.Context, params map[string]string, sort string) ([]dto.EmployeeResponse, error) {
	var emps []models.Employee
	query := s.db.WithContext(ctx).Model(&models.Employee{})
	for k, v := range params {
		query = query.Where(k+" = ?", v)
	}
	query = applyEmployeeSort(query, sort)
	if err := query.Find(&emps).Error; err != nil {
		return nil, err
	}
	var res []dto.EmployeeResponse
	for _, e := range emps {
		res = append(res, *toEmployeeResponse(&e))
	}
	return res, nil
}

func toEmployeeResponse(m *models.Employee) *dto.EmployeeResponse {
	return &dto.EmployeeResponse{
		EmployeeID: m.EmployeeID,
		Position:   m.Position,
		IsActive:   m.IsActive,
	}
}
