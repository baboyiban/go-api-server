package models

import (
	"time"
)

type Package struct {
	PackageID     int       `json:"package_id" gorm:"column:package_id;type:int;primaryKey;autoIncrement"`
	PackageType   string    `json:"package_type" gorm:"column:package_type;type:varchar(50);not null;uniqueIndex:unique_package_info"`
	RegionID      string    `json:"region_id" gorm:"column:region_id;type:char(3);not null;uniqueIndex:unique_package_info"`
	Region        Region    `json:"-" gorm:"foreignKey:RegionID;references:RegionID"`
	PackageStatus string    `json:"package_status" gorm:"column:package_status;type:enum('등록됨','A차운송중','투입됨','B차운송중','완료됨';default:'등록됨'"`
	RegisteredAt  time.Time `json:"registered_at" gorm:"column:registered_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

func (Package) TableName() string {
	return "package"
}
