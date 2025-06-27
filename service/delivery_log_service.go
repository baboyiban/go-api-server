package service

import (
	"context"
	"time"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/utils"
	"gorm.io/gorm"
)

var allowedDeliveryLogSortFields = map[string]bool{
	"trip_id":               true,
	"package_id":            true,
	"region_id":             true,
	"load_order":            true,
	"registered_at":         true,
	"first_transport_time":  true,
	"input_time":            true,
	"second_transport_time": true,
	"completed_at":          true,
}

func applyDeliveryLogSort(query *gorm.DB, sort string) *gorm.DB {
	if sort == "" {
		return query
	}
	field := sort
	desc := false
	if sort[0] == '-' {
		field = sort[1:]
		desc = true
	}
	if allowedDeliveryLogSortFields[field] {
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

type DeliveryLogService struct {
	db *gorm.DB
}

func NewDeliveryLogService(db *gorm.DB) *DeliveryLogService {
	return &DeliveryLogService{db: db}
}

func (s *DeliveryLogService) CreateDeliveryLog(ctx context.Context, req dto.CreateDeliveryLogRequest) (*dto.DeliveryLogResponse, error) {
	log := models.DeliveryLog{
		TripID:              req.TripID,
		PackageID:           req.PackageID,
		RegionID:            req.RegionID,
		LoadOrder:           req.LoadOrder,
		RegisteredAt:        time.Now(),
		FirstTransportTime:  utils.ParseTimePtr(req.FirstTransportTime),
		InputTime:           utils.ParseTimePtr(req.InputTime),
		SecondTransportTime: utils.ParseTimePtr(req.SecondTransportTime),
		CompletedAt:         utils.ParseTimePtr(req.CompletedAt),
	}
	if req.RegisteredAt != nil {
		if t := utils.ParseTimePtr(req.RegisteredAt); t != nil {
			log.RegisteredAt = *t
		}
	}
	if err := s.db.WithContext(ctx).Create(&log).Error; err != nil {
		return nil, err
	}
	resp := dto.ToDeliveryLogResponse(&log)
	return &resp, nil
}

func (s *DeliveryLogService) GetDeliveryLogByID(ctx context.Context, tripID int) (*dto.DeliveryLogResponse, error) {
	var log models.DeliveryLog
	if err := s.db.WithContext(ctx).Where("trip_id = ?", tripID).First(&log).Error; err != nil {
		return nil, err
	}
	resp := dto.ToDeliveryLogResponse(&log)
	return &resp, nil
}

func (s *DeliveryLogService) DeleteDeliveryLog(ctx context.Context, tripID int) error {
	result := s.db.WithContext(ctx).Where("trip_id = ?", tripID).Delete(&models.DeliveryLog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *DeliveryLogService) UpdateDeliveryLog(ctx context.Context, tripID int, req dto.UpdateDeliveryLogRequest) (*dto.DeliveryLogResponse, error) {
	var log models.DeliveryLog
	if err := s.db.WithContext(ctx).Where("trip_id = ?", tripID).First(&log).Error; err != nil {
		return nil, err
	}
	log.LoadOrder = req.LoadOrder
	if req.RegisteredAt != nil {
		if t := utils.ParseTimePtr(req.RegisteredAt); t != nil {
			log.RegisteredAt = *t
		}
	}
	log.FirstTransportTime = utils.ParseTimePtr(req.FirstTransportTime)
	log.InputTime = utils.ParseTimePtr(req.InputTime)
	log.SecondTransportTime = utils.ParseTimePtr(req.SecondTransportTime)
	log.CompletedAt = utils.ParseTimePtr(req.CompletedAt)
	if err := s.db.WithContext(ctx).Save(&log).Error; err != nil {
		return nil, err
	}
	resp := dto.ToDeliveryLogResponse(&log)
	return &resp, nil
}

func (s *DeliveryLogService) ListDeliveryLogs(ctx context.Context, sort string) ([]dto.DeliveryLogResponse, error) {
	var logs []models.DeliveryLog
	query := s.db.WithContext(ctx).Model(&models.DeliveryLog{})
	query = applyDeliveryLogSort(query, sort)
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	res := make([]dto.DeliveryLogResponse, 0, len(logs))
	for i := range logs {
		res = append(res, dto.ToDeliveryLogResponse(&logs[i]))
	}
	return res, nil
}

func (s *DeliveryLogService) SearchDeliveryLogs(ctx context.Context, params map[string]string, sort string) ([]dto.DeliveryLogResponse, error) {
	var logs []models.DeliveryLog
	query := s.db.WithContext(ctx).Model(&models.DeliveryLog{})

	dateFields := map[string]bool{
		"registered_at":         true,
		"first_transport_time":  true,
		"input_time":            true,
		"second_transport_time": true,
		"completed_at":          true,
	}

	for k, v := range params {
		if dateFields[k] {
			query = query.Where("DATE("+k+") = ?", v)
		} else {
			query = query.Where(k+" = ?", v)
		}
	}

	query = applyDeliveryLogSort(query, sort)
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	res := make([]dto.DeliveryLogResponse, 0, len(logs))
	for i := range logs {
		res = append(res, dto.ToDeliveryLogResponse(&logs[i]))
	}
	return res, nil
}
