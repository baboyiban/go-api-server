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

type EmergencyLogHandler struct {
	service *service.EmergencyLogService
}

func NewEmergencyLogHandler(s *service.EmergencyLogService) *EmergencyLogHandler {
	return &EmergencyLogHandler{service: s}
}

// CreateEmergencyLog
func (h *EmergencyLogHandler) CreateEmergencyLog(c *gin.Context) {
	var req dto.CreateEmergencyLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	log, err := h.service.CreateEmergencyLog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create emergency_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, log)
}

// GetEmergencyLogByID
func (h *EmergencyLogHandler) GetEmergencyLogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid emergency_log id"})
		return
	}
	log, err := h.service.GetEmergencyLogByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "EmergencyLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get emergency_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

// UpdateEmergencyLog
func (h *EmergencyLogHandler) UpdateEmergencyLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid emergency_log id"})
		return
	}
	var req dto.UpdateEmergencyLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	log, err := h.service.UpdateEmergencyLog(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "EmergencyLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update emergency_log", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

// DeleteEmergencyLog
func (h *EmergencyLogHandler) DeleteEmergencyLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid emergency_log id"})
		return
	}
	err = h.service.DeleteEmergencyLog(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "EmergencyLog not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete emergency_log", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// ListEmergencyLogs
func (h *EmergencyLogHandler) ListEmergencyLogs(c *gin.Context) {
	sortParam := c.Query("sort")
	logs, err := h.service.ListEmergencyLogs(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list emergency_logs", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

// SearchEmergencyLogs
func (h *EmergencyLogHandler) SearchEmergencyLogs(c *gin.Context) {
	params := dto.ExtractAllowedParams(c.Request.URL.Query(), constants.EmergencyLogAllowedFields)
	sortParam := c.Query("sort")
	logs, err := h.service.SearchEmergencyLogs(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search emergency_logs", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
