package models

import (
	"time"
)

type OperationDelivery struct {
	OperationDeliveryID int        `gorm:"column:운행_ID;type:int"`
	DeliveryID          int        `gorm:"column:택배_ID;type:int"`
	Delivery            Delivery   `json:"-"`
	ZoneID              string     `gorm:"column:구역_ID;type:char(3)"`
	Zone                Zone       `json:"-"`
	LoadingOrder        int        `gorm:"column:적재_순번;type:int"`
	RegistrationTime    *time.Time `gorm:"column:등록_시각;type:datetime"`
	ATransportTime      *time.Time `gorm:"column:A차운송_시각;type:datetime"`
	InputTime           *time.Time `gorm:"column:투입_시각;type:datetime"`
	BTransportTime      *time.Time `gorm:"column:B차운송_시각;type:datetime"`
	CompletionTime      *time.Time `gorm:"column:완료_시각;type:datetime"`
}

func (OperationDelivery) TableName() string {
	return "운행_택배"
}
