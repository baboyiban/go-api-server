package models

type Employee struct {
	EmployeeID int    `json:"employee_id" gorm:"column:employee_id;type:int;primaryKey;autoIncrement"`
	Password   string `json:"password" gorm:"column:password;type:varchar(50);not null"`
	Position   string `json:"position" gorm:"column:position;type:enum('관리직','운송직');not null"`
	IsActive   bool   `json:"is_active" gorm:"column:is_active;type:boolean;not null;default:true"`
}

func (Employee) TableName() string {
	return "employee"
}
