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

func TestNewCompanyStorage(t *testing.T) {
	type args struct {
		db *postgres.Postgres
	}
	tests := []struct {
		name string
		args args
		want storage.CompanyStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompanyStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompanyStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_companyStorageImpl_Create(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.CreateCompanyRequest
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
			c := &companyStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := c.Create(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("companyStorageImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_companyStorageImpl_GetByID(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.GetCompanyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Company
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &companyStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := c.GetByID(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("companyStorageImpl.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("companyStorageImpl.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkpointStorageImpl_Delete(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.DeleteCompanyRequest
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
			c := &checkpointStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := c.Delete(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("checkpointStorageImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
