package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=EmployeeStorage
type EmployeeStorage interface {
	Register(ctx context.Context, request *dto.RegisterEmployeeRequest) error
	GetByPhone(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error)
	ListAll(ctx context.Context, request *dto.ListAllEmployeesRequest) ([]*model.Employee, error)
	Delete(ctx context.Context, request *dto.DeleteEmployeeRequest) error
	Validate(ctx context.Context, request *dto.LoginEmployeeRequest) error
}
