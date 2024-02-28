package service

import (
	"context"
	"course/internal/model"
)

type FieldService interface {
	CreateCardField(ctx context.Context, request *CreateCardFieldRequest) error
	GetCardFields(ctx context.Context, request GetCardFieldsRequest) ([]*model.Field, error)
	GetCardField(ctx context.Context, request GetCardFieldRequest) ([]*model.Field, error)
	DeleteCardField(ctx context.Context, request *DeleteCardFieldRequest) error
}

type CreateCardFieldRequest struct {
	InfoCardID int64
	Type       int64
	Value      string
}

type GetCardFieldsRequest struct {
	InfoCardID int64
}

type GetCardFieldRequest struct {
	InfoCardID int64
	FieldType  int64
}

type DeleteCardFieldRequest struct {
	FieldID int64
}
