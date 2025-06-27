package dto

import (
	"github.com/baboyiban/go-api-server/models"
	"github.com/baboyiban/go-api-server/utils"
)

type CreateRegionRequest struct {
	RegionID    string `json:"region_id" binding:"required,len=3"`
	RegionName  string `json:"region_name" binding:"required"`
	CoordX      int    `json:"coord_x"`
	CoordY      int    `json:"coord_y"`
	MaxCapacity int    `json:"max_capacity"`
}

type UpdateRegionRequest struct {
	RegionName      string  `json:"region_name" binding:"required"`
	CoordX          int     `json:"coord_x"`
	CoordY          int     `json:"coord_y"`
	MaxCapacity     int     `json:"max_capacity"`
	CurrentCapacity int     `json:"current_capacity"`
	IsFull          bool    `json:"is_full"`
	SaturatedAt     *string `json:"saturated_at,omitempty"`
}

type RegionResponse struct {
	RegionID        string  `json:"region_id"`
	RegionName      string  `json:"region_name"`
	CoordX          int     `json:"coord_x"`
	CoordY          int     `json:"coord_y"`
	MaxCapacity     int     `json:"max_capacity"`
	CurrentCapacity int     `json:"current_capacity"`
	IsFull          bool    `json:"is_full"`
	SaturatedAt     *string `json:"saturated_at,omitempty"`
}

func ToRegionResponse(region *models.Region) RegionResponse {
	return RegionResponse{
		RegionID:        region.RegionID,
		RegionName:      region.RegionName,
		CoordX:          region.CoordX,
		CoordY:          region.CoordY,
		MaxCapacity:     region.MaxCapacity,
		CurrentCapacity: region.CurrentCapacity,
		IsFull:          region.IsFull,
		SaturatedAt:     utils.FormatTimePtr(region.SaturatedAt),
	}
}
