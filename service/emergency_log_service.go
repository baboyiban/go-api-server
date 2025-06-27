package service

import (
	"context"
	"time"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"gorm.io/gorm"
)

var allowedEmergencyLogSortFields = map[string]bool{
	"trip_id":            true,
	"vehicle_id":         true,
	"call_time":          true,
	"reason":             true,
	"employee_id":        true,
	"needs_confirmation": true,
}

func applyEmergencyLogSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedEmergencyLogSortFields[field] {
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

type EmergencyLogService struct {
	db *gorm.DB
}

func NewEmergencyLogService(db *gorm.DB) *EmergencyLogService {
	return &EmergencyLogService{db: db}
}

func (s *EmergencyLogService) CreateEmergencyLog(ctx context.Context, req dto.CreateEmergencyLogRequest) (*dto.EmergencyLogResponse, error) {
	em := models.EmergencyLog{
		TripID:     req.TripID,
		VehicleID:  req.VehicleID,
		Reason:     req.Reason,
		EmployeeID: req.EmployeeID,
		CallTime:   time.Now(),
	}
	if req.NeedsConfirmation != nil {
		em.NeedsConfirmation = *req.NeedsConfirmation
	}
	if err := s.db.WithContext(ctx).Create(&em).Error; err != nil {
		return nil, err
	}
	resp := dto.ToEmergencyLogResponse(&em)
	return &resp, nil
}

func (s *EmergencyLogService) GetEmergencyLogByID(ctx context.Context, id int) (*dto.EmergencyLogResponse, error) {
	var em models.EmergencyLog
	if err := s.db.WithContext(ctx).Where("trip_id = ?", id).First(&em).Error; err != nil {
		return nil, err
	}
	resp := dto.ToEmergencyLogResponse(&em)
	return &resp, nil
}

func (s *EmergencyLogService) UpdateEmergencyLog(ctx context.Context, id int, req dto.UpdateEmergencyLogRequest) (*dto.EmergencyLogResponse, error) {
	var em models.EmergencyLog
	if err := s.db.WithContext(ctx).Where("trip_id = ?", id).First(&em).Error; err != nil {
		return nil, err
	}
	if req.Reason != "" {
		em.Reason = req.Reason
	}
	if req.NeedsConfirmation != nil {
		em.NeedsConfirmation = *req.NeedsConfirmation
	}
	if err := s.db.WithContext(ctx).Save(&em).Error; err != nil {
		return nil, err
	}
	resp := dto.ToEmergencyLogResponse(&em)
	return &resp, nil
}

func (s *EmergencyLogService) DeleteEmergencyLog(ctx context.Context, id int) error {
	result := s.db.WithContext(ctx).Where("trip_id = ?", id).Delete(&models.EmergencyLog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *EmergencyLogService) ListEmergencyLogs(ctx context.Context, sort string) ([]dto.EmergencyLogResponse, error) {
	var logs []models.EmergencyLog
	query := s.db.WithContext(ctx).Model(&models.EmergencyLog{})
	query = applyEmergencyLogSort(query, sort)
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	res := make([]dto.EmergencyLogResponse, 0, len(logs))
	for i := range logs {
		res = append(res, dto.ToEmergencyLogResponse(&logs[i]))
	}
	return res, nil
}

func (s *EmergencyLogService) SearchEmergencyLogs(ctx context.Context, params map[string]string, sort string) ([]dto.EmergencyLogResponse, error) {
	var logs []models.EmergencyLog
	query := s.db.WithContext(ctx).Model(&models.EmergencyLog{})
	for k, v := range params {
		if k == "call_time" {
			query = query.Where("DATE("+k+") = ?", v)
		} else {
			query = query.Where(k+" = ?", v)
		}
	}
	query = applyEmergencyLogSort(query, sort)
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	res := make([]dto.EmergencyLogResponse, 0, len(logs))
	for i := range logs {
		res = append(res, dto.ToEmergencyLogResponse(&logs[i]))
	}
	return res, nil
}
