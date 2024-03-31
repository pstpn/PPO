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
