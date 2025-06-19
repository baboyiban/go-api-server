package service

import (
	"context"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/utils"
	"gorm.io/gorm"
)

var allowedTripLogBSortFields = map[string]bool{
	"trip_id":       true,
	"vehicle_id":    true,
	"start_time":    true,
	"end_time":      true,
	"status":        true,
	"destination_1": true,
	"destination_2": true,
	"destination_3": true,
}

func applyTripLogBSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedTripLogBSortFields[field] {
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

type TripLogBService struct {
	db *gorm.DB
}

func NewTripLogBService(db *gorm.DB) *TripLogBService {
	return &TripLogBService{db: db}
}

func (s *TripLogBService) CreateTripLogB(ctx context.Context, req dto.CreateTripLogBRequest) (*dto.TripLogBResponse, error) {
	trip := models.TripLogB{
		VehicleID:    req.VehicleID,
		StartTime:    utils.ParseTimePtr(req.StartTime),
		EndTime:      utils.ParseTimePtr(req.EndTime),
		Status:       req.Status,
		Destination1: req.Destination1,
		Destination2: req.Destination2,
		Destination3: req.Destination3,
	}
	if trip.Status == "" {
		trip.Status = "비운행중"
	}
	if err := s.db.WithContext(ctx).Create(&trip).Error; err != nil {
		return nil, err
	}
	return toTripLogBResponse(&trip), nil
}

func (s *TripLogBService) GetTripLogBByID(ctx context.Context, id int) (*dto.TripLogBResponse, error) {
	var trip models.TripLogB
	if err := s.db.WithContext(ctx).Where("trip_id = ?", id).First(&trip).Error; err != nil {
		return nil, err
	}
	return toTripLogBResponse(&trip), nil
}

func (s *TripLogBService) DeleteTripLogB(ctx context.Context, id int) error {
	result := s.db.WithContext(ctx).Where("trip_id = ?", id).Delete(&models.TripLogB{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *TripLogBService) UpdateTripLogB(ctx context.Context, id int, req dto.UpdateTripLogBRequest) (*dto.TripLogBResponse, error) {
	var trip models.TripLogB
	if err := s.db.WithContext(ctx).Where("trip_id = ?", id).First(&trip).Error; err != nil {
		return nil, err
	}
	trip.StartTime = utils.ParseTimePtr(req.StartTime)
	trip.EndTime = utils.ParseTimePtr(req.EndTime)
	if req.Status != "" {
		trip.Status = req.Status
	}
	trip.Destination1 = req.Destination1
	trip.Destination2 = req.Destination2
	trip.Destination3 = req.Destination3
	if err := s.db.WithContext(ctx).Save(&trip).Error; err != nil {
		return nil, err
	}
	return toTripLogBResponse(&trip), nil
}

func (s *TripLogBService) ListTripLogBs(ctx context.Context, sort string) ([]dto.TripLogBResponse, error) {
	var trips []models.TripLogB
	query := s.db.WithContext(ctx).Model(&models.TripLogB{})
	query = applyTripLogBSort(query, sort)
	if err := query.Find(&trips).Error; err != nil {
		return nil, err
	}
	var res []dto.TripLogBResponse
	for _, t := range trips {
		res = append(res, *toTripLogBResponse(&t))
	}
	return res, nil
}

func (s *TripLogBService) SearchTripLogBs(ctx context.Context, params map[string]string, sort string) ([]dto.TripLogBResponse, error) {
	var trips []models.TripLogB
	query := s.db.WithContext(ctx).Model(&models.TripLogB{})
	for k, v := range params {
		query = query.Where(k+" = ?", v)
	}
	query = applyTripLogBSort(query, sort)
	if err := query.Find(&trips).Error; err != nil {
		return nil, err
	}
	var res []dto.TripLogBResponse
	for _, t := range trips {
		res = append(res, *toTripLogBResponse(&t))
	}
	return res, nil
}

func toTripLogBResponse(m *models.TripLogB) *dto.TripLogBResponse {
	return &dto.TripLogBResponse{
		TripID:       m.TripID,
		VehicleID:    m.VehicleID,
		StartTime:    utils.FormatTimePtr(m.StartTime),
		EndTime:      utils.FormatTimePtr(m.EndTime),
		Status:       m.Status,
		Destination1: m.Destination1,
		Destination2: m.Destination2,
		Destination3: m.Destination3,
	}
}
