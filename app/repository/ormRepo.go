package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	ohost     = "localhost"
	oport     = 5432
	ouser     = "postgres"
	opassword = "user123"
	odbname   = "datb1"
)

var (
	ormdb *gorm.DB
	err   error
)

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	EmpID     int
	Age       int
}

func OpenDB() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ohost, oport, ouser, opassword, odbname)
	//fmt.Println(psqlconn)
	// open database
	var err error
	ormdb, err = gorm.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	ormdb.AutoMigrate(&Employee{})
	fmt.Println("connected")
}

func OrmRead() []Employee {
	OpenDB()
	defer ormdb.Close()
	var emp []Employee
	ormdb.Find(&emp)
	return emp

}

func OrmCreate(emp Employee) {
	OpenDB()
	defer ormdb.Close()

	ormdb.Create(&emp)

}

func OrmUpdate(emp Employee) {
	OpenDB()
	defer ormdb.Close()
	var oldEmp Employee
	ormdb.Where("Emp_Id=?", emp.EmpID).Find(&oldEmp)
	oldEmp.FirstName = emp.FirstName
	oldEmp.LastName = emp.LastName
	oldEmp.Age = emp.Age
	ormdb.Save(&oldEmp)
}

func OrmDelete(id int64) {
	OpenDB()
	defer ormdb.Close()

	var emp Employee
	ormdb.Where("Emp_Id=?", id).Find(&emp)
	ormdb.Delete(&emp)

}
