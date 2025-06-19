package handlers

import (
	"net/http"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/service"
	"github.com/baboyiban/go-api-server/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

// @Summary      로그인
// @Description  직원 ID와 비밀번호로 로그인합니다. 성공 시 JWT 토큰을 반환합니다.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      dto.LoginRequest  true  "로그인 정보"
// @Success      200    {object}  dto.LoginResponse
// @Failure      400    {object}  dto.ErrorResponse
// @Failure      401    {object}  dto.ErrorResponse
// @Router       /api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request"})
		return
	}
	token, emp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
		Employee: dto.EmployeeResponse{
			EmployeeID: emp.EmployeeID,
			Position:   emp.Position,
			IsActive:   emp.IsActive,
		},
	})
}

// @Summary      내 정보 조회
// @Description  JWT 토큰을 이용해 로그인한 직원의 정보를 반환합니다.
// @Tags         auth
// @Produce      json
// @Success      200  {object}  dto.EmployeeResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Router       /api/auth/me [get]
func (h *AuthHandler) Me(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Missing or invalid token"})
		return
	}
	tokenStr := authHeader[7:]
	claims, err := utils.ParseJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid token"})
		return
	}
	employeeID, ok := claims["employee_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid token claims"})
		return
	}
	var emp models.Employee
	if err := h.service.DB().Where("employee_id = ?", int(employeeID)).First(&emp).Error; err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "User not found"})
		return
	}
	c.JSON(http.StatusOK, dto.EmployeeResponse{
		EmployeeID: emp.EmployeeID,
		Position:   emp.Position,
		IsActive:   emp.IsActive,
	})
}
