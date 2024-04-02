package dto

type GetEmployeeRequest struct {
	PhoneNumber string
}

type DeleteEmployeeRequest struct {
	EmployeeID int64
}
