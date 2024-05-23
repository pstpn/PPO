package dto

type CreateDocumentFieldRequest struct {
	DocumentID string
	Type       int64
	Value      string
}

type GetDocumentFieldRequest struct {
	DocumentID string
	FieldType  int64
}

type ListDocumentFieldsRequest struct {
	DocumentID string
}

type DeleteDocumentFieldRequest struct {
	FieldID string
}
