package postgres

import (
	"context"

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

func (c *checkpointStorageImpl) CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) error {
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

func (c *checkpointStorageImpl) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	return nil, nil
}
