package service

import (
	"context"
	"time"

	"course/internal/model"
)

type InfoCardService interface {
	CreateInfoCard(ctx context.Context, request *CreateInfoCardRequest) error
	ValidateInfoCard(ctx context.Context, request *ValidateInfoCardRequest) error
	GetInfoCard(ctx context.Context, request *GetInfoCardRequest) (*model.InfoCard, error)
	ListAllInfoCards(ctx context.Context, request *ListAllInfoCardsRequest) ([]*model.InfoCard, error)
	DeleteInfoCard(ctx context.Context, request *DeleteInfoCardRequest) error
}

type CreateInfoCardRequest struct {
	EmployeeID  int64
	IsConfirmed bool
	CreatedDate *time.Time
}

type ValidateInfoCardRequest struct {
	InfoCardID  int64
	IsConfirmed bool
}

type GetInfoCardRequest struct {
	InfoCardID int64
}

// ListAllInfoCardsRequest TODO: pagination, sort, filter
type ListAllInfoCardsRequest struct {
}

type DeleteInfoCardRequest struct {
	InfoCardID int64
}
