package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	InternalID          int    `json:"내부_ID" gorm:"type:int;primaryKey;autoIncrement"`
	VehicleID           string `json:"차량_ID" gorm:"type:varchar(15);unique"`
	CurrentLoadQuantity int    `json:"현재_적재_수량" gorm:"type:int;default:0"`
	MaxLoadQuantity     int    `json:"최대_적재_수량" gorm:"type:int;default:5"`
	LEDStatus           string `json:"LED_상태" gorm:"type:enum('초록','노랑','빨강');default:'초록'"`
	ConfirmRequired     bool   `json:"담당_확인_필요" gorm:"type:boolean;defeault:false"`
	CurrentCoordinateX  int    `json:"현재_좌표_X" gorm:"type:int"`
	CurrentCoordinateY  int    `json:"현재_좌표_Y" gorm:"type:int"`
}
