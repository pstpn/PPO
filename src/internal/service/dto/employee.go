package dto

type GetEmployeeRequest struct {
	PhoneNumber string
}

// ListAllEmployeesRequest TODO: pagination, sort, filter
type ListAllEmployeesRequest struct {
}

type DeleteEmployeeRequest struct {
	EmployeeID int64
}
