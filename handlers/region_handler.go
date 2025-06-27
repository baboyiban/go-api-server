package handlers

import (
	"net/http"

	"github.com/baboyiban/go-api-server/constants"
	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegionHandler struct {
	service *service.RegionService
}

func NewRegionHandler(s *service.RegionService) *RegionHandler {
	return &RegionHandler{service: s}
}

// CreateRegion godoc
// @Summary      지역 생성
// @Description  새로운 지역을 생성합니다.
// @Tags         region
// @Accept       json
// @Produce      json
// @Param        region  body      dto.CreateRegionRequest  true  "지역 정보"
// @Success      201     {object}  dto.RegionResponse
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /api/region [post]
func (h *RegionHandler) CreateRegion(c *gin.Context) {
	var req dto.CreateRegionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	region, err := h.service.CreateRegion(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create region", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, region)
}

// GetRegionByID godoc
// @Summary      지역 단건 조회
// @Description  region_id로 지역을 조회합니다.
// @Tags         region
// @Produce      json
// @Param        id   path      string  true  "지역 ID"
// @Success      200  {object}  dto.RegionResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/region/{id} [get]
func (h *RegionHandler) GetRegionByID(c *gin.Context) {
	id := c.Param("id")
	region, err := h.service.GetRegionByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Region not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get region", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, region)
}

// DeleteRegion godoc
// @Summary      지역 삭제
// @Description  region_id로 지역을 삭제합니다.
// @Tags         region
// @Produce      json
// @Param        id   path      string  true  "지역 ID"
// @Success      204  "No Content"
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/region/{id} [delete]
func (h *RegionHandler) DeleteRegion(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteRegion(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Region not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete region", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateRegion godoc
// @Summary      지역 정보 수정
// @Description  region_id로 지역 정보를 수정합니다.
// @Tags         region
// @Accept       json
// @Produce      json
// @Param        id      path      string                   true  "지역 ID"
// @Param        region  body      dto.UpdateRegionRequest  true  "수정할 지역 정보"
// @Success      200     {object}  dto.RegionResponse
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      404     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /api/region/{id} [put]
func (h *RegionHandler) UpdateRegion(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateRegionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	region, err := h.service.UpdateRegion(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Region not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update region", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, region)
}

// ListRegions godoc
// @Summary      모든 지역 조회
// @Description  모든 지역 정보를 반환합니다.
// @Tags         region
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드"
// @Success      200   {array}   dto.RegionResponse
// @Router       /api/region [get]
func (h *RegionHandler) ListRegions(c *gin.Context) {
	sortParam := c.Query("sort")
	regions, err := h.service.ListRegions(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list regions", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, regions)
}

// SearchRegions godoc
// @Summary      지역 검색
// @Description  쿼리 파라미터로 지역을 검색합니다.
// @Tags         region
// @Produce      json
// @Param        region_id   query     string  false  "지역 ID"
// @Param        name        query     string  false  "지역명"
// @Param        sort        query     string  false  "정렬 필드"
// @Success      200  {array}   dto.RegionResponse
// @Router       /api/region/search [get]
func (h *RegionHandler) SearchRegions(c *gin.Context) {
	params := dto.ExtractAllowedParams(c.Request.URL.Query(), constants.RegionAllowedFields)
	sortParam := c.Query("sort")
	regions, err := h.service.SearchRegions(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search regions", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, regions)
}
