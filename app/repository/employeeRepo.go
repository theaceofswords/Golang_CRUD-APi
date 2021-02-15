package repository

import (
	"fmt"
	"golang-training/app/models"

	"github.com/jinzhu/gorm"
)

type CRUD interface {
	ReadEmployee() ([]models.Employee, error)
	FindByID(id int) (models.Employee, error)
	CreateEmployee(emp models.Employee) (int, error)
	UpdateEmployee(emp models.Employee) (models.Employee, error)
	DeleteEmployee(id int) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) ReadEmployee() ([]models.Employee, error) {
	var emp []models.Employee
	err := p.DB.Find(&emp).Debug().Error
	return emp, err
}

func (p *repo) FindByID(id int) (models.Employee, error) {
	var emp models.Employee
	err := p.DB.Where("emp_id=?", id).Find(&emp).Error
	return emp, err
}

func (p *repo) CreateEmployee(emp models.Employee) (int, error) {
	//result := p.DB.Create(&emp)

	insertStatement := `INSERT INTO employees VALUES ($1,$2,$3,$4) RETURNING emp_id`

	result := p.DB.Exec(insertStatement, emp.FirstName, emp.LastName, emp.EmpID, emp.Age)

	fmt.Println("=====", result)
	//id := emp.ID
	//err := result.Error

	return emp.EmpID, nil
	//return emp.EmpID, err
}

func (p *repo) UpdateEmployee(emp models.Employee) (models.Employee, error) {

	updateStatement := `UPDATE employees SET first_name = $2, last_name = $3, age = $4 WHERE emp_id = $1`
	result := p.DB.Exec(updateStatement, emp.EmpID, emp.FirstName, emp.LastName, emp.Age)
	fmt.Printf("%T", result)

	var newEmp models.Employee
	err := p.DB.Where("emp_id=?", emp.EmpID).Find(&newEmp).Error

	return newEmp, err
}

func (p *repo) DeleteEmployee(id int) error {
	var emp models.Employee
	p.DB.Where("Emp_Id=?", id).Find(&emp)
	//err := p.DB.Delete(&emp).Error
	deleteStatement := `DELETE FROM tbl_employee WHERE id = $1`
	err := p.DB.Exec(deleteStatement, id).Error
	//count, _ := result.RowsAffected()
	//fmt.Println(result)

	return err
}

func CreateRepository(db *gorm.DB) CRUD {
	return &repo{
		DB: db,
	}
}
