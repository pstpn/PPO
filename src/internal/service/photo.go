package service

import (
	"context"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, request *CreatePhotoRequest) error
	UpdatePhoto(ctx context.Context, request *UpdatePhotoRequest) error
	DeletePhoto(ctx context.Context, request *DeletePhotoRequest) error
}

type CreatePhotoRequest struct {
	DocumentID int64
	Key        string
}

type UpdatePhotoRequest struct {
	PhotoID int64
	Key     string
}

type DeletePhotoRequest struct {
	PhotoID int64
}
