package model

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
