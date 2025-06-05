package models

import (
	"time"

	"gorm.io/gorm"
)

type OperationDelivery struct {
	gorm.Model
	OperationDeliveryID uint       `json:"운행_ID" gorm:"type:int"`
	DeliveryID          int        `json:"택배_ID" gorm:"type:int"`
	Delivery            Delivery   ``
	ZoneID              string     `json:"구역_ID" gorm:"type:char(3)"`
	Zone                Zone       ``
	LoadingOrder        int        `json:"적재_순번" gorm:"type:int"`
	RegistrationTime    *time.Time `json:"등록_시각" gorm:"type:datetime"`
	ATransportTime      *time.Time `json:"A차운송_시각" gorm:"type:datetime"`
	InputTime           *time.Time `json:"투입_시각" gorm:"type:datetime"`
	BTransportTime      *time.Time `json:"B차운송_시각" gorm:"type:datetime"`
	CompletionTime      *time.Time `json:"완료_시각" gorm:"type:datetime"`
}
