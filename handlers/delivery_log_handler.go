package handlers

import (
	"net/http"
	"strconv"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeliveryLogHandler struct {
	service *service.DeliveryLogService
}

func NewDeliveryLogHandler(s *service.DeliveryLogService) *DeliveryLogHandler {
	return &DeliveryLogHandler{service: s}
}

// CreateDeliveryLog godoc
// @Summary      배송 로그 생성
// @Description  새로운 배송 로그를 생성합니다.
// @Tags         delivery_log
// @Accept       json
// @Produce      json
// @Param        delivery_log  body      dto.CreateDeliveryLogRequest  true  "배송 로그 정보"
// @Success      201           {object}  dto.DeliveryLogResponse
// @Failure      400           {object}  dto.ErrorResponse
// @Failure      500           {object}  dto.ErrorResponse
// @Router       /api/delivery-log [post]
func (h *DeliveryLogHandler) CreateDeliveryLog(c *gin.Context) {
	var req dto.CreateDeliveryLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	log, err := h.service.CreateDeliveryLog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create delivery_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, log)
}

// GetDeliveryLogByID godoc
// @Summary      배송 로그 단건 조회
// @Description  trip_id로 배송 로그를 조회합니다.
// @Tags         delivery_log
// @Produce      json
// @Param        id   path      int  true  "배송 로그 trip_id"
// @Success      200  {object}  dto.DeliveryLogResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/delivery-log/{id} [get]
func (h *DeliveryLogHandler) GetDeliveryLogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid delivery_log id"})
		return
	}
	log, err := h.service.GetDeliveryLogByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "DeliveryLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get delivery_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

// DeleteDeliveryLog godoc
// @Summary      배송 로그 삭제
// @Description  trip_id로 배송 로그를 삭제합니다.
// @Tags         delivery_log
// @Produce      json
// @Param        id   path      int  true  "배송 로그 trip_id"
// @Success      204  "No Content"
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/delivery-log/{id} [delete]
func (h *DeliveryLogHandler) DeleteDeliveryLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid delivery_log id"})
		return
	}
	err = h.service.DeleteDeliveryLog(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "DeliveryLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete delivery_log", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateDeliveryLog godoc
// @Summary      배송 로그 정보 수정
// @Description  trip_id로 배송 로그 정보를 수정합니다.
// @Tags         delivery_log
// @Accept       json
// @Produce      json
// @Param        id           path      int                          true  "배송 로그 trip_id"
// @Param        delivery_log body      dto.UpdateDeliveryLogRequest true  "수정할 배송 로그 정보"
// @Success      200          {object}  dto.DeliveryLogResponse
// @Failure      400          {object}  dto.ErrorResponse
// @Failure      404          {object}  dto.ErrorResponse
// @Failure      500          {object}  dto.ErrorResponse
// @Router       /api/delivery-log/{id} [put]
func (h *DeliveryLogHandler) UpdateDeliveryLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid delivery_log id"})
		return
	}
	var req dto.UpdateDeliveryLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	log, err := h.service.UpdateDeliveryLog(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "DeliveryLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update delivery_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

// ListDeliveryLogs godoc
// @Summary      모든 배송 로그 조회
// @Description  모든 배송 로그 정보를 반환합니다.
// @Tags         delivery_log
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드 (예: -registration_time, -trip_id 등)"
// @Success      200   {array}   dto.DeliveryLogResponse
// @Router       /api/delivery-log [get]
func (h *DeliveryLogHandler) ListDeliveryLogs(c *gin.Context) {
	sortParam := c.Query("sort")
	logs, err := h.service.ListDeliveryLogs(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list delivery_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

// SearchDeliveryLogs godoc
// @Summary      배송 로그 검색
// @Description  쿼리 파라미터로 배송 로그를 검색합니다.
// @Tags         delivery_log
// @Produce      json
// @Param        trip_id               query     int     false  "trip_id"
// @Param        package_id            query     int     false  "package_id"
// @Param        region_id             query     string  false  "region_id"
// @Param        load_order            query     int     false  "load_order"
// @Param        sort                  query     string  false  "정렬 필드 (예: -registration_time, -trip_id 등)"
// @Success      200  {array}   dto.DeliveryLogResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Router       /api/delivery-log/search [get]
func (h *DeliveryLogHandler) SearchDeliveryLogs(c *gin.Context) {
	params := map[string]string{}
	for _, key := range []string{"trip_id", "package_id", "region_id", "load_order"} {
		if v := c.Query(key); v != "" {
			params[key] = v
		}
	}
	sortParam := c.Query("sort")
	logs, err := h.service.SearchDeliveryLogs(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search delivery_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
