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

	res := s.empRepo.ReadEmployee()

	//res := ReadEmployee()
	fmt.Println(res)

	var err error = nil
	require.NoError(s.T(), err)

	for _, element := range res {
		require.Nil(s.T(), deep.Equal(models.Employee{EmpID: empID, FirstName: firstName,
			LastName: lastName, Age: age}, element))
	}

}

// func NewMock() (*sql.DB, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	fmt.Println("connected")

// 	return db, mock
// }
//func TestReadEmployee(t *testing.T) {
//db, mock := NewMock()
// query := "SELECT id, first_name, last_name,  FROM employees"
// rows := sqlmock.NewRows([]string{"id", "firstName", "lastName"})
// mock.ExpectQuery(query).WillReturnRows(rows)

// users, err := repo.Find()
// assert.NotEmpty(t, users)
// assert.NoError(t, err)
// assert.Len(t, users, 1)

//}
