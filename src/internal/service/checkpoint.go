package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type CheckpointService interface {
	CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) error
	ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error)
}

type checkpointServiceImpl struct {
	logger            logger.Interface
	checkpointStorage storage.CheckpointStorage
}

func (c *checkpointServiceImpl) CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) error {
	err := c.checkpointStorage.CreatePassage(ctx, request)
	if err != nil {
		return fmt.Errorf("create passage: %w", err)
	}

	return nil
}

func (c *checkpointServiceImpl) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	passages, err := c.checkpointStorage.ListPassages(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list passages: %w", err)
	}

	return passages, nil
}
