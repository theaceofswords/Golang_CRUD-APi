package repository

import (
	"database/sql"
	"fmt"
	"golang-training/app/models"
	"regexp"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	empRepo  CRUD
	employee *models.Employee
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.empRepo = CreateRepository(s.DB)

}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestReadEmployee() {

	empList := []models.Employee{models.Employee{
		FirstName: "test",
		LastName:  "name",
		EmpID:     1,
		Age:       34},

		models.Employee{
			FirstName: "test2",
			LastName:  "name2",
			EmpID:     2,
			Age:       65}}
	//fmt.Println(empList)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "employees"`)).
		WillReturnRows(sqlmock.NewRows([]string{"emp_id", "first_name", "last_name", "age"}).
			AddRow(1, "test", "name", 34).
			AddRow(2, "test2", "name2", 65))

	res, err := s.empRepo.ReadEmployee()

	//fmt.Println(res)

	require.NoError(s.T(), err)

	// for _, element := range res {
	// 	require.Nil(s.T(), deep.Equal(models.Employee{EmpID: empID, FirstName: firstName,
	// 		LastName: lastName, Age: age}, element))
	// }
	require.Nil(s.T(), deep.Equal(empList, res))

}

func (s *Suite) TestFindByID() {
	var (
		firstName = "test"
		lastName  = "name"
		empID     = 1
		age       = 34
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "employees" WHERE "employees"."deleted_at" IS NULL AND ((emp_id=$1))`)).
		WithArgs(empID).
		WillReturnRows(sqlmock.NewRows([]string{"emp_id", "first_name", "last_name", "age"}).
			AddRow(empID, firstName, lastName, age))

	res, err := s.empRepo.FindByID(empID)

	fmt.Println(res)

	require.NoError(s.T(), err)

	require.Nil(s.T(), deep.Equal(models.Employee{EmpID: empID, FirstName: firstName,
		LastName: lastName, Age: age}, res))

}

func (s *Suite) TestUCreateEmployee() {
	var (
		firstName = "test2"
		lastName  = "name2"
		empID     = 2021
		age       = 24
	)

	emp := models.Employee{
		FirstName: firstName,
		LastName:  lastName,
		EmpID:     empID,
		Age:       age,
	}

	//fmt.Println(emp)
	//now := time.Now()
	// s.mock.ExpectBegin()
	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`INSERT INTO "employees" ("created_at","updated_at","deleted_at","first_name","last_name","emp_id","age") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "employees"."id"`)).
	// 	WithArgs(now, now, nil, firstName, lastName, empID, age).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"emp_id"}).AddRow(empID))
	s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO employees VALUES ($1,$2,$3,$4) RETURNING emp_id`)).
		WithArgs(firstName, lastName, empID, age).
		WillReturnResult(sqlmock.NewResult(int64(empID), 1))

	// s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO employees VALUES ($1,$2,$3,$4) RETURNING employees.emp_id`)).
	// 	WithArgs(firstName, lastName, empID, age).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"emp_id"}).AddRow(empID))
	//s.mock.ExpectCommit()
	// s.mock.ExpectQuery(`INSERT INTO employees VALUES ($1,$2,$3,$4), emp.FirstName, emp.LastName, emp.EmpID, emp.Age)`)

	pkID, err := s.empRepo.CreateEmployee(emp)
	require.Equal(s.T(), pkID, empID, "Id returned was not equal")
	require.NoError(s.T(), err)
}

func (s *Suite) TestZUpdateEmployee() {
	var (
		firstName = "tester"
		lastName  = "namer"
		empID     = 1
		age       = 34
	)
	emp := models.Employee{
		FirstName: firstName,
		LastName:  lastName,
		EmpID:     empID,
		Age:       age,
	}
	//now := time.Now()

	//s.mock.ExpectBegin()
	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`INSERT INTO "employees" ("created_at","updated_at","deleted_at","first_name","last_name","emp_id","age") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "employees"."id"`)).
	// 	WithArgs(now, now, nil, firstName, lastName, empID, age).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"emp_id"}).AddRow(empID))
	//================
	// s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO employees VALUES ($1,$2,$3,$4) RETURNING emp_id`)).
	// 	WithArgs(firstName, lastName, empID, age).
	// 	WillReturnResult(sqlmock.NewResult(1, 1))
	//=====================
	//s.mock.ExpectCommit()

	s.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE employees SET first_name = $2, last_name = $3, age = $4 WHERE emp_id = $1`)).
		WithArgs(empID, firstName, lastName, age).
		WillReturnResult(sqlmock.NewResult(int64(empID), 1))

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "employees" WHERE "employees"."deleted_at" IS NULL AND ((emp_id=$1))`)).
		WithArgs(empID).
		WillReturnRows(sqlmock.NewRows([]string{"emp_id", "first_name", "last_name", "age"}).
			AddRow(empID, firstName, lastName, age))

	res, err := s.empRepo.UpdateEmployee(emp)
	//fmt.Println(res)

	require.Nil(s.T(), deep.Equal(emp, res))

	require.NoError(s.T(), err)

}

func (s *Suite) TestZZDeleteEmployee() {
	var (
		firstName = "tester"
		lastName  = "namer"
		empID     = 1
		age       = 34
	)
	// emp := models.Employee{
	// 	FirstName: firstName,
	// 	LastName:  lastName,
	// 	EmpID:     empID,
	// 	Age:       age,
	// }
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "employees" WHERE "employees"."deleted_at" IS NULL AND ((Emp_Id=$1))`)).
		WithArgs(empID).
		WillReturnRows(sqlmock.NewRows([]string{"emp_id", "first_name", "last_name", "age"}).
			AddRow(empID, firstName, lastName, age))

	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`UPDATE "employees" SET "deleted_at"=$1  WHERE "employees"."deleted_at" IS NULL`)).
	// 	WithArgs(time.Now())

	s.mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM tbl_employee WHERE id = $1`)).
		WithArgs(empID).
		WillReturnResult(sqlmock.NewResult(int64(empID), 1))
	err := s.empRepo.DeleteEmployee(empID)

	require.NoError(s.T(), err)
}
