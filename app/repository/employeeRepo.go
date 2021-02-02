package repository

import (
	"golang-training/app/models"

	"github.com/jinzhu/gorm"
)

type CRUD interface {
	ReadEmployee() []models.Employee
	CreateEmployee(emp models.Employee)
	UpdateEmployee(emp models.Employee)
	DeleteEmployee(id int64)
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) ReadEmployee() []models.Employee {
	var emp []models.Employee
	p.DB.Find(&emp)
	return emp
}

func (p *repo) CreateEmployee(emp models.Employee) {}
func (p *repo) UpdateEmployee(emp models.Employee) {}
func (p *repo) DeleteEmployee(id int64)            {}

func CreateRepository(db *gorm.DB) CRUD {
	return &repo{
		DB: db,
	}
}
