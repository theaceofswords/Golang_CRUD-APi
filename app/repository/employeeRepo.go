package repository

import (
	"golang-training/app/models"

	"github.com/jinzhu/gorm"
)

type CRUD interface {
	ReadEmployee() ([]models.Employee, error)
	FindByID(id int) (models.Employee, error)
	CreateEmployee(emp models.Employee) (int, error)
	UpdateEmployee(emp models.Employee) error
	DeleteEmployee(id int64)
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

	insertStatement := `INSERT INTO employees VALUES ($1,$2,$3,$4)`
	//p.DB.Exec(insertStatement)
	p.DB.Exec(insertStatement, emp.FirstName, emp.LastName, emp.EmpID, emp.Age)
	//p.DB.Exec(`INSERT INTO "employees" ("first_name","last_name","emp_id","age") VALUES ($1,$2,$3,$4)`,emp.FirstName, emp.LastName, emp.EmpID, emp.Age)

	//id := emp.ID
	//err := result.Error

	return 2, nil
	//return emp.EmpID, err
}

func (p *repo) UpdateEmployee(emp models.Employee) error {
	var oldEmp models.Employee
	err := p.DB.Where("emp_id=?", emp.EmpID).Find(&oldEmp).Error
	if err != nil {
		return err
	}
	oldEmp.FirstName = emp.FirstName
	oldEmp.LastName = emp.LastName
	oldEmp.Age = emp.Age
	//err = p.DB.Save(&oldEmp).Error
	insertStatement := `INSERT INTO employees VALUES ($1,$2,$3,$4)`
	p.DB.Exec(insertStatement, oldEmp.FirstName, oldEmp.LastName, oldEmp.EmpID, oldEmp.Age)
	return err
}

func (p *repo) DeleteEmployee(id int64) {}

func CreateRepository(db *gorm.DB) CRUD {
	return &repo{
		DB: db,
	}
}
