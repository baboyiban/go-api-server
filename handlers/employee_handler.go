package handlers

import (
	"net/http"
	"strconv"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

// CreateEmployee godoc
// @Summary      직원 생성
// @Description  새로운 직원을 생성합니다.
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param        employee  body      dto.CreateEmployeeRequest  true  "직원 정보"
// @Success      201       {object}  dto.EmployeeResponse
// @Failure      400       {object}  dto.ErrorResponse
// @Failure      500       {object}  dto.ErrorResponse
// @Router       /api/employee [post]
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	emp, err := h.service.CreateEmployee(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create employee", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, emp)
}

// GetEmployeeByID godoc
// @Summary      직원 단건 조회
// @Description  직원 ID로 직원 정보를 조회합니다.
// @Tags         employee
// @Produce      json
// @Param        id   path      int  true  "직원 ID"
// @Success      200  {object}  dto.EmployeeResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/employee/{id} [get]
func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid employee id"})
		return
	}
	emp, err := h.service.GetEmployeeByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Employee not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get employee", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, emp)
}

// DeleteEmployee godoc
// @Summary      직원 삭제
// @Description  직원 ID로 직원을 삭제합니다.
// @Tags         employee
// @Produce      json
// @Param        id   path      int  true  "직원 ID"
// @Success      204  "No Content"
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/employee/{id} [delete]
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid employee id"})
		return
	}
	err = h.service.DeleteEmployee(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Employee not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete employee", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateEmployee godoc
// @Summary      직원 정보 수정
// @Description  직원 ID로 직원 정보를 수정합니다.
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param        id       path      int                      true  "직원 ID"
// @Param        employee body      dto.UpdateEmployeeRequest true  "수정할 직원 정보"
// @Success      200      {object}  dto.EmployeeResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /api/employee/{id} [put]
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid employee id"})
		return
	}
	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	emp, err := h.service.UpdateEmployee(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Employee not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update employee", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, emp)
}

// ListEmployees godoc
// @Summary      모든 직원 조회
// @Description  모든 직원 정보를 반환합니다.
// @Tags         employee
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드 (예: -employee_id, -position 등)"
// @Success      200   {array}   dto.EmployeeResponse
// @Router       /api/employee [get]
func (h *EmployeeHandler) ListEmployees(c *gin.Context) {
	sortParam := c.Query("sort")
	emps, err := h.service.ListEmployees(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list employees", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, emps)
}

// SearchEmployees godoc
// @Summary      직원 검색
// @Description  쿼리 파라미터로 직원을 검색합니다.
// @Tags         employee
// @Produce      json
// @Param        employee_id  query     int     false  "직원 ID"
// @Param        position     query     string  false  "직책"
// @Param        is_active    query     bool    false  "활성 여부"
// @Param        sort         query     string  false  "정렬 필드 (예: -employee_id, -position 등)"
// @Success      200  {array}   dto.EmployeeResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Router       /api/employee/search [get]
func (h *EmployeeHandler) SearchEmployees(c *gin.Context) {
	params := map[string]string{}
	for _, key := range []string{"employee_id", "position", "is_active"} {
		if v := c.Query(key); v != "" {
			params[key] = v
		}
	}
	sortParam := c.Query("sort")
	emps, err := h.service.SearchEmployees(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search employees", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, emps)
}
