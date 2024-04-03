package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

type fieldStorageImpl struct {
	*postgres.Postgres
}

func NewFieldStorage(db *postgres.Postgres) storage.FieldStorage {
	return &fieldStorageImpl{db}
}

func (f *fieldStorageImpl) Create(ctx context.Context, request *dto.CreateDocumentFieldRequest) error {
	query := f.Builder.
		Insert(fieldTable).
		Columns(
			documentIdField,
			typeField,
			valueField,
		).
		Values(
			request.DocumentID,
			request.Type,
			request.Value,
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = f.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (f *fieldStorageImpl) Get(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error) {
	query := f.Builder.
		Select(
			idField,
			documentIdField,
			typeField,
			valueField,
		).
		From(fieldTable).
		Where(
			squirrel.And{
				squirrel.Eq{documentIdField: request.DocumentID},
				squirrel.Eq{typeField: request.FieldType},
			},
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := f.Pool.QueryRow(ctx, sql, args...)

	var field model.Field
	err = row.Scan(
		&field.ID,
		&field.DocumentID,
		&field.Type,
		&field.Value,
	)
	if err != nil {
		return nil, err
	}

	return &field, nil
}

func (f *fieldStorageImpl) ListCardFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error) {
	query := f.Builder.
		Select(
			idField,
			documentIdField,
			typeField,
			valueField,
		).
		From(fieldTable).
		Where(squirrel.Eq{documentIdField: request.DocumentID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := f.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var fields []*model.Field
	for rows.Next() {
		var field model.Field
		err = rows.Scan(
			&field.ID,
			&field.DocumentID,
			&field.Type,
			&field.Value,
		)
		if err != nil {
			return nil, err
		}

		fields = append(fields, &field)
	}

	return fields, nil
}

func (f *fieldStorageImpl) Delete(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error {
	query := f.Builder.
		Delete(fieldTable).
		Where(squirrel.Eq{idField: request.FieldID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = f.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
