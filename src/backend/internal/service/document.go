package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type DocumentService interface {
	CreateDocument(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error)
	GetDocument(ctx context.Context, request *dto.GetDocumentRequest) (*model.Document, error)
	ListEmployeeDocuments(ctx context.Context, request *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error)
	DeleteDocument(ctx context.Context, request *dto.DeleteDocumentRequest) error
}

type documentServiceImpl struct {
	logger          logger.Interface
	documentStorage storage.DocumentStorage
}

func (d *documentServiceImpl) CreateDocument(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error) {
	document, err := d.documentStorage.Create(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("create document: %w", err)
	}

	return document, nil
}

func (d *documentServiceImpl) GetDocument(ctx context.Context, request *dto.GetDocumentRequest) (*model.Document, error) {
	document, err := d.documentStorage.GetByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get document: %w", err)
	}

	return document, nil
}

func (d *documentServiceImpl) ListEmployeeDocuments(ctx context.Context, request *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error) {
	documents, err := d.documentStorage.List(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("list employee documents: %w", err)
	}

	return documents, nil
}

func (d *documentServiceImpl) DeleteDocument(ctx context.Context, request *dto.DeleteDocumentRequest) error {
	err := d.documentStorage.Delete(ctx, request)
	if err != nil {
		return fmt.Errorf("delete document: %w", err)
	}

	return nil
}
