package models

type Employee struct {
	EmployeeID int    `gorm:"column:직원_ID;type:int;primaryKey;autoIncrement"`
	Name       string `gorm:"column:이름;type:varchar(10)"`
	Password   string `gorm:"column:비밀번호;type:varchar(50)"`
	Position   string `gorm:"column:직책;type:enum('Manager','Transportation')"`
	Active     bool   `gorm:"column:활성여부;type:boolean;default:true"`
}

func (Employee) TableName() string {
	return "직원"
}
