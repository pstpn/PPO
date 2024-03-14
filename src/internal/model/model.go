package model

import "time"

type PostType int64

const (
	SecurityEmployee PostType = iota
	DefaultEmployee
)

func ToPostType(post int64) *PostType {
	postType := PostType(post)
	return &postType
}

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

func ToCompanyID(id int64) *CompanyID {
	companyID := CompanyID(id)
	return &companyID
}

type Company struct {
	ID   *CompanyID
	Name string
	City string
}

type EmployeeID int64

func ToEmployeeID(id int64) *EmployeeID {
	employeeID := EmployeeID(id)
	return &employeeID
}

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

func ToInfoCardID(id int64) *InfoCardID {
	infoCardID := InfoCardID(id)
	return &infoCardID
}

type InfoCard struct {
	ID                *InfoCardID
	CreatedEmployeeID *EmployeeID
	IsConfirmed       bool
	CreatedDate       *time.Time
}

type PhotoID int64

func ToPhotoID(id int64) *PhotoID {
	photoID := PhotoID(id)
	return &photoID
}

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
	PhotoID    *PhotoID
}

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
