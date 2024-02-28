package service

import (
	"context"
)

type FieldService interface {
	CreateCardField(ctx context.Context, request CreateCardFieldRequest) error
	DeleteCardField(ctx context.Context, request *DeleteCardFieldRequest) error
}

type CreateCardFieldRequest struct {
	InfoCardID int64
	Type       int64
	Value      string
}

type DeleteCardFieldRequest struct {
	FieldID int64
}
