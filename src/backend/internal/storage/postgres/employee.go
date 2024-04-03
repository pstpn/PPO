package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

type employeeStorageImpl struct {
	*postgres.Postgres
}

func NewEmployeeStorage(db *postgres.Postgres) storage.EmployeeStorage {
	return &employeeStorageImpl{db}
}

func (e *employeeStorageImpl) Register(ctx context.Context, request *dto.RegisterEmployeeRequest) error {
	query := e.Builder.
		Insert(employeeTable).
		Columns(
			phoneNumberField,
			fullNameField,
			companyIdField,
			postField,
			passwordField,
			dateOfBirthField,
		).
		Values(
			request.PhoneNumber,
			request.FullName,
			request.CompanyID,
			request.Post,
			request.Password,
			request.DateOfBirth,
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeStorageImpl) GetByPhone(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error) {
	query := e.Builder.
		Select(
			idField,
			phoneNumberField,
			fullNameField,
			companyIdField,
			postField,
			passwordField,
			dateOfBirthField,
		).
		From(employeeTable).
		Where(squirrel.Eq{phoneNumberField: request.PhoneNumber})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := e.Pool.QueryRow(ctx, sql, args...)

	var employee model.Employee
	var hashedPassword string
	err = row.Scan(
		&employee.ID,
		&employee.PhoneNumber,
		&employee.FullName,
		&employee.CompanyID,
		&employee.Post,
		&hashedPassword,
		&employee.DateOfBirth,
	)
	if err != nil {
		return nil, err
	}

	employee.Password = &model.Password{
		Value:    hashedPassword,
		IsHashed: true,
	}

	return &employee, nil
}

func (e *employeeStorageImpl) Delete(ctx context.Context, request *dto.DeleteEmployeeRequest) error {
	query := e.Builder.
		Delete(employeeTable).
		Where(squirrel.Eq{idField: request.EmployeeID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
