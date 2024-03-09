package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type InfoCardStorage interface {
	Create(ctx context.Context, request *service.CreateInfoCardRequest) error
	Validate(ctx context.Context, request *service.ValidateInfoCardRequest) error
	GetByID(ctx context.Context, request *service.GetInfoCardRequest) (*model.InfoCard, error)
	List(ctx context.Context, request *service.ListInfoCardsRequest) ([]*model.InfoCard, error)
	Delete(ctx context.Context, request *service.DeleteInfoCardRequest) error
}

type infoCardStorageImpl struct {
	db *postgres.Postgres
}
