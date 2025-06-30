package models

import "time"

type ChangeNotify struct {
	ID        int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name      string    `json:"table_name" gorm:"column:table_name;type:varchar(64);not null"`
	Action    string    `json:"action" gorm:"column:action;type:varchar(16);not null"`
	ChangedAt time.Time `json:"changed_at" gorm:"column:changed_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

func (ChangeNotify) TableName() string {
	return "change_notify"
}
