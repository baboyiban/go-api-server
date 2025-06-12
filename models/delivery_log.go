package models

import (
	"time"
)

type DeliveryLog struct {
	DeliveryID          int        `json:"delivery_id" gorm:"column:delivery_id;type:int;primaryKey;autoIncrement"`
	TripID              int        `json:"trip_id" gorm:"column:trip_id;type:int;not null;"`
	TripLog             TripLog    `json:"-" gorm:"foreignKey:TripID;references:TripID"`
	PackageID           int        `json:"package_id" gorm:"column:package_id;type:int;not null"`
	Package             Package    `json:"-" gorm:"foreignKey:PackageID;references:PackageID"`
	RegionID            string     `json:"region_id" gorm:"column:region_id;type:char(3)"`
	Region              Region     `json:"-" gorm:"foreignKey:RegionID;references:RegionID"`
	LoadOrder           int        `json:"load_order" gorm:"column:load_order;type:int"`
	RegistrationTime    time.Time  `json:"registration_time" gorm:"column:registration_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	FirstTransportTime  *time.Time `json:"first_transport_time" gorm:"column:first_transport_time;type:datetime"`
	InputTime           *time.Time `json:"input_time" gorm:"column:input_time;type:datetime"`
	SecondTransportTime *time.Time `json:"second_transport_time" gorm:"column:second_transport_time;type:datetime"`
	CompletionAt        *time.Time `json:"completion_at" gorm:"column:completion_at;type:datetime"`
}

func (DeliveryLog) TableName() string {
	return "delivery_log"
}
