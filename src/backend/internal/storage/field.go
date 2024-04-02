package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=FieldStorage
type FieldStorage interface {
	Create(ctx context.Context, request *dto.CreateCardFieldRequest) error
	Get(ctx context.Context, request *dto.GetCardFieldRequest) (*model.Field, error)
	ListCardFields(ctx context.Context, request *dto.ListCardFieldsRequest) ([]*model.Field, error)
	Delete(ctx context.Context, request *dto.DeleteCardFieldRequest) error
}
