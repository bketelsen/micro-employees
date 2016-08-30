package db

import (
	"database/sql"
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

func Read(id int64) (*emp.Employee, error) {
	employee := Employee{}
	err := db.QueryRowx("select * from employees where emp_no = ?", id).StructScan(&employee)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return nil, errors.New("not found")
		}
		log.Println(err)
		return nil, err
	}

	e := &emp.Employee{
		EmpNo: int64(employee.Number),
		// ASSIGNMENT = populate the rest
	}

	return e, nil
}
