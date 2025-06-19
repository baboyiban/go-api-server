package handlers

import (
	"net/http"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

// Login godoc
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
	token, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}
