package postgres

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type EmployeeStorage interface {
	Register(ctx context.Context, request *service.RegisterEmployeeRequest) error
	GetByPhone(ctx context.Context, request *service.GetEmployeeRequest) (*model.Employee, error)
	ListAll(ctx context.Context, request *service.ListAllEmployeesRequest) ([]*model.Employee, error)
	Delete(ctx context.Context, request *service.DeleteEmployeeRequest) error
	Validate(ctx context.Context, request *service.LoginEmployeeRequest) error
}

type employeeStorageImpl struct {
	db *postgres.Postgres
}
