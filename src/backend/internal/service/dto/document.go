package dto

type CreateDocumentRequest struct {
	InfoCardID   int64
	DocumentType int64
}

type GetDocumentRequest struct {
	DocumentID int64
}

type ListEmployeeDocumentsRequest struct {
	EmployeeID int64
}

type DeleteDocumentRequest struct {
	DocumentID int64
}
