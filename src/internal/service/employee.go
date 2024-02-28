package service

import (
	"context"
	"time"

	"course/internal/model"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, request *CreateEmployeeRequest) error
	GetEmployee(ctx context.Context, request *GetEmployeeRequest) (*model.Employee, error)
	ListAllEmployees(ctx context.Context, request *ListAllEmployeesRequest) ([]*model.Employee, error)
	DeleteEmployee(ctx context.Context, request *DeleteEmployeeRequest) error
}

type CreateEmployeeRequest struct {
	PhoneNumber string
	FullName    string
	CompanyID   int64
	Post        int64
	DateOfBirth *time.Time
}

type GetEmployeeRequest struct {
	EmployeeID int64
}

// ListAllEmployeesRequest TODO: pagination, sort, filter
type ListAllEmployeesRequest struct {
}

type DeleteEmployeeRequest struct {
	EmployeeID int64
}
