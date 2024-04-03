package model

import "time"

type PostType int64

const (
	SecurityEmployee PostType = iota
	DefaultEmployee
	UnknownEmployee
)

func ToPostTypeFromInt(post int64) *PostType {
	postType := PostType(post)
	return &postType
}

func ToPostTypeFromString(post string) *PostType {
	var postType PostType
	switch post {
	case "Сотрудник СБ":
		postType = SecurityEmployee
	case "Сотрудник":
		postType = DefaultEmployee
	default:
		postType = UnknownEmployee
	}

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
	return *p == SecurityEmployee
}

type EmployeeID int64

func ToEmployeeID(id int64) *EmployeeID {
	employeeID := EmployeeID(id)
	return &employeeID
}

type Password struct {
	Value    string
	IsHashed bool
}

type Employee struct {
	ID          *EmployeeID
	FullName    string
	PhoneNumber string
	CompanyID   *CompanyID
	Post        *PostType
	Password    *Password
	DateOfBirth *time.Time
}
