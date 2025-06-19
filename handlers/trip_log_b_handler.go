package handlers

import (
	"net/http"
	"strconv"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TripLogBHandler struct {
	service *service.TripLogBService
}

func NewTripLogBHandler(s *service.TripLogBService) *TripLogBHandler {
	return &TripLogBHandler{service: s}
}

// CreateTripLogB godoc
// @Summary      B차 운행 로그 생성
// @Description  새로운 B차 운행 로그를 생성합니다.
// @Tags         trip_log_b
// @Accept       json
// @Produce      json
// @Param        trip_log_b  body      dto.CreateTripLogBRequest  true  "B차 운행 로그 정보"
// @Success      201         {object}  dto.TripLogBResponse
// @Failure      400         {object}  dto.ErrorResponse
// @Failure      500         {object}  dto.ErrorResponse
// @Router       /api/trip-log-b [post]
func (h *TripLogBHandler) CreateTripLogB(c *gin.Context) {
	var req dto.CreateTripLogBRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	trip, err := h.service.CreateTripLogB(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create trip_log_b", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, trip)
}

// GetTripLogBByID godoc
// @Summary      B차 운행 로그 단건 조회
// @Description  trip_id로 B차 운행 로그를 조회합니다.
// @Tags         trip_log_b
// @Produce      json
// @Param        id   path      int  true  "B차 운행 로그 trip_id"
// @Success      200  {object}  dto.TripLogBResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/trip-log-b/{id} [get]
func (h *TripLogBHandler) GetTripLogBByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid trip_log_b id"})
		return
	}
	trip, err := h.service.GetTripLogBByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "TripLogB not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get trip_log_b", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

// DeleteTripLogB godoc
// @Summary      B차 운행 로그 삭제
// @Description  trip_id로 B차 운행 로그를 삭제합니다.
// @Tags         trip_log_b
// @Produce      json
// @Param        id   path      int  true  "B차 운행 로그 trip_id"
// @Success      204  "No Content"
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/trip-log-b/{id} [delete]
func (h *TripLogBHandler) DeleteTripLogB(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid trip_log_b id"})
		return
	}
	err = h.service.DeleteTripLogB(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "TripLogB not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete trip_log_b", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateTripLogB godoc
// @Summary      B차 운행 로그 정보 수정
// @Description  trip_id로 B차 운행 로그 정보를 수정합니다.
// @Tags         trip_log_b
// @Accept       json
// @Produce      json
// @Param        id          path      int                        true  "B차 운행 로그 trip_id"
// @Param        trip_log_b  body      dto.UpdateTripLogBRequest  true  "수정할 B차 운행 로그 정보"
// @Success      200         {object}  dto.TripLogBResponse
// @Failure      400         {object}  dto.ErrorResponse
// @Failure      404         {object}  dto.ErrorResponse
// @Failure      500         {object}  dto.ErrorResponse
// @Router       /api/trip-log-b/{id} [put]
func (h *TripLogBHandler) UpdateTripLogB(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid trip_log_b id"})
		return
	}
	var req dto.UpdateTripLogBRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	trip, err := h.service.UpdateTripLogB(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "TripLogB not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update trip_log_b", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

// ListTripLogBs godoc
// @Summary      모든 B차 운행 로그 조회
// @Description  모든 B차 운행 로그 정보를 반환합니다.
// @Tags         trip_log_b
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드 (예: -trip_id, -start_time 등)"
// @Success      200   {array}   dto.TripLogBResponse
// @Router       /api/trip-log-b [get]
func (h *TripLogBHandler) ListTripLogBs(c *gin.Context) {
	sortParam := c.Query("sort")
	trips, err := h.service.ListTripLogBs(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list trip_log_bs", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}

// SearchTripLogBs godoc
// @Summary      B차 운행 로그 검색
// @Description  쿼리 파라미터로 B차 운행 로그를 검색합니다.
// @Tags         trip_log_b
// @Produce      json
// @Param        trip_id        query     int     false  "trip_id"
// @Param        vehicle_id     query     string  false  "차량 ID"
// @Param        status         query     string  false  "상태"
// @Param        destination_1  query     string  false  "목적지1"
// @Param        destination_2  query     string  false  "목적지2"
// @Param        destination_3  query     string  false  "목적지3"
// @Param        sort           query     string  false  "정렬 필드 (예: -trip_id, -start_time 등)"
// @Success      200  {array}   dto.TripLogBResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Router       /api/trip-log-b/search [get]
func (h *TripLogBHandler) SearchTripLogBs(c *gin.Context) {
	params := map[string]string{}
	for _, key := range []string{"trip_id", "vehicle_id", "status", "destination_1", "destination_2", "destination_3"} {
		if v := c.Query(key); v != "" {
			params[key] = v
		}
	}
	sortParam := c.Query("sort")
	trips, err := h.service.SearchTripLogBs(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search trip_log_bs", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}
