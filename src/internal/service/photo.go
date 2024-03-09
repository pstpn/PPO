package service

import (
	"context"
	storage "course/internal/storage/postgres"

	"course/internal/model"
	"course/pkg/logger"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, request *CreatePhotoRequest) error
	GetPhoto(ctx context.Context, request *GetPhotoRequest) (*model.Photo, error)
	UpdatePhoto(ctx context.Context, request *UpdatePhotoRequest) error
	DeletePhoto(ctx context.Context, request *DeletePhotoRequest) error
}

type photoServiceImpl struct {
	logger       logger.Interface
	photoStorage storage.PhotoStorage
}

type CreatePhotoRequest struct {
	DocumentID int64
	Data       []byte
}

func (p *photoServiceImpl) CreatePhoto(ctx context.Context, request *CreatePhotoRequest) error {

}

type GetPhotoRequest struct {
	DocumentID int64
}

type UpdatePhotoRequest struct {
	DocumentID int64
	Data       []byte
}

type DeletePhotoRequest struct {
	PhotoID int64
}
