package dto

type CreateDeliveryLogRequest struct {
	TripID              int     `json:"trip_id" binding:"required"`
	PackageID           int     `json:"package_id" binding:"required"`
	RegionID            string  `json:"region_id" binding:"required"`
	LoadOrder           int     `json:"load_order"`
	RegisteredAt        *string `json:"registered_at"`         // RFC3339
	FirstTransportTime  *string `json:"first_transport_time"`  // RFC3339
	InputTime           *string `json:"input_time"`            // RFC3339
	SecondTransportTime *string `json:"second_transport_time"` // RFC3339
	CompletedAt         *string `json:"completed_at"`          // RFC3339
}

type UpdateDeliveryLogRequest struct {
	LoadOrder           int     `json:"load_order"`
	RegisteredAt        *string `json:"registered_at"`
	FirstTransportTime  *string `json:"first_transport_time"`
	InputTime           *string `json:"input_time"`
	SecondTransportTime *string `json:"second_transport_time"`
	CompletedAt         *string `json:"completed_at"`
}

type DeliveryLogResponse struct {
	TripID              int     `json:"trip_id"`
	PackageID           int     `json:"package_id"`
	RegionID            string  `json:"region_id"`
	LoadOrder           int     `json:"load_order"`
	RegisteredAt        *string `json:"registered_at"`
	FirstTransportTime  *string `json:"first_transport_time"`
	InputTime           *string `json:"input_time"`
	SecondTransportTime *string `json:"second_transport_time"`
	CompletedAt         *string `json:"completed_at"`
}
