package postgres

import (
	"context"
	"strings"
	"time"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

const connURL = "postgresql://postgres:admin@localhost:5432/tests"

func NewTestStorage() *postgres.Postgres {
	conn, err := postgres.New(connURL)
	if err != nil {
		panic(err)
	}

	initTestCompanyStorage(NewCompanyStorage(conn))
	initTestEmployeeStorage(NewEmployeeStorage(conn))
	initTestInfoCardStorage(NewInfoCardStorage(conn))
	initTestDocumentStorage(NewDocumentStorage(conn))
	initTestFieldStorage(NewFieldStorage(conn))
	initTestCheckpointStorage(NewCheckpointStorage(conn))
	initTestPhotoMetaStorage(NewPhotoMetaStorage(conn))

	return conn
}

func initTestCompanyStorage(storage storage.CompanyStorage) {
	_, err := storage.Create(context.TODO(), &dto.CreateCompanyRequest{
		Name: "Test",
		City: "Test",
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}
}

func initTestEmployeeStorage(storage storage.EmployeeStorage) {
	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	_, err := storage.Register(context.TODO(), &dto.RegisterEmployeeRequest{
		PhoneNumber: "123",
		FullName:    "123",
		CompanyID:   1,
		Post:        1,
		Password: &model.Password{
			Value:    "123",
			IsHashed: true,
		},
		DateOfBirth: &tm,
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}
}

func initTestInfoCardStorage(storage storage.InfoCardStorage) {
	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	_, err := storage.Create(context.TODO(), &dto.CreateInfoCardRequest{
		EmployeeID:  1,
		IsConfirmed: true,
		CreatedDate: &tm,
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}
}

func initTestDocumentStorage(storage storage.DocumentStorage) {
	_, err := storage.Create(context.TODO(), &dto.CreateDocumentRequest{
		InfoCardID:   1,
		DocumentType: 1,
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}
}

func initTestFieldStorage(storage storage.FieldStorage) {
	_, err := storage.Create(context.TODO(), &dto.CreateDocumentFieldRequest{
		DocumentID: 1,
		Type:       1,
		Value:      "123",
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}
}

func initTestCheckpointStorage(storage storage.CheckpointStorage) {
	//tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	_, err := storage.CreateCheckpoint(context.TODO(), &dto.CreateCheckpointRequest{
		PhoneNumber: "123123",
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}
	//_, err = storage.CreatePassage(context.TODO(), &dto.CreatePassageRequest{
	//	CheckpointID: 1,
	//	DocumentID:   1,
	//	Type:         1,
	//	Time:         &tm,
	//})
	//if err != nil && !strings.Contains(err.Error(), "constraint") {
	//	panic(err)
	//}
}

func initTestPhotoMetaStorage(storage storage.PhotoMetaStorage) {
	_, err := storage.SaveKey(context.TODO(), &dto.CreatePhotoKeyRequest{
		DocumentID: model.ToDocumentID(1),
		Key:        model.ToPhotoKey("123321"),
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}
}
