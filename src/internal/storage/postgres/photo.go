package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type PhotoStorage interface {
	Save(ctx context.Context, request *service.CreatePhotoRequest, photoKey *model.PhotoKey) error
	Get(ctx context.Context, request *service.GetPhotoRequest) (*model.PhotoKey, error)
	Update(ctx context.Context, request *service.UpdatePhotoRequest, photoKey *model.PhotoKey) error
	Delete(ctx context.Context, request *service.DeletePhotoRequest) error
}

type photoStorageImpl struct {
	db *postgres.Postgres
}
