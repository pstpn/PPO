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

func TestNewFieldStorage(t *testing.T) {
	type args struct {
		db *postgres.Postgres
	}
	tests := []struct {
		name string
		args args
		want storage.FieldStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFieldStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFieldStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldStorageImpl_Create(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.CreateDocumentFieldRequest
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
			f := &fieldStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := f.Create(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("fieldStorageImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fieldStorageImpl_Get(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.GetDocumentFieldRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fieldStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := f.Get(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("fieldStorageImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldStorageImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldStorageImpl_ListCardFields(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.ListDocumentFieldsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fieldStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := f.ListCardFields(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("fieldStorageImpl.ListCardFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldStorageImpl.ListCardFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldStorageImpl_Delete(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.DeleteDocumentFieldRequest
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
			f := &fieldStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := f.Delete(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("fieldStorageImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
