package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeID uint   `json:"직원_ID" gorm:"type:int;primaryKey;autoIncrement"`
	Name       string `json:"이름" gorm:"type:varchar(10)"`
	Password   string `json:"비밀번호" gorm:"type:varchar(50)"`
	Position   string `json:"직책" gorm:"type:enum('Manager','Transportation')"`
	Active     bool   `json:"활성여부" gorm:"type:boolean;default:true"`
}
