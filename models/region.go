package models

import (
	"time"
)

type Region struct {
	RegionID        string     `json:"region_id" gorm:"column:region_id;type:char(3);primaryKey"`
	RegionName      string     `json:"region_name" gorm:"column:region_name;type:varchar(50);not null"`
	CoordX          int        `json:"coord_x" gorm:"column:coord_x;type:int"`
	CoordY          int        `json:"coord_y" gorm:"column:coord_y;type:int"`
	MaxCapacity     int        `json:"max_capacity" gorm:"column:max_capacity;type:int;not null;default:0"`
	CurrentCapacity int        `json:"current_capacity" gorm:"column:current_capacity;type:int;not null;default:0"`
	IsFull          bool       `json:"is_full" gorm:"column:is_full;type:boolean;not null;default:false"`
	SaturatedAt     *time.Time `json:"saturated_at" gorm:"column:saturated_at;type:datetime"`
}

func (Region) TableName() string {
	return "region"
}
