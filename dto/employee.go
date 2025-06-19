package dto

type CreateEmployeeRequest struct {
	Password string `json:"password" binding:"required"`
	Position string `json:"position" binding:"required,oneof=관리직 운송직"`
	IsActive *bool  `json:"is_active"` // optional, default true
}

type UpdateEmployeeRequest struct {
	Password string `json:"password"`
	Position string `json:"position" binding:"omitempty,oneof=관리직 운송직"`
	IsActive *bool  `json:"is_active"`
}

type EmployeeResponse struct {
	EmployeeID int    `json:"employee_id"`
	Position   string `json:"position"`
	IsActive   bool   `json:"is_active"`
}
