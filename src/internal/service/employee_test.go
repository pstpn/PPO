package service

import (
	"context"
	"course/internal/model"
	"reflect"
	"testing"
)

func Test_employeeServiceImpl_GetEmployee(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *GetEmployeeRequest
	}
	tests := []struct {
		name    string
		e       *employeeServiceImpl
		args    args
		want    *model.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.GetEmployee(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("employeeServiceImpl.GetEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeServiceImpl.GetEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeServiceImpl_ListAllEmployees(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *ListAllEmployeesRequest
	}
	tests := []struct {
		name    string
		e       *employeeServiceImpl
		args    args
		want    []*model.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.ListAllEmployees(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("employeeServiceImpl.ListAllEmployees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeServiceImpl.ListAllEmployees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeServiceImpl_DeleteEmployee(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *DeleteEmployeeRequest
	}
	tests := []struct {
		name    string
		e       *employeeServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.DeleteEmployee(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("employeeServiceImpl.DeleteEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
