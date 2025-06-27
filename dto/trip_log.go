package dto

import (
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/utils"
)

type CreateTripLogRequest struct {
	VehicleID   string  `json:"vehicle_id" binding:"required"`
	StartTime   *string `json:"start_time"` // RFC3339 string
	EndTime     *string `json:"end_time"`   // RFC3339 string
	Status      string  `json:"status"`     // "운행중" or "비운행중"
	Destination *string `json:"destination"`
}

type UpdateTripLogRequest struct {
	StartTime   *string `json:"start_time"`
	EndTime     *string `json:"end_time"`
	Status      string  `json:"status"`
	Destination *string `json:"destination"`
}

type TripLogResponse struct {
	TripID      int     `json:"trip_id"`
	VehicleID   string  `json:"vehicle_id"`
	StartTime   *string `json:"start_time"`
	EndTime     *string `json:"end_time"`
	Status      string  `json:"status"`
	Destination *string `json:"destination"`
}

func ToTripLogResponse(v *models.TripLog) TripLogResponse {
	return TripLogResponse{
		TripID:      v.TripID,
		VehicleID:   v.VehicleID,
		StartTime:   utils.FormatTimePtr(v.StartTime),
		EndTime:     utils.FormatTimePtr(v.EndTime),
		Status:      v.Status,
		Destination: v.Destination,
	}
}
