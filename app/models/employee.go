package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	FirstName string `gorm:"column:first_name" `
	LastName  string `gorm:"column:last_name`
	EmpID     int    `gorm:"column:emp_id;primary_key"`
	Age       int    `gorm:"column:age`
}

type GetParam struct {
	EmpId int
}
