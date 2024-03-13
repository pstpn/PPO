package postgres

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type CheckpointStorage interface {
	CreatePassage(ctx context.Context, request *service.CreatePassageRequest) error
	ListPassages(ctx context.Context, request *service.ListPassagesRequest) ([]*model.Passage, error)
}

type checkpointStorageImpl struct {
	db *postgres.Postgres
}
