package model

import "time"

type Employee struct {
	ID *EmployeeID
}

type EmployeeID int64

type InfoCard struct {
	ID            *InfoCardID
	PhoneNumber   string
	FullName      string
	Birthday      *time.Time
	EmploymentDay *time.Time
	Verified      bool
}

type InfoCardID int64

type Field struct {
	ID   *FieldID
	Type *FieldType
}

type FieldID int64

type FieldType string

const (
	DateOfIssue FieldType = "date of issue"
)

type Change struct {
	ID         *ChangeID
	EmployeeID *EmployeeID
	FieldID    *FieldID
	OldValue   string
	NewValue   string
}

type ChangeID int64
