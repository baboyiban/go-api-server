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

type PackageHandler struct {
	service *service.PackageService
}

func NewPackageHandler(s *service.PackageService) *PackageHandler {
	return &PackageHandler{service: s}
}

// CreatePackage godoc
// @Summary      택배 생성
// @Description  새로운 택배를 생성합니다.
// @Tags         package
// @Accept       json
// @Produce      json
// @Param        package  body      dto.CreatePackageRequest  true  "택배 정보"
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
// @Summary      택배 단건 조회
// @Description  package_id로 택배를 조회합니다.
// @Tags         package
// @Produce      json
// @Param        id   path      int  true  "택배 ID"
// @Success      200  {object}  dto.PackageResponse
// @Failure      400  {object}  dto.ErrorResponse
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
// @Summary      택배 삭제
// @Description  package_id로 택배를 삭제합니다.
// @Tags         package
// @Produce      json
// @Param        id   path      int  true  "택배 ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.ErrorResponse
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
// @Summary      택배 정보 수정
// @Description  package_id로 택배 정보를 수정합니다.
// @Tags         package
// @Accept       json
// @Produce      json
// @Param        id       path      int                      true  "택배 ID"
// @Param        package  body      dto.UpdatePackageRequest  true  "수정할 택배 정보"
// @Success      200      {object}  dto.PackageResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      404      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
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
// @Summary      모든 택배 조회
// @Description  모든 택배 정보를 반환합니다.
// @Tags         package
// @Produce      json
// @Param        sort  query     string  false  "정렬 필드"
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
// @Summary      택배 검색
// @Description  쿼리 파라미터로 택배를 검색합니다.
// @Tags         package
// @Produce      json
// @Param        package_id   query     int     false  "택배 ID"
// @Param        recipient    query     string  false  "수령인"
// @Param        status       query     string  false  "상태"
// @Param        sort         query     string  false  "정렬 필드"
// @Success      200  {array}   dto.PackageResponse
// @Router       /api/package/search [get]
func (h *PackageHandler) SearchPackages(c *gin.Context) {
	params := dto.ExtractAllowedParams(c.Request.URL.Query(), constants.PackageAllowedFields)
	sortParam := c.Query("sort")
	pkgs, err := h.service.SearchPackages(c.Request.Context(), params, sortParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to search packages", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}
