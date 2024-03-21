package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=InfoCardStorage
type InfoCardStorage interface {
	Create(ctx context.Context, request *dto.CreateInfoCardRequest) error
	Validate(ctx context.Context, request *dto.ValidateInfoCardRequest) error
	GetByID(ctx context.Context, request *dto.GetInfoCardRequest) (*model.InfoCard, error)
	List(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.InfoCard, error)
	Delete(ctx context.Context, request *dto.DeleteInfoCardRequest) error
}
