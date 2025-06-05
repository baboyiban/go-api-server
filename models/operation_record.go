package models

type OperationRecord struct {
	OperationID     int     `gorm:"column:운행_ID;type:int;primaryKey;autoInrement"`
	VehicleID       int     `gorm:"column:차량_ID;type:varchar(15);not null"`
	Vehicle         Vehicle `json:"-"`
	OperationStart  int     `gorm:"column:운행_시작;type:datetime"`
	EndTime         int     `gorm:"column:종료_시각;type:datetime"`
	OperationStatus string  `gorm:"column:운행_상태;type:enum('운행중','비운행중');default:'비운행중'"`
}

func (OperationRecord) TableName() string {
	return "운행_기록"
}
