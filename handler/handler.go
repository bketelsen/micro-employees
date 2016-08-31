package handler

import (
	"golang.org/x/net/context"

	emp "github.com/bketelsen/micro-employees/proto/employees"
)

type Employee struct{}

func (s *Employee) Get(ctx context.Context, req *emp.GetRequest, rsp *emp.GetResponse) error {
	// Assignment
	// Get Employee from the db package
	// Add to Response
}
