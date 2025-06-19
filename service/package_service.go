package service

import (
	"context"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"gorm.io/gorm"
)

var allowedPackageSortFields = map[string]bool{
	"package_id":     true,
	"package_type":   true,
	"region_id":      true,
	"package_status": true,
	"registered_at":  true,
}

func applyPackageSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedPackageSortFields[field] {
		order := field
		if desc {
			order += " DESC"
		} else {
			order += " ASC"
		}
		return query.Order(order)
	}
	return query
}

type PackageService struct {
	db *gorm.DB
}

func NewPackageService(db *gorm.DB) *PackageService {
	return &PackageService{db: db}
}

func (s *PackageService) CreatePackage(ctx context.Context, req dto.CreatePackageRequest) (*models.Package, error) {
	pkg := models.Package{
		PackageType:   req.PackageType,
		RegionID:      req.RegionID,
		PackageStatus: req.PackageStatus,
	}
	if err := s.db.WithContext(ctx).Create(&pkg).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (s *PackageService) GetPackageByID(ctx context.Context, id int) (*models.Package, error) {
	var pkg models.Package
	if err := s.db.WithContext(ctx).Where("package_id = ?", id).First(&pkg).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (s *PackageService) DeletePackage(ctx context.Context, id int) error {
	result := s.db.WithContext(ctx).Where("package_id = ?", id).Delete(&models.Package{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *PackageService) UpdatePackage(ctx context.Context, id int, req dto.UpdatePackageRequest) (*models.Package, error) {
	var pkg models.Package
	if err := s.db.WithContext(ctx).Where("package_id = ?", id).First(&pkg).Error; err != nil {
		return nil, err
	}
	if req.PackageType != "" {
		pkg.PackageType = req.PackageType
	}
	if req.RegionID != "" {
		pkg.RegionID = req.RegionID
	}
	if req.PackageStatus != "" {
		pkg.PackageStatus = req.PackageStatus
	}
	if err := s.db.WithContext(ctx).Save(&pkg).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (s *PackageService) ListPackages(ctx context.Context, sort string) ([]models.Package, error) {
	var pkgs []models.Package
	query := s.db.WithContext(ctx).Model(&models.Package{})
	query = applyPackageSort(query, sort)
	if err := query.Find(&pkgs).Error; err != nil {
		return nil, err
	}
	return pkgs, nil
}

func (s *PackageService) SearchPackages(ctx context.Context, params map[string]string, sort string) ([]models.Package, error) {
	var pkgs []models.Package
	query := s.db.WithContext(ctx).Model(&models.Package{})
	for k, v := range params {
		query = query.Where(k+" = ?", v)
	}
	query = applyPackageSort(query, sort)
	if err := query.Find(&pkgs).Error; err != nil {
		return nil, err
	}
	return pkgs, nil
}
