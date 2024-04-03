package model

type FieldID int64

func ToFieldID(id int64) *FieldID {
	fieldID := FieldID(id)
	return &fieldID
}

type FieldType int64

const (
	DateOfRelease FieldType = iota
	UnknownFieldType
)

func ToFieldTypeFromInt(field int64) *FieldType {
	fieldType := FieldType(field)
	return &fieldType
}

func ToFieldTypeFromString(field string) *FieldType {
	var fieldType FieldType
	switch field {
	case "Дата выпуска":
		fieldType = DateOfRelease
	default:
		fieldType = UnknownFieldType
	}

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
