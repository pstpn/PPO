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

type documentStorageImpl struct {
	*postgres.Postgres
}

func NewDocumentStorage(db *postgres.Postgres) storage.DocumentStorage {
	return &documentStorageImpl{db}
}

func (d *documentStorageImpl) Create(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error) {
	query := d.Builder.
		Insert(documentTable).
		Columns(
			serialNumberField,
			infoCardIdField,
			typeField,
		).
		Values(
			request.SerialNumber,
			request.InfoCardID,
			model.ToDocumentTypeFromInt(request.DocumentType).String(),
		).
		Suffix(returningDocumentColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := d.Pool.QueryRow(ctx, sql, args...)

	return d.rowToModel(row)
}

func (d *documentStorageImpl) GetByID(ctx context.Context, request *dto.GetDocumentRequest) (*model.Document, error) {
	query := d.Builder.
		Select(
			idField,
			serialNumberField,
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

	return d.rowToModel(row)
}

func (d *documentStorageImpl) List(ctx context.Context, request *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error) {
	query := d.Builder.
		Select(
			fullColName(documentTable, idField),
			serialNumberField,
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
		document, err := d.rowToModel(rows)
		if err != nil {
			return nil, err
		}

		documents = append(documents, document)
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

func (d *documentStorageImpl) rowToModel(row pgx.Row) (*model.Document, error) {
	var document model.Document
	var documentType string
	err := row.Scan(
		&document.ID,
		&document.SerialNumber,
		&document.InfoCardID,
		&documentType,
	)
	if err != nil {
		return nil, err
	}
	document.Type = model.ToDocumentTypeFromString(documentType)

	return &document, nil
}
