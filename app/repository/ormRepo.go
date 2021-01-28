package repository

import (
	"golang-training/app/models"
	"golang-training/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ReadEmployee() []models.Employee {
	db := config.OpenDB()
	defer db.Close()
	var emp []models.Employee
	db.Find(&emp)
	return emp

}

func CreateEmployee(emp models.Employee) {
	db := config.OpenDB()
	defer db.Close()

	db.Create(&emp)

}

func UpdateEmployee(emp models.Employee) {
	db := config.OpenDB()
	defer db.Close()

	var oldEmp models.Employee
	db.Where("Emp_Id=?", emp.EmpID).Find(&oldEmp)
	oldEmp.FirstName = emp.FirstName
	oldEmp.LastName = emp.LastName
	oldEmp.Age = emp.Age
	db.Save(&oldEmp)
}

func DeleteEmployee(id int64) {
	db := config.OpenDB()
	defer db.Close()

	var emp models.Employee
	db.Where("Emp_Id=?", id).Find(&emp)
	db.Delete(&emp)

}
