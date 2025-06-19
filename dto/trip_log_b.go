package dto

type CreateTripLogBRequest struct {
	VehicleID    string  `json:"vehicle_id" binding:"required"`
	StartTime    *string `json:"start_time"` // RFC3339 string
	EndTime      *string `json:"end_time"`   // RFC3339 string
	Status       string  `json:"status"`     // "운행중" or "비운행중"
	Destination1 *string `json:"destination_1"`
	Destination2 *string `json:"destination_2"`
	Destination3 *string `json:"destination_3"`
}

type UpdateTripLogBRequest struct {
	StartTime    *string `json:"start_time"`
	EndTime      *string `json:"end_time"`
	Status       string  `json:"status"`
	Destination1 *string `json:"destination_1"`
	Destination2 *string `json:"destination_2"`
	Destination3 *string `json:"destination_3"`
}

type TripLogBResponse struct {
	TripID       int     `json:"trip_id"`
	VehicleID    string  `json:"vehicle_id"`
	StartTime    *string `json:"start_time"`
	EndTime      *string `json:"end_time"`
	Status       string  `json:"status"`
	Destination1 *string `json:"destination_1"`
	Destination2 *string `json:"destination_2"`
	Destination3 *string `json:"destination_3"`
}
