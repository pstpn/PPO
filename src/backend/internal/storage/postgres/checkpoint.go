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

type checkpointStorageImpl struct {
	*postgres.Postgres
}

func NewCheckpointStorage(db *postgres.Postgres) storage.CheckpointStorage {
	return &checkpointStorageImpl{db}
}

func (c *checkpointStorageImpl) CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) (*model.Passage, error) {
	query := c.Builder.
		Insert(passageTable).
		Columns(
			checkpointIdField,
			documentIdField,
			typeField,
			timeField,
		).
		Values(
			request.CheckpointID,
			request.DocumentID,
			model.ToPassageType(request.Type).String(),
			request.Time,
		).
		Suffix(returningPassageColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(ctx, sql, args...)

	return c.rowToModel(row)
}

func (c *checkpointStorageImpl) GetPassage(ctx context.Context, request *dto.GetPassageRequest) (*model.Passage, error) {
	query := c.Builder.
		Select(
			idField,
			checkpointIdField,
			documentIdField,
			typeField,
			timeField,
		).
		From(passageTable).
		Where(squirrel.Eq{idField: request.PassageID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(ctx, sql, args...)

	return c.rowToModel(row)
}

func (c *checkpointStorageImpl) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	query := c.Builder.
		Select(
			fullColName(passageTable, idField),
			checkpointIdField,
			documentIdField,
			fullColName(passageTable, typeField),
			timeField,
		).
		From(passageTable).
		Join(on(
			passageTable,
			documentTable,
			documentIdField,
			idField,
		)).
		Where(squirrel.Eq{infoCardIdField: request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := c.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var passages []*model.Passage
	for rows.Next() {
		passage, err := c.rowToModel(rows)
		if err != nil {
			return nil, err
		}

		passages = append(passages, passage)
	}

	return passages, nil
}

func (c *checkpointStorageImpl) DeletePassage(ctx context.Context, request *dto.DeletePassageRequest) error {
	query := c.Builder.
		Delete(passageTable).
		Where(squirrel.Eq{idField: request.PassageID})

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

func (c *checkpointStorageImpl) rowToModel(row pgx.Row) (*model.Passage, error) {
	var passage model.Passage
	err := row.Scan(
		&passage.ID,
		&passage.CheckpointID,
		&passage.DocumentID,
		&passage.Type,
		&passage.Time,
	)
	if err != nil {
		return nil, err
	}

	return &passage, nil
}
