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
