package models

import (
	"time"
)

type Package struct {
	PackageID     int       `json:"package_id" gorm:"column:package_id;type:int;primaryKey;autoIncrement"`
	PackageType   string    `json:"package_type" gorm:"column:package_type;type:varchar(50);not null"`
	RegionID      string    `json:"region_id" gorm:"column:region_id;type:char(3);not null"`
	Region        Region    `json:"-" gorm:"foreignKey:RegionID;references:RegionID"`
	PackageStatus string    `json:"package_status" gorm:"column:package_status;type:enum('registered','A_transport','input','B_transport','completed')"`
	RegisteredAt  time.Time `json:"registered_at" gorm:"column:registered_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

func (Package) TableName() string {
	return "package"
}
