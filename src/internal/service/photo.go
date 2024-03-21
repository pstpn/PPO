package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, request *dto.CreatePhotoRequest) error
	GetPhoto(ctx context.Context, request *dto.GetPhotoRequest) (*model.Photo, error)
	UpdatePhoto(ctx context.Context, request *dto.UpdatePhotoRequest) error
	DeletePhoto(ctx context.Context, request *dto.DeletePhotoRequest) error
}

type photoServiceImpl struct {
	logger       logger.Interface
	photoStorage storage.PhotoStorage
}

func (p *photoServiceImpl) CreatePhoto(ctx context.Context, request *dto.CreatePhotoRequest) error {
	// TODO: Crop face from document

	key, err := p.photoStorage.Save(ctx, request.Data)
	if err != nil {
		return fmt.Errorf("save photo: %w", err)
	}

	err = p.photoStorage.SaveKey(ctx, &dto.CreatePhotoKeyRequest{
		DocumentID: model.ToDocumentID(request.DocumentID),
		Key:        key,
	})
	if err != nil {
		return fmt.Errorf("save photo key: %w", err)
	}

	return nil
}

func (p *photoServiceImpl) GetPhoto(ctx context.Context, request *dto.GetPhotoRequest) (*model.Photo, error) {
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

func (p *photoServiceImpl) UpdatePhoto(ctx context.Context, request *dto.UpdatePhotoRequest) error {
	meta, err := p.photoStorage.GetKey(ctx, &dto.GetPhotoRequest{DocumentID: request.DocumentID})
	if err != nil {
		return fmt.Errorf("get photo key: %w", err)
	}

	err = p.photoStorage.Update(ctx, meta.PhotoKey, request.Data)
	if err != nil {
		return fmt.Errorf("update photo by key: %w", err)
	}

	return nil
}

func (p *photoServiceImpl) DeletePhoto(ctx context.Context, request *dto.DeletePhotoRequest) error {
	meta, err := p.photoStorage.GetKey(ctx, &dto.GetPhotoRequest{DocumentID: request.DocumentID})
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
