package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type EmployeeService interface {
	GetEmployee(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error)
	DeleteEmployee(ctx context.Context, request *dto.DeleteEmployeeRequest) error
}

type employeeServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage
}

func (e *employeeServiceImpl) GetEmployee(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error) {
	employee, err := e.employeeStorage.GetByPhone(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get employee: %w", err)
	}

	return employee, nil
}

func (e *employeeServiceImpl) DeleteEmployee(ctx context.Context, request *dto.DeleteEmployeeRequest) error {
	err := e.employeeStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete employee: %w", err)
	}

	return nil
}
