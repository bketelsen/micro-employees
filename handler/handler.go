package handler

import (
	"github.com/bketelsen/micro-employees/db"
	emp "github.com/bketelsen/micro-employees/proto/employees"
	"golang.org/x/net/context"
)

type Employee struct{}

func (s *Employee) Get(ctx context.Context, req *emp.GetRequest, rsp *emp.GetResponse) error {
	// Assignment // Get Employee from the db package
	e, err := db.Read(ctx, req.EmpNo)
	if err != nil {
		return err
	} // Add to Response i
	rsp.Employee = e
	return nil
}
