package model

type FieldID int64

func ToFieldID(id int64) *FieldID {
	fieldID := FieldID(id)
	return &fieldID
}

type FieldType int64

const (
	DateOfRelease FieldType = iota
)

func ToFieldType(field int64) *FieldType {
	fieldType := FieldType(field)
	return &fieldType
}

func (f *FieldType) String() string {
	switch *f {
	case DateOfRelease:
		return "Дата выпуска"
	default:
		return "Неизвестное поле"
	}
}

type Field struct {
	ID         *FieldID
	DocumentID *DocumentID
	Type       *FieldType
	Value      string
}
