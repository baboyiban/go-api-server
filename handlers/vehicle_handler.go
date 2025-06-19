package handlers

import (
	"net/http"
	"strconv"

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
// @Description  차량 ID로 차량 정보를 조회합니다.
// @Tags         vehicle
// @Produce      json
// @Param        id   path      int  true  "차량 Internal ID"
// @Success      200  {object}  dto.VehicleResponse
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
// @Description  차량 ID로 차량을 삭제합니다.
// @Tags         vehicle
// @Produce      json
// @Param        id   path      int  true  "차량 Internal ID"
// @Success      204  "No Content"
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
// @Description  차량 ID로 차량 정보를 수정합니다.
// @Tags         vehicle
// @Accept       json
// @Produce      json
// @Param        id      path      int                      true  "차량 Internal ID"
// @Param        vehicle body      dto.UpdateVehicleRequest  true  "수정할 차량 정보"
// @Success      200     {object}  dto.VehicleResponse
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      404     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
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
// @Param        sort  query     string  false  "정렬 필드 (예: -internal_id, -vehicle_id 등)"
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
// @Param        internal_id        query     int     false  "차량 Internal ID"
// @Param        vehicle_id         query     string  false  "차량 ID"
// @Param        current_load       query     int     false  "현재 적재량"
// @Param        max_load           query     int     false  "최대 적재량"
// @Param        led_status         query     string  false  "LED 상태"
// @Param        needs_confirmation query     bool    false  "확인 필요 여부"
// @Param        coord_x            query     int     false  "X 좌표"
// @Param        coord_y            query     int     false  "Y 좌표"
// @Param        sort               query     string  false  "정렬 필드 (예: -internal_id, -vehicle_id 등)"
// @Success      200  {array}   dto.VehicleResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Router       /api/vehicle/search [get]
func (h *VehicleHandler) SearchVehicles(c *gin.Context) {
	params := map[string]string{}
	for _, key := range []string{"internal_id", "vehicle_id", "current_load", "max_load", "led_status", "needs_confirmation", "coord_x", "coord_y"} {
		if v := c.Query(key); v != "" {
			params[key] = v
		}
	}
	sortParam := c.Query("sort")
	vehicles, err := h.service.SearchVehicles(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search vehicles", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}
