package dto

type CreatePackageRequest struct {
	PackageType   string `json:"package_type" binding:"required"`
	RegionID      string `json:"region_id" binding:"required,len=3"`
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
