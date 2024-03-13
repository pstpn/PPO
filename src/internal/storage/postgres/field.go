package postgres

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type FieldStorage interface {
	Create(ctx context.Context, request *service.CreateCardFieldRequest) error
	Get(ctx context.Context, request *service.GetCardFieldRequest) (*model.Field, error)
	ListCardFields(ctx context.Context, request *service.ListCardFieldsRequest) ([]*model.Field, error)
	Delete(ctx context.Context, request *service.DeleteCardFieldRequest) error
}

type fieldStorageImpl struct {
	db *postgres.Postgres
}
