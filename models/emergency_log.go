package models

import "time"

type EmergencyLog struct {
	TripID            int       `json:"trip_id" gorm:"column:trip_id;type:int;not null;primaryKey"`
	VehicleID         string    `json:"vehicle_id" gorm:"column:vehicle_id;type:varchar(15);not null;primaryKey"`
	CallTime          time.Time `json:"call_time" gorm:"column:call_time;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	Reason            string    `json:"reason" gorm:"column:reason;type:enum('차량 관련 호출','택배 관련 호출','운송 관련 호출');not null"`
	EmployeeID        int       `json:"employee_id" gorm:"column:employee_id;type:int;not null"`
	NeedsConfirmation bool      `json:"needs_confirmation" gorm:"column:needs_confirmation;type:boolean;not null;default:true"`
}

func (EmergencyLog) TableName() string {
	return "emergency_log"
}
