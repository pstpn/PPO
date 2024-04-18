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

func NewDocumentService(logger logger.Interface, documentStorage storage.DocumentStorage) DocumentService {
	return &documentServiceImpl{
		logger:          logger,
		documentStorage: documentStorage,
	}
}

func (d *documentServiceImpl) CreateDocument(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error) {
	d.logger.Infof("create document with serial number %s", request.SerialNumber)

	document, err := d.documentStorage.Create(ctx, request)
	if err != nil {
		d.logger.Errorf("create document: %s", err.Error())
		return nil, fmt.Errorf("create document: %w", err)
	}

	return document, nil
}

func (d *documentServiceImpl) GetDocument(ctx context.Context, request *dto.GetDocumentRequest) (*model.Document, error) {
	d.logger.Infof("get document by ID %d", request.DocumentID)

	document, err := d.documentStorage.GetByID(ctx, request)
	if err != nil {
		d.logger.Errorf("get document: %s", err.Error())
		return nil, fmt.Errorf("get document: %w", err)
	}

	return document, nil
}

func (d *documentServiceImpl) ListEmployeeDocuments(ctx context.Context, request *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error) {
	d.logger.Infof("list employee documents by employee ID %d", request.EmployeeID)

	documents, err := d.documentStorage.List(ctx, request)
	if err != nil {
		d.logger.Errorf("list employee documents: %s", err.Error())
		return nil, fmt.Errorf("list employee documents: %w", err)
	}

	return documents, nil
}

func (d *documentServiceImpl) DeleteDocument(ctx context.Context, request *dto.DeleteDocumentRequest) error {
	d.logger.Infof("delete document by ID %d", request.DocumentID)

	err := d.documentStorage.Delete(ctx, request)
	if err != nil {
		d.logger.Errorf("delete document: %s", err.Error())
		return fmt.Errorf("delete document: %w", err)
	}

	return nil
}
