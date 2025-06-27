package dto

import (
	"time"

	"github.com/baboyiban/go-api-server/models"
)

type CreatePackageRequest struct {
	PackageType   string `json:"package_type" binding:"required"`
	RegionID      string `json:"region_id" binding:"required"`
	PackageStatus string `json:"package_status"`
}

type UpdatePackageRequest struct {
	PackageType   string `json:"package_type"`
	RegionID      string `json:"region_id"`
	PackageStatus string `json:"package_status"`
}

type PackageResponse struct {
	PackageID     int    `json:"package_id"`
	PackageType   string `json:"package_type"`
	RegionID      string `json:"region_id"`
	PackageStatus string `json:"package_status"`
	RegisteredAt  string `json:"registered_at"`
}

func ToPackageResponse(pkg *models.Package) PackageResponse {
	return PackageResponse{
		PackageID:     pkg.PackageID,
		PackageType:   pkg.PackageType,
		RegionID:      pkg.RegionID,
		PackageStatus: pkg.PackageStatus,
		RegisteredAt:  pkg.RegisteredAt.Format(time.RFC3339),
	}
}
