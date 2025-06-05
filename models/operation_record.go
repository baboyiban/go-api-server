package models

import "gorm.io/gorm"

type OperationRecord struct {
	gorm.Model
	OperationID     uint    `json:"운행_ID" gorm:"type:int;primaryKey;autoInrement"`
	VehicleID       int     `json:"차량_ID" gorm:"type:varchar(15);not null"`
	Vehicle         Vehicle ``
	OperationStart  int     `json:"운행_시작" gorm:"type:datetime"`
	EndTime         int     `json:"종료_시각" gorm:"type:datetime"`
	OperationStatus string  `json:"운행_상태" gorm:"type:enum('운행중','비운행중');default:'비운행중'"`
}
