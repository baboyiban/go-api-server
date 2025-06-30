package dto

import (
	"time"

	"github.com/baboyiban/go-api-server/models"
)

type ChangeNotifyResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"table_name"`
	Action    string    `json:"action"`
	ChangedAt time.Time `json:"changed_at"`
}

func ToChangeNotifyResponse(m *models.ChangeNotify) ChangeNotifyResponse {
	return ChangeNotifyResponse{
		ID:        m.ID,
		Name:      m.Name,
		Action:    m.Action,
		ChangedAt: m.ChangedAt,
	}
}
