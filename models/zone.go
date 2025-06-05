package models

import "time"

type Zone struct {
	ZoneID          string `gorm:"primaryKey;type:char(3)"`
	ZoneName        string `gorm:"type:varchar(50)"`
	CoordinateX     int
	CoordinateY     int
	MaxCapacity     int
	CurrentCapacity int  `gorm:"default:0"`
	IsFull          bool `gorm:"default:false"`
	FullAt          *time.Time
}
