package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage/mocks"
)

func Test_employeeServiceImpl_GetEmployee(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.GetEmployeeRequest
	}

	type storages struct {
		employeeStorage struct {
			storageArgs   args
			storageReturn struct {
				employee *model.Employee
				err      error
			}
		}
	}

	employeeMockStorage := mocks.NewEmployeeStorage(t)
	tests := []struct {
		name    string
		e       *employeeServiceImpl
		args    args
		want    *model.Employee
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect phone number",
			e: &employeeServiceImpl{
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetEmployeeRequest{PhoneNumber: "kjc123"},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				employeeStorage: struct {
					storageArgs   args
					storageReturn struct {
						employee *model.Employee
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetEmployeeRequest{PhoneNumber: "kjc123"},
					},
					storageReturn: struct {
						employee *model.Employee
						err      error
					}{
						employee: nil,
						err:      fmt.Errorf("incorrect phone number"),
					},
				},
			},
		},
		{
			name: "success",
			e: &employeeServiceImpl{
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetEmployeeRequest{PhoneNumber: "123"},
			},
			want: &model.Employee{
				ID:          model.ToEmployeeID(1),
				FullName:    "Stepa Stepan Stepanovich",
				PhoneNumber: "123",
				CompanyID:   model.ToCompanyID(1),
				Post:        model.ToPostTypeFromInt(1),
				Password: &model.Password{
					Value:    "OHiuoup98u",
					IsHashed: true,
				},
				DateOfBirth: nil,
			},
			wantErr: false,

			storages: storages{
				employeeStorage: struct {
					storageArgs   args
					storageReturn struct {
						employee *model.Employee
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetEmployeeRequest{PhoneNumber: "123"},
					},
					storageReturn: struct {
						employee *model.Employee
						err      error
					}{
						employee: &model.Employee{
							ID:          model.ToEmployeeID(1),
							FullName:    "Stepa Stepan Stepanovich",
							PhoneNumber: "123",
							CompanyID:   model.ToCompanyID(1),
							Post:        model.ToPostTypeFromInt(1),
							Password: &model.Password{
								Value:    "OHiuoup98u",
								IsHashed: true,
							},
							DateOfBirth: nil,
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		employeeMockStorage.
			On("GetByPhone",
				tt.storages.employeeStorage.storageArgs.ctx,
				tt.storages.employeeStorage.storageArgs.request,
			).
			Return(
				tt.storages.employeeStorage.storageReturn.employee,
				tt.storages.employeeStorage.storageReturn.err,
			).
			Once()
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

func Test_employeeServiceImpl_DeleteEmployee(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.DeleteEmployeeRequest
	}

	type storages struct {
		employeeStorage struct {
			storageArgs   args
			storageReturn struct {
				err error
			}
		}
	}

	employeeMockStorage := mocks.NewEmployeeStorage(t)
	tests := []struct {
		name    string
		e       *employeeServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect employee ID",
			e: &employeeServiceImpl{
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteEmployeeRequest{EmployeeID: -1},
			},
			wantErr: true,

			storages: storages{
				employeeStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteEmployeeRequest{EmployeeID: -1},
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect employeeID"),
					},
				},
			},
		},
		{
			name: "success",
			e: &employeeServiceImpl{
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteEmployeeRequest{EmployeeID: 1},
			},
			wantErr: false,

			storages: storages{
				employeeStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteEmployeeRequest{EmployeeID: 1},
					},
					storageReturn: struct {
						err error
					}{
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		employeeMockStorage.
			On("Delete",
				tt.storages.employeeStorage.storageArgs.ctx,
				tt.storages.employeeStorage.storageArgs.request,
			).
			Return(
				tt.storages.employeeStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.DeleteEmployee(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("employeeServiceImpl.DeleteEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
