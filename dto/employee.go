package dto

import "github.com/baboyiban/go-api-server/models"

type CreateEmployeeRequest struct {
	Password string `json:"password" binding:"required"`
	Position string `json:"position" binding:"required,oneof=관리직 운송직"`
	IsActive *bool  `json:"is_active"` // optional, default true
}

type UpdateEmployeeRequest struct {
	Password string `json:"password"`
	Position string `json:"position" binding:"oneof=관리직 운송직"`
	IsActive *bool  `json:"is_active"`
}

type EmployeeResponse struct {
	EmployeeID int    `json:"employee_id"`
	Position   string `json:"position"`
	IsActive   bool   `json:"is_active"`
}

func ToEmployeeResponse(emp *models.Employee) EmployeeResponse {
	return EmployeeResponse{
		EmployeeID: emp.EmployeeID,
		Position:   emp.Position,
		IsActive:   emp.IsActive,
	}
}
