package handlers

import (
	"log"
	"net/http"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/service"
	"github.com/baboyiban/go-api-server/utils"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related endpoints
type AuthHandler struct {
	service *service.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

// @Summary      로그인
// @Description  직원 ID와 비밀번호로 로그인합니다. 성공 시 JWT 토큰을 HttpOnly Secure 쿠키로 반환합니다.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      dto.LoginRequest  true  "로그인 정보"
// @Success      200    {object}  dto.LoginResponse "JWT 토큰이 HttpOnly Secure 쿠키(token)로도 반환됨"
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

	// SameSite 속성 설정 - 이 부분을 추가해야 합니다
	// c.SetSameSite(http.SameSiteNoneMode) // 크로스 도메인 요청을 위해 None으로 설정

	// domain := os.Getenv("BACKEND_DOMAIN")
	// // JWT를 HttpOnly Secure 쿠키로 설정
	// c.SetCookie(
	// 	"token",
	// 	token,
	// 	60*60*8, // 8시간
	// 	"/",
	// 	domain, // 도메인 (필요시 지정)
	// 	true,   // Secure (운영환경에서는 true, 개발환경에서는 false 가능)
	// 	true,   // HttpOnly
	// )

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token, // 필요 없다면 바디에서 제거해도 됨
		Employee: dto.EmployeeResponse{
			EmployeeID: emp.EmployeeID,
			Position:   emp.Position,
			IsActive:   emp.IsActive,
		},
	})
}

// @Summary      내 정보 조회
// @Description  JWT 토큰을 Authorization 헤더 또는 HttpOnly 쿠키(token)로 전달하여 로그인한 직원의 정보를 반환합니다.
// @Tags         auth
// @Produce      json
// @Success      200  {object}  dto.EmployeeResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/auth/me [get]
func (h *AuthHandler) Me(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	log.Println("Authorization:", authHeader)

	var tokenStr string

	if len(authHeader) >= 8 && authHeader[:7] == "Bearer " {
		tokenStr = authHeader[7:]
	} else {
		// 쿠키에서 토큰 시도
		cookie, err := c.Cookie("token")
		log.Println("token:", cookie)

		if err != nil || cookie == "" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Missing or invalid token"})
			return
		}
		tokenStr = cookie
	}

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
