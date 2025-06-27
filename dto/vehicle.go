package dto

type CreateVehicleRequest struct {
	VehicleID string `json:"vehicle_id" binding:"required"`
	MaxLoad   int    `json:"max_load"`
	CoordX    *int   `json:"coord_x,omitempty"`
	CoordY    *int   `json:"coord_y,omitempty"`
	AICoordX  *int   `json:"AI_coord_x,omitempty"`
	AICoordY  *int   `json:"AI_coord_y,omitempty"`
}

type UpdateVehicleRequest struct {
	MaxLoad           int    `json:"max_load"`
	LedStatus         string `json:"led_status"`
	NeedsConfirmation bool   `json:"needs_confirmation"`
	CoordX            *int   `json:"coord_x,omitempty"`
	CoordY            *int   `json:"coord_y,omitempty"`
	AICoordX          *int   `json:"AI_coord_x,omitempty"`
	AICoordY          *int   `json:"AI_coord_y,omitempty"`
}

type VehicleResponse struct {
	InternalID        int    `json:"internal_id"`
	VehicleID         string `json:"vehicle_id"`
	CurrentLoad       int    `json:"current_load"`
	MaxLoad           int    `json:"max_load"`
	LedStatus         string `json:"led_status"`
	NeedsConfirmation bool   `json:"needs_confirmation"`
	CoordX            *int   `json:"coord_x,omitempty"`
	CoordY            *int   `json:"coord_y,omitempty"`
	AICoordX          *int   `json:"AI_coord_x,omitempty"`
	AICoordY          *int   `json:"AI_coord_y,omitempty"`
}
