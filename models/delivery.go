package models

import (
	"time"
)

type Delivery struct {
	DeliveryID       int        `gorm:"column:택배_ID;type:int;primaryKey;autoIncrement"`
	ZoneID           string     `gorm:"column:구역_ID;type:char(3)"`
	Zone             Zone       `json:"-"`
	DeliveryType     string     `gorm:"column:택배_종류;type:varchar(50)"`
	CurrentStatus    string     `gorm:"column:현재_상태;type:enum('미등록','등록됨','A차운송중','투입됨','B차운송중');default:'등록됨'"`
	RegistrationTime *time.Time `gorm:"column:등록_시각;type:datetime"`
}

func (Delivery) TableName() string {
	return "택배"
}
