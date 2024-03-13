package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type PhotoStorage interface {
	SaveKey(ctx context.Context, documentID *model.DocumentID, photoKey *model.PhotoKey) error
	GetKey(ctx context.Context, request *service.GetPhotoRequest) (*model.PhotoMeta, error)
	UpdateKey(ctx context.Context, documentID *model.DocumentID, photoKey *model.PhotoKey) error
	DeleteKey(ctx context.Context, request *service.DeletePhotoRequest) error
}

type photoStorageImpl struct {
	db *postgres.Postgres
}
