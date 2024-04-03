package postgres

import (
	"context"
	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
	"reflect"
	"testing"
)

func TestNewDocumentStorage(t *testing.T) {
	type args struct {
		db *postgres.Postgres
	}
	tests := []struct {
		name string
		args args
		want storage.DocumentStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDocumentStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDocumentStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_documentStorageImpl_Create(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.CreateDocumentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &documentStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := d.Create(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("documentStorageImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_documentStorageImpl_GetByID(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.GetDocumentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Document
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &documentStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := d.GetByID(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("documentStorageImpl.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("documentStorageImpl.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_documentStorageImpl_List(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.ListEmployeeDocumentsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Document
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &documentStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := d.List(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("documentStorageImpl.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("documentStorageImpl.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_documentStorageImpl_Delete(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.DeleteDocumentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &documentStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := d.Delete(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("documentStorageImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
