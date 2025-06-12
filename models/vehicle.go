package models

type Vehicle struct {
	InternalID        int    `json:"internal_id" gorm:"column:internal_id;type:int;primaryKey;autoIncrement"`
	VehicleID         string `json:"vehicle_id" gorm:"column:vehicle_id;type:varchar(15);unique"`
	CurrentLoad       int    `json:"current_load" gorm:"column:current_load;type:int;not null;default:0"`
	MaxLoad           int    `json:"max_load" gorm:"column:max_load;type:int;not null;default:5"`
	LedStatus         string `json:"led_status" gorm:"column:led_status;type:varchar(10)"`
	NeedsConfirmation bool   `json:"needs_confirmation" gorm:"column:needs_confirmation;type:boolean;not null;default:false"`
	CoordX            int    `json:"coord_x" gorm:"column:coord_x;type:int"`
	CoordY            int    `json:"coord_y" gorm:"column:coord_y;type:int"`
}

func (Vehicle) TableName() string {
	return "vehicle"
}
