package handlers

import (
	"net/http"
	"strconv"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PackageHandler struct {
	service *service.PackageService
}

func NewPackageHandler(s *service.PackageService) *PackageHandler {
	return &PackageHandler{service: s}
}

// CreatePackage godoc
// @Summary      패키지 생성
// @Description  새로운 패키지를 생성합니다.
// @Tags         package
// @Accept       json
// @Produce      json
// @Param        package  body      dto.CreatePackageRequest  true  "패키지 정보"
// @Success      201      {object}  dto.PackageResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /api/package [post]
func (h *PackageHandler) CreatePackage(c *gin.Context) {
	var req dto.CreatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	pkg, err := h.service.CreatePackage(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create package", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pkg)
}

// GetPackageByID godoc
// @Summary      패키지 단건 조회
// @Description  패키지 ID로 패키지 정보를 조회합니다.
// @Tags         package
// @Produce      json
// @Param        id   path      int  true  "패키지 ID"
// @Success      200  {object}  dto.PackageResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/package/{id} [get]
func (h *PackageHandler) GetPackageByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid package id"})
		return
	}
	pkg, err := h.service.GetPackageByID(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Package not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to get package", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// DeletePackage godoc
// @Summary      패키지 삭제
// @Description  패키지 ID로 패키지를 삭제합니다.
// @Tags         package
// @Produce      json
// @Param        id   path      int  true  "패키지 ID"
// @Success      204  "No Content"
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/package/{id} [delete]
func (h *PackageHandler) DeletePackage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid package id"})
		return
	}
	err = h.service.DeletePackage(c.Request.Context(), id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Package not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete package", Details: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdatePackage godoc
// @Summary      패키지 정보 수정
// @Description  패키지 ID로 패키지 정보를 수정합니다.
// @Tags         package
// @Accept       json
// @Produce      json
// @Param        id      path      int                     true  "패키지 ID"
// @Param        package body      dto.UpdatePackageRequest true  "수정할 패키지 정보"
// @Success      200     {object}  dto.PackageResponse
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      404     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /api/package/{id} [put]
func (h *PackageHandler) UpdatePackage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid package id"})
		return
	}
	var req dto.UpdatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}
	pkg, err := h.service.UpdatePackage(c.Request.Context(), id, req)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Package not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update package", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// ListPackages godoc
// @Summary      모든 패키지 조회
// @Description  모든 패키지 정보를 반환합니다.
// @Tags         package
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드 (예: -registered_at는 최신순, package_id 등)"
// @Success      200   {array}   dto.PackageResponse
// @Router       /api/package [get]
func (h *PackageHandler) ListPackages(c *gin.Context) {
	sortParam := c.Query("sort")
	pkgs, err := h.service.ListPackages(c.Request.Context(), sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to list packages", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// SearchPackages godoc
// @Summary      패키지 검색
// @Description  쿼리 파라미터로 패키지를 검색합니다.
// @Tags         package
// @Produce      json
// @Param        package_id     query     int     false  "패키지 ID"
// @Param        package_type   query     string  false  "패키지 타입"
// @Param        region_id      query     string  false  "지역 ID"
// @Param        package_status query     string  false  "패키지 상태"
// @Param        registered_at  query     string  false  "등록 시각(RFC3339)"
// @Param        sort           query     string  false  "정렬 필드 (예: -registered_at, -package_id 등)"
// @Success      200  {array}   dto.PackageResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Router       /api/package/search [get]
func (h *PackageHandler) SearchPackages(c *gin.Context) {
	params := map[string]string{}
	for _, key := range []string{"package_id", "package_type", "region_id", "package_status", "registered_at"} {
		if v := c.Query(key); v != "" {
			params[key] = v
		}
	}
	sortParam := c.Query("sort")
	pkgs, err := h.service.SearchPackages(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search packages", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}
