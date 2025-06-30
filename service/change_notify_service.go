package service

import (
	"context"

	"github.com/baboyiban/go-api-server/dto"
	"github.com/baboyiban/go-api-server/models"
	"gorm.io/gorm"
)

type ChangeNotifyService struct {
	db *gorm.DB
}

func NewChangeNotifyService(db *gorm.DB) *ChangeNotifyService {
	return &ChangeNotifyService{db: db}
}

func (s *ChangeNotifyService) ListChangeNotifies(ctx context.Context) ([]dto.ChangeNotifyResponse, error) {
	var notifies []models.ChangeNotify
	if err := s.db.WithContext(ctx).Order("changed_at DESC").Find(&notifies).Error; err != nil {
		return nil, err
	}
	res := make([]dto.ChangeNotifyResponse, 0, len(notifies))
	for i := range notifies {
		res = append(res, dto.ToChangeNotifyResponse(&notifies[i]))
	}
	return res, nil
}
