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

type TripLogHandler struct {
	service *service.TripLogService
}

func NewTripLogHandler(s *service.TripLogService) *TripLogHandler {
	return &TripLogHandler{service: s}
}

// CreateTripLog godoc
// @Summary      차량 운행 로그 생성
// @Description  새로운 차량 운행 로그를 생성합니다.
// @Tags         trip_log
// @Accept       json
// @Produce      json
// @Param        trip_log  body      dto.CreateTripLogRequest  true  "차량 운행 로그 정보"
// @Success      201         {object}  dto.TripLogResponse
// @Failure      400         {object}  dto.ErrorResponse
// @Failure      500         {object}  dto.ErrorResponse
// @Router       /api/trip-log [post]
func (h *TripLogHandler) CreateTripLog(c *gin.Context) {
	var req dto.CreateTripLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	trip, err := h.service.CreateTripLog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create trip_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, trip)
}

// GetTripLogByID godoc
// @Summary      차량 운행 로그 단건 조회
// @Description  trip_id로 차량 운행 로그를 조회합니다.
// @Tags         trip_log
// @Produce      json
// @Param        id   path      int  true  "차량 운행 로그 trip_id"
// @Success      200  {object}  dto.TripLogResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/trip-log/{id} [get]
func (h *TripLogHandler) GetTripLogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid trip_log id"})
		return
	}
	trip, err := h.service.GetTripLogByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "TripLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get trip_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

// DeleteTripLog godoc
// @Summary      차량 운행 로그 삭제
// @Description  trip_id로 차량 운행 로그를 삭제합니다.
// @Tags         trip_log
// @Produce      json
// @Param        id   path      int  true  "차량 운행 로그 trip_id"
// @Success      204  "No Content"
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/trip-log/{id} [delete]
func (h *TripLogHandler) DeleteTripLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid trip_log id"})
		return
	}
	err = h.service.DeleteTripLog(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "TripLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete trip_log", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateTripLog godoc
// @Summary      차량 운행 로그 정보 수정
// @Description  trip_id로 차량 운행 로그 정보를 수정합니다.
// @Tags         trip_log
// @Accept       json
// @Produce      json
// @Param        id          path      int                        true  "차량 운행 로그 trip_id"
// @Param        trip_log    body      dto.UpdateTripLogRequest   true  "수정할 차량 운행 로그 정보"
// @Success      200         {object}  dto.TripLogResponse
// @Failure      400         {object}  dto.ErrorResponse
// @Failure      404         {object}  dto.ErrorResponse
// @Failure      500         {object}  dto.ErrorResponse
// @Router       /api/trip-log/{id} [put]
func (h *TripLogHandler) UpdateTripLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid trip_log id"})
		return
	}
	var req dto.UpdateTripLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	trip, err := h.service.UpdateTripLog(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "TripLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update trip_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

// ListTripLogs godoc
// @Summary      모든 차량 운행 로그 조회
// @Description  모든 차량 운행 로그 정보를 반환합니다.
// @Tags         trip_log
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드 (예: -trip_id, -start_time 등)"
// @Success      200   {array}   dto.TripLogResponse
// @Router       /api/trip-log [get]
func (h *TripLogHandler) ListTripLogs(c *gin.Context) {
	sortParam := c.Query("sort")
	trips, err := h.service.ListTripLogs(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list trip_logs", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}

// SearchTripLogs godoc
// @Summary      모든 차량 운행 로그 검색
// @Description  쿼리 파라미터로 모든 차량 운행 로그를 검색합니다.
// @Tags         trip_log
// @Produce      json
// @Param        trip_id        query     int     false  "trip_id"
// @Param        vehicle_id     query     string  false  "차량 ID"
// @Param        start_time     query     string  false  "출발 시각 (YYYY-MM-DD)"
// @Param        end_time       query     string  false  "도착 시각 (YYYY-MM-DD)"
// @Param        status         query     string  false  "상태"
// @Param        destination    query     string  false  "목적지"
// @Param        sort           query     string  false  "정렬 필드 (예: -trip_id, -start_time 등)"
// @Success      200  {array}   dto.TripLogResponse
// @Router       /api/trip-log/search [get]
func (h *TripLogHandler) SearchTripLogs(c *gin.Context) {
	params := dto.ExtractAllowedParams(c.Request.URL.Query(), constants.TripLogAllowedFields)
	sortParam := c.Query("sort")
	trips, err := h.service.SearchTripLogs(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search trip_logs", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}
