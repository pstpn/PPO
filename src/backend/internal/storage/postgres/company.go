package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

type companyStorageImpl struct {
	*postgres.Postgres
}

func NewCompanyStorage(db *postgres.Postgres) storage.CompanyStorage {
	return &companyStorageImpl{db}
}

func (c *companyStorageImpl) Create(ctx context.Context, request *dto.CreateCompanyRequest) error {
	query := c.Builder.
		Insert(companyTable).
		Columns(
			nameField,
			cityField,
		).
		Values(
			request.Name,
			request.City,
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = c.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyStorageImpl) GetByID(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error) {
	query := c.Builder.
		Select(
			idField,
			nameField,
			cityField,
		).
		From(companyTable).
		Where(squirrel.Eq{idField: request.CompanyID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(ctx, sql, args...)

	var company model.Company
	err = row.Scan(&company.ID, &company.Name, &company.City)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (c *checkpointStorageImpl) Delete(ctx context.Context, request *dto.DeleteCompanyRequest) error {
	query := c.Builder.
		Delete(companyTable).
		Where(squirrel.Eq{idField: request.CompanyID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = c.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
