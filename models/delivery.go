package models

import (
	"time"
)

type Delivery struct {
	DeliveryID       uint       `json:"택배_ID" gorm:"column:택배_ID;type:int;primaryKey;autoIncrement"`
	DeliveryType     string     `json:"택배_종류" gorm:"column:택배_종류;type:varchar(50)"`
	ZoneID           string     `json:"구역_ID" gorm:"column:구역_ID;type:char(3)"`
	Zone             Zone       ``
	CurrentStatus    string     `json:"현재_상태" gorm:"column:현재_상태;type:enum('미등록','등록됨','A차운송중','투입됨','B차운송중');default:'등록됨'"`
	RegistrationTime *time.Time `json:"등록_시각" gorm:"column:등록_시각;type:datetime"`
}

func (Delivery) TableName() string {
	return "택배"
}
