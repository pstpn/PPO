package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

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

func (i *infoCardStorageImpl) Create(ctx context.Context, request *dto.CreateInfoCardRequest) (*model.InfoCard, error) {
	query := i.Builder.
		Insert(infoCardTable).
		Columns(
			createdEmployeePhoneNumberField,
			isConfirmedField,
			createdDateField,
		).
		Values(
			request.EmployeePhoneNumber,
			request.IsConfirmed,
			request.CreatedDate,
		).
		Suffix(returningInfoCardColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := i.Pool.QueryRow(ctx, sql, args...)

	return i.rowToModel(row)
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
			createdEmployeePhoneNumberField,
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

	return i.rowToModel(row)
}

func (i *infoCardStorageImpl) List(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.InfoCard, error) {
	query := i.Builder.
		Select(
			fullColName(infoCardTable, idField),
			createdEmployeePhoneNumberField,
			isConfirmedField,
			createdDateField,
		).
		From(infoCardTable).
		Join(on(
			infoCardTable,
			employeeTable,
			createdEmployeePhoneNumberField,
			phoneNumberField,
		))

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
		infoCard, err := i.rowToModel(rows)
		if err != nil {
			return nil, err
		}

		infoCards = append(infoCards, infoCard)
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

func (i *infoCardStorageImpl) rowToModel(row pgx.Row) (*model.InfoCard, error) {
	var infoCard model.InfoCard
	err := row.Scan(
		&infoCard.ID,
		&infoCard.CreatedEmployeePhoneNumber,
		&infoCard.IsConfirmed,
		&infoCard.CreatedDate,
	)
	if err != nil {
		return nil, err
	}

	return &infoCard, nil
}
