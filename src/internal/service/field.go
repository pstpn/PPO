package service

import (
	"context"
	"fmt"

	"course/internal/model"
	storage "course/internal/storage/postgres"
	"course/pkg/logger"
)

type FieldService interface {
	CreateCardField(ctx context.Context, request *CreateCardFieldRequest) error
	GetCardField(ctx context.Context, request *GetCardFieldRequest) (*model.Field, error)
	ListCardFields(ctx context.Context, request *ListCardFieldsRequest) ([]*model.Field, error)
	DeleteCardField(ctx context.Context, request *DeleteCardFieldRequest) error
}

type fieldServiceImpl struct {
	logger       logger.Interface
	fieldStorage storage.FieldStorage
}

type CreateCardFieldRequest struct {
	InfoCardID int64
	Type       int64
	Value      string
}

func (f *fieldServiceImpl) CreateCardField(ctx context.Context, request *CreateCardFieldRequest) error {
	err := f.fieldStorage.Create(ctx, request)
	if err != nil {
		return fmt.Errorf("create info card field: %w", err)
	}

	return nil
}

type GetCardFieldRequest struct {
	InfoCardID int64
	FieldType  int64
}

func (f *fieldServiceImpl) GetCardField(ctx context.Context, request *GetCardFieldRequest) (*model.Field, error) {
	field, err := f.fieldStorage.Get(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get info card field: %w", err)
	}

	return field, nil
}

type ListCardFieldsRequest struct {
	InfoCardID int64
}

func (f *fieldServiceImpl) ListCardFields(ctx context.Context, request *ListCardFieldsRequest) ([]*model.Field, error) {
	fields, err := f.fieldStorage.ListCardFields(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list info card fields: %w", err)
	}

	return fields, nil
}

type DeleteCardFieldRequest struct {
	FieldID int64
}

func (f *fieldServiceImpl) DeleteCardField(ctx context.Context, request *DeleteCardFieldRequest) error {
	err := f.fieldStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete info card field: %w", err)
	}

	return nil
}
