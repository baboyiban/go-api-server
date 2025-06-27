package dto

import "github.com/baboyiban/go-api-server/models"

type CreateEmergencyLogRequest struct {
	TripID            int    `json:"trip_id" binding:"required"`
	VehicleID         string `json:"vehicle_id" binding:"required"`
	Reason            string `json:"reason" binding:"required,oneof=차량 관련 호출 택배 관련 호출 운송 관련 호출"`
	EmployeeID        int    `json:"employee_id" binding:"required"`
	NeedsConfirmation *bool  `json:"needs_confirmation"`
}

type UpdateEmergencyLogRequest struct {
	Reason            string `json:"reason"`
	NeedsConfirmation *bool  `json:"needs_confirmation"`
}

type EmergencyLogResponse struct {
	TripID            int    `json:"trip_id"`
	VehicleID         string `json:"vehicle_id"`
	CallTime          string `json:"call_time"`
	Reason            string `json:"reason"`
	EmployeeID        int    `json:"employee_id"`
	NeedsConfirmation bool   `json:"needs_confirmation"`
}

func ToEmergencyLogResponse(m *models.EmergencyLog) EmergencyLogResponse {
	return EmergencyLogResponse{
		TripID:            m.TripID,
		VehicleID:         m.VehicleID,
		CallTime:          m.CallTime.Format("2006-01-02T15:04:05Z07:00"),
		Reason:            m.Reason,
		EmployeeID:        m.EmployeeID,
		NeedsConfirmation: m.NeedsConfirmation,
	}
}
