package dto

type CreateCardFieldRequest struct {
	InfoCardID int64
	Type       int64
	Value      string
}

type GetCardFieldRequest struct {
	InfoCardID int64
	FieldType  int64
}
type ListCardFieldsRequest struct {
	InfoCardID int64
}

type DeleteCardFieldRequest struct {
	FieldID int64
}
