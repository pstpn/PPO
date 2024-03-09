package service

import (
	"context"
	"fmt"

	"course/internal/model"
	storage "course/internal/storage/postgres"
	"course/pkg/logger"
)

type DocumentService interface {
	CreateDocument(ctx context.Context, request *CreateDocumentRequest) error
	GetDocument(ctx context.Context, request *GetDocumentRequest) (*model.Document, error)
	ListEmployeeDocuments(ctx context.Context, request *ListEmployeeDocumentsRequest) ([]*model.Document, error)
	DeleteDocument(ctx context.Context, request *DeleteDocumentRequest) error
}

type documentServiceImpl struct {
	logger          logger.Interface
	documentStorage storage.DocumentStorage
}

type CreateDocumentRequest struct {
	InfoCardID   int64
	DocumentType int64
}

func (d *documentServiceImpl) CreateDocument(ctx context.Context, request *CreateDocumentRequest) error {
	err := d.documentStorage.Create(ctx, request)
	if err != nil {
		return fmt.Errorf("create document: %w", err)
	}

	return nil
}

type GetDocumentRequest struct {
	DocumentID int64
}

func (d *documentServiceImpl) GetDocument(ctx context.Context, request *GetDocumentRequest) (*model.Document, error) {
	document, err := d.documentStorage.GetByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get document: %w", err)
	}

	return document, nil
}

type ListEmployeeDocumentsRequest struct {
	EmployeeID int64
}

func (d *documentServiceImpl) ListEmployeeDocuments(ctx context.Context, request *ListEmployeeDocumentsRequest) ([]*model.Document, error) {
	documents, err := d.documentStorage.List(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list employee documents: %w", err)
	}

	return documents, nil
}

type DeleteDocumentRequest struct {
	DocumentID int64
}

func (d *documentServiceImpl) DeleteDocument(ctx context.Context, request *DeleteDocumentRequest) error {
	err := d.documentStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete document: %w", err)
	}

	return nil
}
