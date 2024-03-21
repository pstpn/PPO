package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=CheckpointStorage
type CheckpointStorage interface {
	CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) error
	ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error)
}
