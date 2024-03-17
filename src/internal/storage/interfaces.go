package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

type PhotoKeyStorage interface {
	SaveKey(ctx context.Context, request *dto.CreatePhotoKeyRequest) error
	GetKey(ctx context.Context, request *dto.GetPhotoRequest) (*model.PhotoMeta, error)
	UpdateKey(ctx context.Context, request *dto.UpdatePhotoKeyRequest) error
	DeleteKey(ctx context.Context, request *dto.DeletePhotoRequest) error
}

type PhotoDataStorage interface {
	Save(ctx context.Context, data []byte) (*model.PhotoKey, error)
	Get(ctx context.Context, key *model.PhotoKey) ([]byte, error)
	Update(ctx context.Context, key *model.PhotoKey, data []byte) error
	Delete(ctx context.Context, key *model.PhotoKey) error
}

//go:generate mockery --inpackage --name=PhotoStorage
type PhotoStorage interface {
	PhotoDataStorage
	PhotoKeyStorage
}

//go:generate mockery --inpackage --name=CheckpointStorage
type CheckpointStorage interface {
	CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) error
	ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error)
}

//go:generate mockery --inpackage --name=CompanyStorage
type CompanyStorage interface {
	GetByID(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error)
}

//go:generate mockery --inpackage --name=DocumentStorage
type DocumentStorage interface {
	Create(ctx context.Context, request *dto.CreateDocumentRequest) error
	GetByID(ctx context.Context, request *dto.GetDocumentRequest) (*model.Document, error)
	List(ctx context.Context, request *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error)
	Delete(ctx context.Context, request *dto.DeleteDocumentRequest) error
}

//go:generate mockery --inpackage --name=EmployeeStorage
type EmployeeStorage interface {
	Register(ctx context.Context, request *dto.RegisterEmployeeRequest) error
	GetByPhone(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error)
	ListAll(ctx context.Context, request *dto.ListAllEmployeesRequest) ([]*model.Employee, error)
	Delete(ctx context.Context, request *dto.DeleteEmployeeRequest) error
	Validate(ctx context.Context, request *dto.LoginEmployeeRequest) error
}

//go:generate mockery --inpackage --name=FieldStorage
type FieldStorage interface {
	Create(ctx context.Context, request *dto.CreateCardFieldRequest) error
	Get(ctx context.Context, request *dto.GetCardFieldRequest) (*model.Field, error)
	ListCardFields(ctx context.Context, request *dto.ListCardFieldsRequest) ([]*model.Field, error)
	Delete(ctx context.Context, request *dto.DeleteCardFieldRequest) error
}

//go:generate mockery --inpackage --name=InfoCardStorage
type InfoCardStorage interface {
	Create(ctx context.Context, request *dto.CreateInfoCardRequest) error
	Validate(ctx context.Context, request *dto.ValidateInfoCardRequest) error
	GetByID(ctx context.Context, request *dto.GetInfoCardRequest) (*model.InfoCard, error)
	List(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.InfoCard, error)
	Delete(ctx context.Context, request *dto.DeleteInfoCardRequest) error
}
