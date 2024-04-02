package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

type PhotoKeyStorage interface {
	SaveKey(ctx context.Context, request *dto.CreatePhotoKeyRequest) error
	GetKey(ctx context.Context, request *dto.GetPhotoRequest) (*model.PhotoMeta, error)
	UpdateKey(ctx context.Context, request *dto.UpdatePhotoKeyRequest) error
	DeleteKey(ctx context.Context, request *dto.DeletePhotoRequest) error
}

type PhotoDataStorage interface {
	Save(ctx context.Context, data []byte) (*model.PhotoKey, error)
	Get(ctx context.Context, key *model.PhotoKey) ([]byte, error)
	Update(ctx context.Context, key *model.PhotoKey, data []byte) error
	Delete(ctx context.Context, key *model.PhotoKey) error
}

//go:generate mockery --name=PhotoStorage
type PhotoStorage interface {
	PhotoDataStorage
	PhotoKeyStorage
}
