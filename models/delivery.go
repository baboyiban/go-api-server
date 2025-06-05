package models

import (
	"time"
)

type Delivery struct {
	delivery_id       uint       `json:"택배_ID" gorm:"type:int"`
	delivery_type     string     `json:"택배_종류" gorm:"type:varchar(50)"`
	zone_id           string     `json:"구역_ID" gorm:"type:char(3)"`
	current_status    string     `json:"현재_상태" gorm:"type:enum('미등록','등록됨','A차운송중','투입됨','B차운송중');default:'등록됨'"`
	registration_time *time.Time `json:"등록_시각" gorm:"type:datetime"`
}
