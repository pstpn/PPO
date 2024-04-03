package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type FieldService interface {
	CreateCardField(ctx context.Context, request *dto.CreateDocumentFieldRequest) error
	GetCardField(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error)
	ListCardFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error)
	DeleteCardField(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error
}

type fieldServiceImpl struct {
	logger       logger.Interface
	fieldStorage storage.FieldStorage
}

func (f *fieldServiceImpl) CreateCardField(ctx context.Context, request *dto.CreateDocumentFieldRequest) error {
	err := f.fieldStorage.Create(ctx, request)
	if err != nil {
		return fmt.Errorf("create info card field: %w", err)
	}

	return nil
}

func (f *fieldServiceImpl) GetCardField(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error) {
	field, err := f.fieldStorage.Get(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get info card field: %w", err)
	}

	return field, nil
}

func (f *fieldServiceImpl) ListCardFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error) {
	fields, err := f.fieldStorage.ListCardFields(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list info card fields: %w", err)
	}

	return fields, nil
}

func (f *fieldServiceImpl) DeleteCardField(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error {
	err := f.fieldStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete info card field: %w", err)
	}

	return nil
}
