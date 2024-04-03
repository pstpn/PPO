package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

type documentStorageImpl struct {
	*postgres.Postgres
}

func NewDocumentStorage(db *postgres.Postgres) storage.DocumentStorage {
	return &documentStorageImpl{db}
}

func (d *documentStorageImpl) Create(ctx context.Context, request *dto.CreateDocumentRequest) error {
	query := d.Builder.
		Insert(documentTable).
		Columns(
			infoCardIdField,
			typeField,
		).
		Values(
			request.InfoCardID,
			request.DocumentType,
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = d.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (d *documentStorageImpl) GetByID(ctx context.Context, request *dto.GetDocumentRequest) (*model.Document, error) {
	query := d.Builder.
		Select(
			idField,
			infoCardIdField,
			typeField,
		).
		From(documentTable).
		Where(squirrel.Eq{idField: request.DocumentID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := d.Pool.QueryRow(ctx, sql, args...)

	var document model.Document
	err = row.Scan(&document.ID, &document.InfoCardID, &document.Type)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (d *documentStorageImpl) List(ctx context.Context, request *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error) {
	query := d.Builder.
		Select(
			fullColName(documentTable, idField),
			infoCardIdField,
			fullColName(documentTable, typeField),
		).
		From(documentTable).
		Join(on(
			documentTable,
			infoCardTable,
			infoCardIdField,
			idField,
		)).
		Where(squirrel.Eq{createdEmployeeIdField: request.EmployeeID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := d.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var documents []*model.Document
	for rows.Next() {
		var document model.Document
		err = rows.Scan(
			&document.ID,
			&document.InfoCardID,
			&document.Type,
		)
		if err != nil {
			return nil, err
		}

		documents = append(documents, &document)
	}

	return documents, nil
}

func (d *documentStorageImpl) Delete(ctx context.Context, request *dto.DeleteDocumentRequest) error {
	query := d.Builder.
		Delete(documentTable).
		Where(squirrel.Eq{idField: request.DocumentID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = d.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
