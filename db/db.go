package db

import (
	"context"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	emp "github.com/bketelsen/micro-employees/proto/employees"
)

var (
	Url      = "docker:docker@tcp(127.0.0.1:3306)/employees?parseTime=true"
	database string
	db       *sqlx.DB
)

var NoRecord = errors.New("0x4f89: NO RECORD")

func Init() {
	var d *sqlx.DB
	var err error

	if d, err = sqlx.Open("mysql", Url); err != nil {
		log.Fatal(err)
	}

	db = d

}

// Employee represents the employee model in the database
// 'db' struct tags tell sqlx how to map data
type Employee struct {
	Number    int       `db:"emp_no"`
	Birthdate time.Time `db:"birth_date"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Gender    string    `db:"gender"`
	HireDate  time.Time `db:"hire_date"`
}

func Read(ctx context.Context, id int64) (*emp.Employee, error) {
	// Assignment
	// Retrieve Employee from mysql
	// Return it

}
