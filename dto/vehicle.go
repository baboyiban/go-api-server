package dto

type CreateVehicleRequest struct {
	VehicleID string `json:"vehicle_id" binding:"required"`
	MaxLoad   int    `json:"max_load"`
}

type UpdateVehicleRequest struct {
	MaxLoad           int    `json:"max_load"`
	LedStatus         string `json:"led_status"`
	NeedsConfirmation bool   `json:"needs_confirmation"`
	CoordX            int    `json:"coord_x"`
	CoordY            int    `json:"coord_y"`
}

type VehicleResponse struct {
	InternalID        int    `json:"internal_id"`
	VehicleID         string `json:"vehicle_id"`
	CurrentLoad       int    `json:"current_load"`
	MaxLoad           int    `json:"max_load"`
	LedStatus         string `json:"led_status"`
	NeedsConfirmation bool   `json:"needs_confirmation"`
	CoordX            int    `json:"coord_x"`
	CoordY            int    `json:"coord_y"`
}
