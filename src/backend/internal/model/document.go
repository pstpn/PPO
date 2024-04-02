package model

type DocumentType int64

const (
	Passport DocumentType = iota
	DrivingLicense
)

func ToDocumentType(document int64) *DocumentType {
	documentType := DocumentType(document)
	return &documentType
}

func (d *DocumentType) String() string {
	switch *d {
	case Passport:
		return "Паспорт"
	case DrivingLicense:
		return "Водительские права"
	default:
		return "Неизвестный документ"
	}
}

type DocumentID int64

func ToDocumentID(id int64) *DocumentID {
	documentID := DocumentID(id)
	return &documentID
}

type Document struct {
	ID         *DocumentID
	InfoCardID *InfoCardID
	Type       *DocumentType
}
