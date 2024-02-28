package service

import (
	"context"

	"course/internal/model"
)

type DocumentService interface {
	CreateDocument(ctx context.Context, request *CreateDocumentRequest) error
	GetDocument(ctx context.Context, request *GetDocumentRequest) (*model.Document, error)
	GetEmployeeDocuments(ctx context.Context, request *GetEmployeeDocumentsRequest) ([]*model.Document, error)
	DeleteDocument(ctx context.Context, request *DeleteDocumentRequest) error
}

type CreateDocumentRequest struct {
	InfoCardID   int64
	DocumentType int64
}

type GetDocumentRequest struct {
	DocumentID int64
}

type GetEmployeeDocumentsRequest struct {
	EmployeeID int64
}

type DeleteDocumentRequest struct {
	DocumentID int64
}
