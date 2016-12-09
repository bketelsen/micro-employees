package handler

import (
	emp "github.com/bketelsen/micro-employees/proto/employees"
	"golang.org/x/net/context"
)

type Employee struct{}

func (s *Employee) Get(ctx context.Context, req *emp.GetRequest, rsp *emp.GetResponse) error {
	// Assignment
	// Get Employee from the db package
	// Add it to the response
	return nil

}
