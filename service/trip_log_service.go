package service

import (
	"context"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/utils"
	"gorm.io/gorm"
)

var allowedTripLogSortFields = map[string]bool{
	"trip_id":     true,
	"vehicle_id":  true,
	"start_time":  true,
	"end_time":    true,
	"status":      true,
	"destination": true,
}

func applyTripLogSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedTripLogSortFields[field] {
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

type TripLogService struct {
	db *gorm.DB
}

func NewTripLogService(db *gorm.DB) *TripLogService {
	return &TripLogService{db: db}
}

func (s *TripLogService) CreateTripLog(ctx context.Context, req dto.CreateTripLogRequest) (*dto.TripLogResponse, error) {
	trip := models.TripLog{
		VehicleID:   req.VehicleID,
		StartTime:   utils.ParseTimePtr(req.StartTime),
		EndTime:     utils.ParseTimePtr(req.EndTime),
		Status:      req.Status,
		Destination: req.Destination,
	}
	if trip.Status == "" {
		trip.Status = "비운행중"
	}
	if err := s.db.WithContext(ctx).Create(&trip).Error; err != nil {
		return nil, err
	}
	return toTripLogResponse(&trip), nil
}

func (s *TripLogService) GetTripLogByID(ctx context.Context, id int) (*dto.TripLogResponse, error) {
	var trip models.TripLog
	if err := s.db.WithContext(ctx).Where("trip_id = ?", id).First(&trip).Error; err != nil {
		return nil, err
	}
	return toTripLogResponse(&trip), nil
}

func (s *TripLogService) DeleteTripLog(ctx context.Context, id int) error {
	result := s.db.WithContext(ctx).Where("trip_id = ?", id).Delete(&models.TripLog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *TripLogService) UpdateTripLog(ctx context.Context, id int, req dto.UpdateTripLogRequest) (*dto.TripLogResponse, error) {
	var trip models.TripLog
	if err := s.db.WithContext(ctx).Where("trip_id = ?", id).First(&trip).Error; err != nil {
		return nil, err
	}
	trip.StartTime = utils.ParseTimePtr(req.StartTime)
	trip.EndTime = utils.ParseTimePtr(req.EndTime)
	if req.Status != "" {
		trip.Status = req.Status
	}
	trip.Destination = req.Destination
	if err := s.db.WithContext(ctx).Save(&trip).Error; err != nil {
		return nil, err
	}
	return toTripLogResponse(&trip), nil
}

func (s *TripLogService) ListTripLogs(ctx context.Context, sort string) ([]dto.TripLogResponse, error) {
	var trips []models.TripLog
	query := s.db.WithContext(ctx).Model(&models.TripLog{})
	query = applyTripLogSort(query, sort)
	if err := query.Find(&trips).Error; err != nil {
		return nil, err
	}
	var res []dto.TripLogResponse
	for _, t := range trips {
		res = append(res, *toTripLogResponse(&t))
	}
	return res, nil
}

func (s *TripLogService) SearchTripLogs(ctx context.Context, params map[string]string, sort string) ([]dto.TripLogResponse, error) {
	var trips []models.TripLog
	query := s.db.WithContext(ctx).Model(&models.TripLog{})
	for k, v := range params {
		if k == "start_time" || k == "end_time" {
			query = query.Where("DATE("+k+") = ?", v)
		} else {
			query = query.Where(k+" = ?", v)
		}
	}
	query = applyTripLogSort(query, sort)
	if err := query.Find(&trips).Error; err != nil {
		return nil, err
	}
	var res []dto.TripLogResponse
	for _, t := range trips {
		res = append(res, *toTripLogResponse(&t))
	}
	return res, nil
}

func toTripLogResponse(m *models.TripLog) *dto.TripLogResponse {
	return &dto.TripLogResponse{
		TripID:      m.TripID,
		VehicleID:   m.VehicleID,
		StartTime:   utils.FormatTimePtr(m.StartTime),
		EndTime:     utils.FormatTimePtr(m.EndTime),
		Status:      m.Status,
		Destination: m.Destination,
	}
}
