package service

import (
	"context"
	"fmt"

	"course/internal/model"
	storage "course/internal/storage/postgres"
	"course/pkg/logger"
)

type EmployeeService interface {
	GetEmployee(ctx context.Context, request *GetEmployeeRequest) (*model.Employee, error)
	ListAllEmployees(ctx context.Context, request *ListAllEmployeesRequest) ([]*model.Employee, error)
	DeleteEmployee(ctx context.Context, request *DeleteEmployeeRequest) error
}

type employeeServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage
}

type GetEmployeeRequest struct {
	PhoneNumber string
}

func (e *employeeServiceImpl) GetEmployee(ctx context.Context, request *GetEmployeeRequest) (*model.Employee, error) {
	employee, err := e.employeeStorage.GetByPhone(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get employee: %w", err)
	}

	return employee, nil
}

// ListAllEmployeesRequest TODO: pagination, sort, filter
type ListAllEmployeesRequest struct {
}

func (e *employeeServiceImpl) ListAllEmployees(ctx context.Context, request *ListAllEmployeesRequest) ([]*model.Employee, error) {
	employees, err := e.employeeStorage.ListAll(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list all employees: %w", err)
	}

	return employees, nil
}

type DeleteEmployeeRequest struct {
	EmployeeID int64
}

func (e *employeeServiceImpl) DeleteEmployee(ctx context.Context, request *DeleteEmployeeRequest) error {
	err := e.employeeStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete employee: %w", err)
	}

	return nil
}
