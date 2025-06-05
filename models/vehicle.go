package models

type Vehicle struct {
	InternalID          int    `gorm:"column:내부_ID;type:int;primaryKey;autoIncrement"`
	VehicleID           string `gorm:"column:차량_ID;type:varchar(15);unique"`
	CurrentLoadQuantity int    `gorm:"column:현재_적재_수량;type:int;default:0"`
	MaxLoadQuantity     int    `gorm:"column:최대_적재_수량;type:int;default:5"`
	LEDStatus           string `gorm:"column:LED_상태;type:enum('초록','노랑','빨강');default:'초록'"`
	ConfirmRequired     bool   `gorm:"column:담당_확인_필요;type:boolean;default:false"`
	CurrentCoordinateX  int    `gorm:"column:현재_좌표_X;type:int"`
	CurrentCoordinateY  int    `gorm:"column:현재_좌표_Y;type:int"`
}

func (Vehicle) TableName() string {
	return "차량"
}
