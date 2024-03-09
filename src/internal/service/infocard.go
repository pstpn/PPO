package service

import (
	"context"
	"fmt"
	"time"

	"course/internal/model"
	storage "course/internal/storage/postgres"
	"course/pkg/logger"
)

type InfoCardService interface {
	CreateInfoCard(ctx context.Context, request *CreateInfoCardRequest) error
	ValidateInfoCard(ctx context.Context, request *ValidateInfoCardRequest) error
	GetInfoCard(ctx context.Context, request *GetInfoCardRequest) (*model.InfoCard, error)
	ListInfoCards(ctx context.Context, request *ListInfoCardsRequest) ([]*model.InfoCard, error)
	DeleteInfoCard(ctx context.Context, request *DeleteInfoCardRequest) error
}

type infoCardServiceImpl struct {
	logger          logger.Interface
	infoCardStorage storage.InfoCardStorage
}

type CreateInfoCardRequest struct {
	EmployeeID  int64
	IsConfirmed bool
	CreatedDate *time.Time
}

func (i *infoCardServiceImpl) CreateInfoCard(ctx context.Context, request *CreateInfoCardRequest) error {
	err := i.infoCardStorage.Create(ctx, request)
	if err != nil {
		return fmt.Errorf("create info card: %w", err)
	}

	return nil
}

type ValidateInfoCardRequest struct {
	InfoCardID  int64
	IsConfirmed bool
}

func (i *infoCardServiceImpl) ValidateInfoCard(ctx context.Context, request *ValidateInfoCardRequest) error {
	err := i.infoCardStorage.Validate(ctx, request)
	if err != nil {
		return fmt.Errorf("validate info card: %w", err)
	}

	return nil
}

type GetInfoCardRequest struct {
	InfoCardID int64
}

func (i *infoCardServiceImpl) GetInfoCard(ctx context.Context, request *GetInfoCardRequest) (*model.InfoCard, error) {
	infoCard, err := i.infoCardStorage.GetByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get info card: %w", err)
	}

	return infoCard, nil
}

// ListInfoCardsRequest TODO: pagination, sort, filter
type ListInfoCardsRequest struct {
}

func (i *infoCardServiceImpl) ListInfoCards(ctx context.Context, request *ListInfoCardsRequest) ([]*model.InfoCard, error) {
	infoCards, err := i.infoCardStorage.List(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list info cards: %w", err)
	}

	return infoCards, nil
}

type DeleteInfoCardRequest struct {
	InfoCardID int64
}

func (i *infoCardServiceImpl) DeleteInfoCard(ctx context.Context, request *DeleteInfoCardRequest) error {
	err := i.infoCardStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete info card: %w", err)
	}

	return nil
}
