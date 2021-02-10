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

func ReadById(id int64) (models.Employee, error) {
	db := config.OpenDB()
	defer db.Close()

	var emp models.Employee
	err := db.Where("Emp_Id=?", id).Find(&emp).Error
	//	err := db.Where("Emp_Id=?", id).Find(&emp).Pluck("emp_id", emp.EmpID).Pluck("first_name", emp.FirstName).Pluck("last_name", emp.LastName).Pluck("age", emp.Age).Error
	return emp, err
}

func CreateEmployee(emp models.Employee) error {
	db := config.OpenDB()
	defer db.Close()

	err := db.Create(&emp).Error
	return err

}

func UpdateEmployee(emp models.Employee) error {
	db := config.OpenDB()
	defer db.Close()

	var oldEmp models.Employee
	err := db.Where("Emp_Id=?", emp.EmpID).Find(&oldEmp).Error
	if err != nil {
		return err
	}
	oldEmp.FirstName = emp.FirstName
	oldEmp.LastName = emp.LastName
	oldEmp.Age = emp.Age
	err = db.Save(&oldEmp).Error
	return err
}

func DeleteEmployee(id int64) error {
	db := config.OpenDB()
	defer db.Close()

	var emp models.Employee
	err := db.Where("Emp_Id=?", id).Find(&emp).Error
	if err != nil {
		return err
	}
	err = db.Delete(&emp).Error
	return err

}
