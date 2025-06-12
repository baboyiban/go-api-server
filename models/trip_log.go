package models

import "time"

type TripLog struct {
	TripID    int        `json:"trip_id" gorm:"column:trip_id;type:int;primaryKey;autoIncrement"`
	VehicleID string     `json:"vehicle_id" gorm:"column:vehicle_id;type:varchar(15);not null"`
	Vehicle   Vehicle    `json:"-" gorm:"foreignKey:VehicleID;references:VehicleID"`
	StartTime *time.Time `json:"start_time" gorm:"column:start_time;type:datetime"`
	EndTime   *time.Time `json:"end_time" gorm:"column:end_time;type:datetime"`
	Status    string     `json:"status" gorm:"column:status;type:enum('운송중','비운송중');not null;default:'비운송중'"`
}

func (TripLog) TableName() string {
	return "trip_log"
}
