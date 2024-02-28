package service

type Service struct {
	Auth           AuthService
	CompanyService CompanyService
	Employee       EmployeeService
	InfoCard       InfoCardService
	Document       DocumentService
	Field          FieldService
	Photo          PhotoService
}

// ServiceImpl TODO
type ServiceImpl struct {
}
