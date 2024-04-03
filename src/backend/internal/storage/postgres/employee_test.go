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

func TestNewEmployeeStorage(t *testing.T) {
	type args struct {
		db *postgres.Postgres
	}
	tests := []struct {
		name string
		args args
		want storage.EmployeeStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeStorageImpl_Register(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.RegisterEmployeeRequest
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
			e := &employeeStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := e.Register(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("employeeStorageImpl.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_employeeStorageImpl_GetByPhone(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.GetEmployeeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employeeStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := e.GetByPhone(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("employeeStorageImpl.GetByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeStorageImpl.GetByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeStorageImpl_Delete(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.DeleteEmployeeRequest
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
			e := &employeeStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := e.Delete(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("employeeStorageImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
