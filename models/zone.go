package models

import (
	"time"
)

type Zone struct {
	ZoneID                 string     `gorm:"column:구역_ID;type:char(3);primaryKey"`
	ZoneName               string     `gorm:"column:구역_명;type:varchar(50)"`
	CoordinateX            int        `gorm:"column:좌표_X;type:int"`
	CoordinateY            int        `gorm:"column:좌표_Y;type:int"`
	MaxStorageQuantity     int        `gorm:"column:최대_보관_수량;type:int"`
	CurrentStorageQuantity int        `gorm:"column:현재_보관_수량;type:int;default:0"`
	SaturationStatus       bool       `gorm:"column:포화_여부;type:boolean"`
	SaturationTime         *time.Time `gorm:"column:포화_시각;type:datetime"`
}

func (Zone) TableName() string {
	return "구역"
}
