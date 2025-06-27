package handlers

import (
	"net/http"
	"strconv"

	"github.com/baboyiban/go-api-server/constants"
	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VehicleHandler struct {
	service *service.VehicleService
}

func NewVehicleHandler(s *service.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: s}
}

// CreateVehicle godoc
// @Summary      차량 생성
// @Description  새로운 차량을 생성합니다.
// @Tags         vehicle
// @Accept       json
// @Produce      json
// @Param        vehicle  body      dto.CreateVehicleRequest  true  "차량 정보"
// @Success      201      {object}  dto.VehicleResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /api/vehicle [post]
func (h *VehicleHandler) CreateVehicle(c *gin.Context) {
	var req dto.CreateVehicleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	vehicle, err := h.service.CreateVehicle(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create vehicle", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, vehicle)
}

// GetVehicleByID godoc
// @Summary      차량 단건 조회
// @Description  vehicle_id로 차량을 조회합니다.
// @Tags         vehicle
// @Produce      json
// @Param        id   path      int  true  "차량 ID"
// @Success      200  {object}  dto.VehicleResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/vehicle/{id} [get]
func (h *VehicleHandler) GetVehicleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid vehicle id"})
		return
	}
	vehicle, err := h.service.GetVehicleByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Vehicle not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get vehicle", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

// DeleteVehicle godoc
// @Summary      차량 삭제
// @Description  vehicle_id로 차량을 삭제합니다.
// @Tags         vehicle
// @Produce      json
// @Param        id   path      int  true  "차량 ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/vehicle/{id} [delete]
func (h *VehicleHandler) DeleteVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid vehicle id"})
		return
	}
	err = h.service.DeleteVehicle(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Vehicle not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete vehicle", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateVehicle godoc
// @Summary      차량 정보 수정
// @Description  vehicle_id로 차량 정보를 수정합니다.
// @Tags         vehicle
// @Accept       json
// @Produce      json
// @Param        id       path      int                      true  "차량 ID"
// @Param        vehicle  body      dto.UpdateVehicleRequest  true  "수정할 차량 정보"
// @Success      200      {object}  dto.VehicleResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /api/vehicle/{id} [put]
func (h *VehicleHandler) UpdateVehicle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid vehicle id"})
		return
	}
	var req dto.UpdateVehicleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	vehicle, err := h.service.UpdateVehicle(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Vehicle not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update vehicle", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

// ListVehicles godoc
// @Summary      모든 차량 조회
// @Description  모든 차량 정보를 반환합니다.
// @Tags         vehicle
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드"
// @Success      200   {array}   dto.VehicleResponse
// @Router       /api/vehicle [get]
func (h *VehicleHandler) ListVehicles(c *gin.Context) {
	sortParam := c.Query("sort")
	vehicles, err := h.service.ListVehicles(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list vehicles", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

// SearchVehicles godoc
// @Summary      차량 검색
// @Description  쿼리 파라미터로 차량을 검색합니다.
// @Tags         vehicle
// @Produce      json
// @Param        vehicle_id   query     int     false  "차량 ID"
// @Param        model        query     string  false  "차량 모델"
// @Param        status       query     string  false  "상태"
// @Param        sort         query     string  false  "정렬 필드"
// @Success      200  {array}   dto.VehicleResponse
// @Router       /api/vehicle/search [get]
func (h *VehicleHandler) SearchVehicles(c *gin.Context) {
	params := dto.ExtractAllowedParams(c.Request.URL.Query(), constants.VehicleAllowedFields)
	sortParam := c.Query("sort")
	vehicles, err := h.service.SearchVehicles(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search vehicles", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}
