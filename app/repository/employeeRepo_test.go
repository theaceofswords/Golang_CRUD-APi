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
	var (
		firstName = "test"
		lastName  = "name"
		empID     = 1
		age       = 34
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "employees"`)).
		WillReturnRows(sqlmock.NewRows([]string{"emp_id", "first_name", "last_name", "age"}).
			AddRow(empID, firstName, lastName, age))

	res, err := s.empRepo.ReadEmployee()

	fmt.Println(res)

	require.NoError(s.T(), err)

	for _, element := range res {
		require.Nil(s.T(), deep.Equal(models.Employee{EmpID: empID, FirstName: firstName,
			LastName: lastName, Age: age}, element))
	}

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
		empID     = 2
		age       = 24
	)

	emp := models.Employee{
		FirstName: firstName,
		LastName:  lastName,
		EmpID:     empID,
		Age:       age,
	}

	fmt.Println(emp)
	//now := time.Now()
	// s.mock.ExpectBegin()
	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`INSERT INTO "employees" ("created_at","updated_at","deleted_at","first_name","last_name","emp_id","age") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "employees"."id"`)).
	// 	WithArgs(now, now, nil, firstName, lastName, empID, age).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"emp_id"}).AddRow(empID))
	s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO employees VALUES ($1,$2,$3,$4)`)).
		WithArgs(firstName, lastName, empID, age)

	// s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO employees VALUES ($1,$2,$3,$4) RETURNING employees.emp_id`)).
	// 	WithArgs(firstName, lastName, empID, age).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"emp_id"}).AddRow(empID))

	// s.mock.ExpectQuery(`INSERT INTO employees VALUES ($1,$2,$3,$4), emp.FirstName, emp.LastName, emp.EmpID, emp.Age)`)

	pkID, err := s.empRepo.CreateEmployee(emp)
	//s.mock.ExpectCommit()
	fmt.Println(pkID)

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

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "employees" WHERE "employees"."deleted_at" IS NULL AND ((emp_id=$1))`)).
		WithArgs(empID).
		WillReturnRows(sqlmock.NewRows([]string{"emp_id", "first_name", "last_name", "age"}).
			AddRow(empID, firstName, lastName, age))
	//s.mock.ExpectBegin()
	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`INSERT INTO "employees" ("created_at","updated_at","deleted_at","first_name","last_name","emp_id","age") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "employees"."id"`)).
	// 	WithArgs(now, now, nil, firstName, lastName, empID, age).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"emp_id"}).AddRow(empID))
	s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO employees VALUES ($1,$2,$3,$4)`)).
		WithArgs(firstName, lastName, empID, age)
	//s.mock.ExpectCommit()

	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`UPDATE "employees" SET ((first_name=$2)), ((last_name=$3)),((age=$4)) WHERE ((emp_id=$1))`)).
	// 	WithArgs(empID, firstName, lastName, age).
	// 	WillReturnRows(sqlmock.NewRows([]string{"emp_id", "first_name", "last_name", "age"}).
	// 		AddRow(empID, firstName, lastName, age))

	err := s.empRepo.UpdateEmployee(emp)

	require.NoError(s.T(), err)

}
