package handler

import (
	"log"

	"github.com/bketelsen/micro-employees/db"
	"golang.org/x/net/context"

	emp "github.com/bketelsen/micro-employees/proto/employees"
)

type Employee struct{}

func (s *Employee) Get(ctx context.Context, req *emp.GetRequest, rsp *emp.GetResponse) error {
	employee, err := db.Read(req.EmpNo)
	if err != nil {
		log.Println(err)
		return err
	}
	rsp.Employee = employee
	return nil
}
