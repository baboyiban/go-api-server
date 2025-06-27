package service

import (
	"context"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"gorm.io/gorm"
)

var allowedVehicleSortFields = map[string]bool{
	"internal_id":        true,
	"vehicle_id":         true,
	"current_load":       true,
	"max_load":           true,
	"led_status":         true,
	"needs_confirmation": true,
	"coord_x":            true,
	"coord_y":            true,
}

func applyVehicleSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedVehicleSortFields[field] {
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

type VehicleService struct {
	db *gorm.DB
}

func NewVehicleService(db *gorm.DB) *VehicleService {
	return &VehicleService{db: db}
}

func (s *VehicleService) CreateVehicle(ctx context.Context, req dto.CreateVehicleRequest) (*models.Vehicle, error) {
	vehicle := models.Vehicle{
		VehicleID: req.VehicleID,
		MaxLoad:   req.MaxLoad,
		CoordX:    req.CoordX,
		CoordY:    req.CoordY,
		AICoordX:  req.AICoordX,
		AICoordY:  req.AICoordY,
	}
	if err := s.db.WithContext(ctx).Create(&vehicle).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (s *VehicleService) GetVehicleByID(ctx context.Context, id int) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	if err := s.db.WithContext(ctx).Where("internal_id = ?", id).First(&vehicle).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (s *VehicleService) DeleteVehicle(ctx context.Context, id int) error {
	result := s.db.WithContext(ctx).Where("internal_id = ?", id).Delete(&models.Vehicle{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *VehicleService) UpdateVehicle(ctx context.Context, id int, req dto.UpdateVehicleRequest) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	if err := s.db.WithContext(ctx).Where("internal_id = ?", id).First(&vehicle).Error; err != nil {
		return nil, err
	}
	vehicle.MaxLoad = req.MaxLoad
	vehicle.LedStatus = req.LedStatus
	vehicle.NeedsConfirmation = req.NeedsConfirmation
	vehicle.CoordX = req.CoordX
	vehicle.CoordY = req.CoordY
	if err := s.db.WithContext(ctx).Save(&vehicle).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (s *VehicleService) ListVehicles(ctx context.Context, sort string) ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	query := s.db.WithContext(ctx).Model(&models.Vehicle{})
	query = applyVehicleSort(query, sort)
	if err := query.Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (s *VehicleService) SearchVehicles(ctx context.Context, params map[string]string, sort string) ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	query := s.db.WithContext(ctx).Model(&models.Vehicle{})
	for k, v := range params {
		query = query.Where(k+" = ?", v)
	}
	query = applyVehicleSort(query, sort)
	if err := query.Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}
