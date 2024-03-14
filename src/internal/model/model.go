package model

import "time"

type PostType int64

const (
	SecurityEmployee PostType = iota
	DefaultEmployee
)

func (p *PostType) String() string {
	switch *p {
	case SecurityEmployee:
		return "Сотрудник СБ"
	case DefaultEmployee:
		return "Сотрудник"
	default:
		return "Неизвестная должность"
	}
}

func (p *PostType) IsAdmin() bool {
	return *p == 1
}

type CompanyID int64

type Company struct {
	ID   *CompanyID
	Name string
	City string
}

type EmployeeID int64

type Employee struct {
	ID          *EmployeeID
	FullName    string
	PhoneNumber string
	CompanyID   *CompanyID
	Post        *PostType
	Password    string
	DateOfBirth *time.Time
}

type InfoCardID int64

type InfoCard struct {
	ID                *InfoCardID
	CreatedEmployeeID *EmployeeID
	IsConfirmed       bool
	CreatedDate       *time.Time
}

type PhotoID int64

type PhotoKey string

func ToPhotoKey(key string) *PhotoKey {
	photoKey := PhotoKey(key)
	return &photoKey
}

type PhotoMeta struct {
	PhotoID  *PhotoID
	PhotoKey *PhotoKey
}

type Photo struct {
	Meta *PhotoMeta
	Data []byte
}

type DocumentType int64

const (
	Passport DocumentType = iota
	DrivingLicense
)

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
	PhotoID    *PhotoID
}

type FieldID int64

type FieldType int64

const (
	DateOfRelease FieldType = iota
)

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

type CheckpointID int64

func ToCheckpointID(id int64) *CheckpointID {
	checkpointID := CheckpointID(id)
	return &checkpointID
}

type PassageType int64

const (
	Entrance PassageType = iota
	Exit
)

func ToPassageType(passage int64) *PassageType {
	passageType := PassageType(passage)
	return &passageType
}

type Passage struct {
	CheckpointID *CheckpointID
	DocumentID   *DocumentID
	Type         *PassageType
	Time         *time.Time
}
