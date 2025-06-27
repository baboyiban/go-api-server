package service

import (
	"context"
	"time"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"gorm.io/gorm"
)

var allowedSortFields = map[string]bool{
	"region_id":        true,
	"region_name":      true,
	"coord_x":          true,
	"coord_y":          true,
	"max_capacity":     true,
	"current_capacity": true,
	"is_full":          true,
	"saturated_at":     true,
}

func applyRegionSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedSortFields[field] {
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

type RegionService struct {
	db *gorm.DB
}

func NewRegionService(db *gorm.DB) *RegionService {
	return &RegionService{db: db}
}

func (s *RegionService) CreateRegion(ctx context.Context, req dto.CreateRegionRequest) (*dto.RegionResponse, error) {
	region := models.Region{
		RegionID:    req.RegionID,
		RegionName:  req.RegionName,
		CoordX:      req.CoordX,
		CoordY:      req.CoordY,
		MaxCapacity: req.MaxCapacity,
	}
	if err := s.db.WithContext(ctx).Create(&region).Error; err != nil {
		return nil, err
	}
	resp := dto.ToRegionResponse(&region)
	return &resp, nil
}

func (s *RegionService) GetRegionByID(ctx context.Context, id string) (*dto.RegionResponse, error) {
	var region models.Region
	if err := s.db.WithContext(ctx).Where("region_id = ?", id).First(&region).Error; err != nil {
		return nil, err
	}
	resp := dto.ToRegionResponse(&region)
	return &resp, nil
}

func (s *RegionService) DeleteRegion(ctx context.Context, id string) error {
	result := s.db.WithContext(ctx).Where("region_id = ?", id).Delete(&models.Region{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *RegionService) UpdateRegion(ctx context.Context, id string, req dto.UpdateRegionRequest) (*dto.RegionResponse, error) {
	var region models.Region
	if err := s.db.WithContext(ctx).Where("region_id = ?", id).First(&region).Error; err != nil {
		return nil, err
	}
	region.RegionName = req.RegionName
	region.CoordX = req.CoordX
	region.CoordY = req.CoordY
	region.MaxCapacity = req.MaxCapacity
	region.CurrentCapacity = req.CurrentCapacity
	region.IsFull = req.IsFull
	if req.SaturatedAt != nil {
		t, _ := time.Parse(time.RFC3339, *req.SaturatedAt)
		region.SaturatedAt = &t
	}
	if err := s.db.WithContext(ctx).Save(&region).Error; err != nil {
		return nil, err
	}
	resp := dto.ToRegionResponse(&region)
	return &resp, nil
}

func (s *RegionService) ListRegions(ctx context.Context, sort string) ([]dto.RegionResponse, error) {
	var regions []models.Region
	query := s.db.WithContext(ctx).Model(&models.Region{})
	query = applyRegionSort(query, sort)
	if err := query.Find(&regions).Error; err != nil {
		return nil, err
	}
	res := make([]dto.RegionResponse, 0, len(regions))
	for i := range regions {
		res = append(res, dto.ToRegionResponse(&regions[i]))
	}
	return res, nil
}

func (s *RegionService) SearchRegions(ctx context.Context, params map[string]string, sort string) ([]dto.RegionResponse, error) {
	var regions []models.Region
	query := s.db.WithContext(ctx).Model(&models.Region{})
	for k, v := range params {
		if k == "saturated_at" {
			query = query.Where("DATE("+k+") = ?", v)
		} else {
			query = query.Where(k+" = ?", v)
		}
	}
	query = applyRegionSort(query, sort)
	if err := query.Find(&regions).Error; err != nil {
		return nil, err
	}
	res := make([]dto.RegionResponse, 0, len(regions))
	for i := range regions {
		res = append(res, dto.ToRegionResponse(&regions[i]))
	}
	return res, nil
}
