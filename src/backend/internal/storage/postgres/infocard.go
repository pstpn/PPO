package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

type infoCardStorageImpl struct {
	*postgres.Postgres
}

func NewInfoCardStorage(db *postgres.Postgres) storage.InfoCardStorage {
	return &infoCardStorageImpl{db}
}

func (i *infoCardStorageImpl) Create(ctx context.Context, request *dto.CreateInfoCardRequest) error {
	query := i.Builder.
		Insert(infoCardTable).
		Columns(
			createdEmployeeIdField,
			isConfirmedField,
			createdDateField,
		).
		Values(
			request.EmployeeID,
			request.IsConfirmed,
			request.CreatedDate,
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = i.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (i *infoCardStorageImpl) Validate(ctx context.Context, request *dto.ValidateInfoCardRequest) error {
	query := i.Builder.
		Update(infoCardTable).
		Set(isConfirmedField, request.IsConfirmed).
		Where(squirrel.Eq{idField: request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = i.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (i *infoCardStorageImpl) GetByID(ctx context.Context, request *dto.GetInfoCardRequest) (*model.InfoCard, error) {
	query := i.Builder.
		Select(
			idField,
			createdEmployeeIdField,
			isConfirmedField,
			createdDateField,
		).
		From(infoCardTable).
		Where(squirrel.Eq{idField: request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := i.Pool.QueryRow(ctx, sql, args...)

	var infoCard model.InfoCard
	err = row.Scan(
		&infoCard.ID,
		&infoCard.CreatedEmployeeID,
		&infoCard.IsConfirmed,
		&infoCard.CreatedDate,
	)
	if err != nil {
		return nil, err
	}

	return &infoCard, nil
}

func (i *infoCardStorageImpl) List(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.InfoCard, error) {
	query := i.Builder.
		Select(
			idField,
			createdEmployeeIdField,
			isConfirmedField,
			createdDateField,
		).
		From(infoCardTable)

	query = request.Pagination.ToSQL(query)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := i.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var infoCards []*model.InfoCard
	for rows.Next() {
		var infoCard model.InfoCard
		err = rows.Scan(
			&infoCard.ID,
			&infoCard.CreatedEmployeeID,
			&infoCard.IsConfirmed,
			&infoCard.CreatedDate,
		)
		if err != nil {
			return nil, err
		}

		infoCards = append(infoCards, &infoCard)
	}

	return infoCards, nil
}

func (i *infoCardStorageImpl) Delete(ctx context.Context, request *dto.DeleteInfoCardRequest) error {
	query := i.Builder.
		Delete(infoCardTable).
		Where(squirrel.Eq{idField: request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = i.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
