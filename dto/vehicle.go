package dto

import "github.com/baboyiban/go-api-server/models"

type CreateVehicleRequest struct {
	VehicleID string `json:"vehicle_id" binding:"required"`
	MaxLoad   int    `json:"max_load"`
	CoordX    *int   `json:"coord_x"`
	CoordY    *int   `json:"coord_y"`
	AICoordX  *int   `json:"AI_coord_x"`
	AICoordY  *int   `json:"AI_coord_y"`
}

type UpdateVehicleRequest struct {
	MaxLoad           int    `json:"max_load"`
	LedStatus         string `json:"led_status"`
	NeedsConfirmation bool   `json:"needs_confirmation"`
	CoordX            *int   `json:"coord_x"`
	CoordY            *int   `json:"coord_y"`
	AICoordX          *int   `json:"AI_coord_x"`
	AICoordY          *int   `json:"AI_coord_y"`
}

type VehicleResponse struct {
	InternalID        int    `json:"internal_id"`
	VehicleID         string `json:"vehicle_id"`
	CurrentLoad       int    `json:"current_load"`
	MaxLoad           int    `json:"max_load"`
	LedStatus         string `json:"led_status"`
	NeedsConfirmation bool   `json:"needs_confirmation"`
	CoordX            *int   `json:"coord_x"`
	CoordY            *int   `json:"coord_y"`
	AICoordX          *int   `json:"AI_coord_x"`
	AICoordY          *int   `json:"AI_coord_y"`
}

func ToVehicleResponse(v *models.Vehicle) VehicleResponse {
	return VehicleResponse{
		InternalID:        v.InternalID,
		VehicleID:         v.VehicleID,
		CurrentLoad:       v.CurrentLoad,
		MaxLoad:           v.MaxLoad,
		LedStatus:         v.LedStatus,
		NeedsConfirmation: v.NeedsConfirmation,
		CoordX:            v.CoordX,
		CoordY:            v.CoordY,
		AICoordX:          v.AICoordX,
		AICoordY:          v.AICoordY,
	}
}
