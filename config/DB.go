package config

import (
	"context"
	"database/sql"
	"fmt"
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
	dbname   = "test1"
)

var (
	ctx context.Context

	db *sql.DB
)

// DBConn Exported
func DBConn() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	insertStatement := `INSERT INTO tbluser VALUES ('name6',18,43)`
	_, err = db.Query(insertStatement)
	if err != nil {
		fmt.Println(err)
	}

	selectStatement := `SELECT name, age FROM tbluser WHERE id=$1;`

	var name string
	var age int
	row := db.QueryRow(selectStatement, 17)
	err = row.Scan(&name, &age)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(name, age)
	default:
		panic(err)
	}

	updateStatement := `UPDATE tbluser SET name =$1 WHERE id =$2`

	_, err = db.Exec(updateStatement, "Name", 24)
	if err != nil {
		panic(err)
	}

	deleteStatement := `DELETE FROM tbluser WHERE id = $1;`
	result, err := db.Exec(deleteStatement, 17)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("%d rows deleted ", count)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Closed")

}
