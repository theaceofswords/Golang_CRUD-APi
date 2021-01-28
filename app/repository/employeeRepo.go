package repository

import (
	"database/sql"
	"fmt"
	"golang-training/app/models"
	"log"

	/*Not using the package identifier in the program,
	so the compiler should ignore the error of not
	using the package identifier but will still invoke the init function.*/
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "user123"
	dbname   = "datb1"
)

var (
	db *sql.DB
	//emp models.[]Employee
)

func ConnectDB() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//fmt.Println(psqlconn)
	// open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
}

func GetAllEmployee() []models.Employee {

	readStatement := `SELECT * FROM tbl_employee;`
	rows, err := db.Query(readStatement)
	checkError(err)
	var emp models.Employee
	empList := []models.Employee{}

	for rows.Next() {
		err := rows.Scan(&emp.FirstName, &emp.LastName, &emp.ID, &emp.Age)
		checkError(err)
		empList = append(empList, emp)

	}
	return empList

}

func AddEmployee(emp models.Employee) []models.Employee {
	insertStatement := `INSERT INTO tbl_employee VALUES($1,$2,$3,$4)`
	_, err := db.Query(insertStatement, emp.FirstName, emp.LastName, emp.ID, emp.Age)
	checkError(err)

	return GetAllEmployee()
}
func UpdateEmployee(emp models.Employee) []models.Employee {
	updateStatement := `UPDATE tbl_employee SET first_name=$1, last_name=$2,age=$3 WHERE id=$4;`
	_, err := db.Exec(updateStatement, emp.FirstName, emp.LastName, emp.Age, emp.ID)
	checkError(err)

	return GetAllEmployee()
}

func RemoveEmployee(id int64) {
	deleteStatement := `DELETE FROM tbl_employee WHERE id = $1`
	result, err := db.Exec(deleteStatement, id)
	checkError(err)
	count, _ := result.RowsAffected()
	fmt.Println("rows deleted: ", count)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
