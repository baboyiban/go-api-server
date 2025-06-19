package dto

type LoginRequest struct {
	EmployeeID int    `json:"employee_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
