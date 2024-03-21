package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type InfoCardService interface {
	CreateInfoCard(ctx context.Context, request *dto.CreateInfoCardRequest) error
	ValidateInfoCard(ctx context.Context, request *dto.ValidateInfoCardRequest) error
	GetInfoCard(ctx context.Context, request *dto.GetInfoCardRequest) (*model.InfoCard, error)
	ListInfoCards(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.InfoCard, error)
	DeleteInfoCard(ctx context.Context, request *dto.DeleteInfoCardRequest) error
}

type infoCardServiceImpl struct {
	logger          logger.Interface
	infoCardStorage storage.InfoCardStorage
}

func (i *infoCardServiceImpl) CreateInfoCard(ctx context.Context, request *dto.CreateInfoCardRequest) error {
	err := i.infoCardStorage.Create(ctx, request)
	if err != nil {
		return fmt.Errorf("create info card: %w", err)
	}

	return nil
}

func (i *infoCardServiceImpl) ValidateInfoCard(ctx context.Context, request *dto.ValidateInfoCardRequest) error {
	err := i.infoCardStorage.Validate(ctx, request)
	if err != nil {
		return fmt.Errorf("validate info card: %w", err)
	}

	return nil
}

func (i *infoCardServiceImpl) GetInfoCard(ctx context.Context, request *dto.GetInfoCardRequest) (*model.InfoCard, error) {
	infoCard, err := i.infoCardStorage.GetByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get info card: %w", err)
	}

	return infoCard, nil
}

func (i *infoCardServiceImpl) ListInfoCards(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.InfoCard, error) {
	infoCards, err := i.infoCardStorage.List(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list info cards: %w", err)
	}

	return infoCards, nil
}

func (i *infoCardServiceImpl) DeleteInfoCard(ctx context.Context, request *dto.DeleteInfoCardRequest) error {
	err := i.infoCardStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete info card: %w", err)
	}

	return nil
}
