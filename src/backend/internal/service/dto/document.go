package dto

type CreateDocumentRequest struct {
	SerialNumber string
	InfoCardID   int64
	DocumentType int64
}

type GetDocumentRequest struct {
	DocumentID int64
}

type ListEmployeeDocumentsRequest struct {
	EmployeePhoneNumber string
}

type DeleteDocumentRequest struct {
	DocumentID int64
}
