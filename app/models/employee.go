package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	EmpID     int
	Age       int
}
