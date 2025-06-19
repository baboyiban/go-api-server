package models

import "time"

type TripLogA struct {
	TripID       int        `json:"trip_id" gorm:"column:trip_id;type:int;primaryKey;autoIncrement"`
	VehicleID    string     `json:"vehicle_id" gorm:"column:vehicle_id;type:varchar(15);not null"`
	StartTime    *time.Time `json:"start_time" gorm:"column:start_time;type:datetime"`
	EndTime      *time.Time `json:"end_time" gorm:"column:end_time;type:datetime"`
	Status       string     `json:"status" gorm:"column:status;type:enum('운행중','비운행중');not null;default:'비운행중'"`
	Destination1 *string    `json:"destination_1" gorm:"column:destination_1;type:char(3)"`
	Destination2 *string    `json:"destination_2" gorm:"column:destination_2;type:char(3)"`
	Destination3 *string    `json:"destination_3" gorm:"column:destination_3;type:char(3)"`
}

func (TripLogA) TableName() string {
	return "trip_log_A"
}
