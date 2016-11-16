package db

import (
	"context"
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

type ReadResponse struct {
	Employee *emp.Employee
	err      error
}

func Read(ctx context.Context, id int64) (*emp.Employee, error) {
	// Assignment
	// Retrieve Employee from mysql
	// Return it

	ch := make(chan *ReadResponse)
	go func() {
		time.Sleep(2 * time.Millisecond)
		var e Employee
		err := db.Get(&e, "select *  from employees where emp_no = ?", id)
		if err != nil {
			if err == sql.ErrNoRows {
				ch <- &ReadResponse{nil, NoRecord}
			}
			ch <- &ReadResponse{nil, NoRecord}
		}
		ep := &emp.Employee{
			EmpNo:     int64(e.Number),
			FirstName: e.FirstName,
			LastName:  e.LastName,
			HireDate:  e.HireDate.Unix(),
			BirthDate: e.Birthdate.Unix(),
		}

		ch <- &ReadResponse{ep, err}
	}()

	select {
	case rr := <-ch:
		return rr.Employee, rr.err

	// we received the signal of cancelation in this channel
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}
