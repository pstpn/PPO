package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/storage"
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
	// TODO: Crop face from document

	key, err := p.photoStorage.Save(ctx, request.Data)
	if err != nil {
		return fmt.Errorf("save photo: %w", err)
	}

	err = p.photoStorage.SaveKey(ctx, model.ToDocumentID(request.DocumentID), key)
	if err != nil {
		return fmt.Errorf("save photo key: %w", err)
	}

	return nil
}

type GetPhotoRequest struct {
	DocumentID int64
}

func (p *photoServiceImpl) GetPhoto(ctx context.Context, request *GetPhotoRequest) (*model.Photo, error) {
	meta, err := p.photoStorage.GetKey(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get photo key: %w", err)
	}

	data, err := p.photoStorage.Get(ctx, meta.PhotoKey)
	if err != nil {
		return nil, fmt.Errorf("get photo by key: %w", err)
	}

	return &model.Photo{
		Meta: meta,
		Data: data,
	}, nil
}

type UpdatePhotoRequest struct {
	DocumentID int64
	Data       []byte
}

func (p *photoServiceImpl) UpdatePhoto(ctx context.Context, request *UpdatePhotoRequest) error {
	meta, err := p.photoStorage.GetKey(ctx, &GetPhotoRequest{DocumentID: request.DocumentID})
	if err != nil {
		return fmt.Errorf("get photo key: %w", err)
	}

	err = p.photoStorage.Update(ctx, meta.PhotoKey, request.Data)
	if err != nil {
		return fmt.Errorf("update photo by key: %w", err)
	}

	return nil
}

type DeletePhotoRequest struct {
	DocumentID int64
}

func (p *photoServiceImpl) DeletePhoto(ctx context.Context, request *DeletePhotoRequest) error {
	meta, err := p.photoStorage.GetKey(ctx, &GetPhotoRequest{DocumentID: request.DocumentID})
	if err != nil {
		return fmt.Errorf("get photo key: %w", err)
	}

	err = p.photoStorage.Delete(ctx, meta.PhotoKey)
	if err != nil {
		return fmt.Errorf("delete photo by key: %w", err)
	}

	err = p.photoStorage.DeleteKey(ctx, request)
	if err != nil {
		return fmt.Errorf("delete photo key: %w", err)
	}

	return nil
}
