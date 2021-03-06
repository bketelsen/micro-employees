// Code generated by protoc-gen-go.
// source: github.com/bketelsen/micro-employees/proto/employees.proto
// DO NOT EDIT!

/*
Package employees is a generated protocol buffer package.

It is generated from these files:
	github.com/bketelsen/micro-employees/proto/employees.proto

It has these top-level messages:
	Employee
	GetRequest
	GetResponse
*/
package employees

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Employee_Gender int32

const (
	Employee_MALE   Employee_Gender = 0
	Employee_FEMALE Employee_Gender = 1
)

var Employee_Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}
var Employee_Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Employee_Gender) String() string {
	return proto.EnumName(Employee_Gender_name, int32(x))
}
func (Employee_Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Employee struct {
	EmpNo     int64  `protobuf:"varint,1,opt,name=emp_no,json=empNo" json:"emp_no,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	BirthDate int64  `protobuf:"varint,4,opt,name=birth_date,json=birthDate" json:"birth_date,omitempty"`
	HireDate  int64  `protobuf:"varint,5,opt,name=hire_date,json=hireDate" json:"hire_date,omitempty"`
}

func (m *Employee) Reset()                    { *m = Employee{} }
func (m *Employee) String() string            { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()               {}
func (*Employee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetRequest struct {
	EmpNo int64 `protobuf:"varint,1,opt,name=emp_no,json=empNo" json:"emp_no,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetResponse struct {
	Employee *Employee `protobuf:"bytes,1,opt,name=employee" json:"employee,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

func init() {
	proto.RegisterType((*Employee)(nil), "Employee")
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
	proto.RegisterEnum("Employee_Gender", Employee_Gender_name, Employee_Gender_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Employees service

type EmployeesClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type employeesClient struct {
	c           client.Client
	serviceName string
}

func NewEmployeesClient(serviceName string, c client.Client) EmployeesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "employees"
	}
	return &employeesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *employeesClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Employees.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Employees service

type EmployeesHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
}

func RegisterEmployeesHandler(s server.Server, hdlr EmployeesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Employees{hdlr}, opts...))
}

type Employees struct {
	EmployeesHandler
}

func (h *Employees) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.EmployeesHandler.Get(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/bketelsen/micro-employees/proto/employees.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x8b, 0x6d, 0x09, 0x3b, 0x78, 0x30, 0x9b, 0x98, 0x10, 0x4d, 0x4d, 0xb3, 0xc6, 0xa4,
	0x17, 0x21, 0xa9, 0x9e, 0xbc, 0x99, 0x88, 0x5c, 0xb4, 0x07, 0x5e, 0x80, 0x40, 0x3b, 0x0a, 0x91,
	0x65, 0x71, 0x77, 0x7a, 0xf0, 0xd1, 0x7c, 0x3b, 0xc3, 0x28, 0xd5, 0x8b, 0xb7, 0x9d, 0xef, 0x9b,
	0x9d, 0xdd, 0xf9, 0xe1, 0xee, 0xb5, 0xa1, 0x7a, 0x5f, 0xc5, 0x5b, 0xa3, 0x93, 0xea, 0x0d, 0x09,
	0x5b, 0x87, 0x5d, 0xa2, 0x9b, 0xad, 0x35, 0xd7, 0xa8, 0xfb, 0xd6, 0x7c, 0x20, 0xba, 0xa4, 0xb7,
	0x86, 0x4c, 0x72, 0xa8, 0x63, 0xae, 0xd5, 0xa7, 0x07, 0x41, 0xfa, 0xc3, 0xe4, 0x29, 0xf8, 0xa8,
	0xfb, 0xa2, 0x33, 0x91, 0xb7, 0xf4, 0x56, 0xd3, 0x7c, 0x8e, 0xba, 0xdf, 0x18, 0xb9, 0x00, 0x78,
	0x69, 0xac, 0xa3, 0xa2, 0x2b, 0x35, 0x46, 0x47, 0x4b, 0x6f, 0x25, 0x72, 0xc1, 0x64, 0x53, 0x6a,
	0x94, 0xe7, 0x20, 0xda, 0x72, 0xb4, 0x53, 0xb6, 0xc1, 0x00, 0x58, 0x2e, 0x00, 0xaa, 0xc6, 0x52,
	0x5d, 0xec, 0x4a, 0xc2, 0x68, 0xc6, 0x63, 0x05, 0x93, 0x87, 0x92, 0xf8, 0x6e, 0xdd, 0x58, 0xfc,
	0xb6, 0x73, 0xb6, 0xc1, 0x00, 0x06, 0xa9, 0x2e, 0xc0, 0xcf, 0xb0, 0xdb, 0xa1, 0x95, 0x01, 0xcc,
	0x9e, 0xef, 0x9f, 0xd2, 0x93, 0x89, 0x04, 0xf0, 0x1f, 0x53, 0x3e, 0x7b, 0xea, 0x12, 0x20, 0x43,
	0xca, 0xf1, 0x7d, 0x8f, 0x8e, 0xfe, 0xf9, 0xbc, 0xba, 0x85, 0x90, 0x9b, 0x5c, 0x6f, 0x3a, 0x87,
	0xf2, 0x0a, 0x82, 0x31, 0x02, 0xee, 0x0b, 0xd7, 0x22, 0x1e, 0xf7, 0xcf, 0x0f, 0x6a, 0x9d, 0x80,
	0x18, 0xa9, 0x93, 0x0a, 0xa6, 0x19, 0x92, 0x0c, 0xe3, 0xdf, 0xd7, 0xce, 0x8e, 0xe3, 0x3f, 0x53,
	0xd5, 0xa4, 0xf2, 0x39, 0xce, 0x9b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x18, 0x16, 0x92,
	0x8c, 0x01, 0x00, 0x00,
}
