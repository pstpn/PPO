package mongodb

import (
	"context"

	"course/internal/model"
	"course/internal/service"
)

type PhotoStorage interface {
	Save(ctx context.Context, request *service.CreatePhotoRequest) (*model.PhotoKey, error)
	Get(ctx context.Context, request *service.GetPhotoRequest) (*model.PhotoData, error)
	Update(ctx context.Context, request *service.UpdatePhotoRequest) (*model.PhotoKey, error)
	Delete(ctx context.Context, key *model.PhotoKey) error
}

type photoStorageImpl struct {
	db *mongodb
}
