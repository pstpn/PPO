package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type DocumentStorage interface {
	Create(ctx context.Context, request *service.CreateDocumentRequest) error
	GetByID(ctx context.Context, request *service.GetDocumentRequest) (*model.Document, error)
	List(ctx context.Context, request *service.ListEmployeeDocumentsRequest) ([]*model.Document, error)
	Delete(ctx context.Context, request *service.DeleteDocumentRequest) error
}

type documentStorageImpl struct {
	db *postgres.Postgres
}
