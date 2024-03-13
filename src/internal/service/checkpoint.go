package service

import (
	"context"
	"fmt"
	"time"

	"course/internal/model"
	storage "course/internal/storage/postgres"
	"course/pkg/logger"
)

type CheckpointService interface {
	CreatePassage(ctx context.Context, request *CreatePassageRequest) error
	ListPassages(ctx context.Context, request *ListPassagesRequest) ([]*model.Passage, error)
}

type checkpointServiceImpl struct {
	logger            logger.Interface
	checkpointStorage storage.CheckpointStorage
}

type CreatePassageRequest struct {
	CheckpointID int64
	DocumentID   int64
	Type         string
	Time         *time.Time
}

func (c *checkpointServiceImpl) CreatePassage(ctx context.Context, request *CreatePassageRequest) error {
	err := c.checkpointStorage.CreatePassage(ctx, request)
	if err != nil {
		return fmt.Errorf("create passage: %w", err)
	}

	return nil
}

type ListPassagesRequest struct {
	CheckpointID int64
}

func (c *checkpointServiceImpl) ListPassages(ctx context.Context, request *ListPassagesRequest) ([]*model.Passage, error) {
	passages, err := c.checkpointStorage.ListPassages(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list passages: %w", err)
	}

	return passages, nil
}
