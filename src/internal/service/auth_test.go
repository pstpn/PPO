package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/stretchr/testify/mock"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage/mocks"
)

func Test_authServiceImpl_RegisterEmployee(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.RegisterEmployeeRequest
	}

	type storages struct {
		employeeStorage struct {
			storageArgs struct {
				ctx     context.Context
				request string
			}
			storageReturn struct {
				err error
			}
		}
	}

	employeeMockStorage := mocks.NewEmployeeStorage(t)
	tests := []struct {
		name    string
		a       *authServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect company ID",
			a: &authServiceImpl{
				logger:          nil,
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.RegisterEmployeeRequest{
					PhoneNumber: "3123124",
					FullName:    "Stepa Stepan Stepanovich",
					CompanyID:   -1,
					Post:        1,
					Password:    "",
					DateOfBirth: nil,
				},
			},
			wantErr: true,

			storages: storages{
				employeeStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request string
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request string
					}{ctx: ctx, request: mock.Anything},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect companyID"),
					},
				},
			},
		},
		{
			name: "success",
			a: &authServiceImpl{
				logger:          nil,
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.RegisterEmployeeRequest{
					PhoneNumber: "3123124",
					FullName:    "Stepa Stepan Stepanovich",
					CompanyID:   1,
					Post:        1,
					Password:    "",
					DateOfBirth: nil,
				},
			},
			wantErr: false,

			storages: storages{
				employeeStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request string
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request string
					}{ctx: ctx, request: mock.Anything},
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
			On("Register",
				tt.storages.employeeStorage.storageArgs.ctx,
				tt.storages.employeeStorage.storageArgs.request).
			Return(tt.storages.employeeStorage.storageReturn.err).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.RegisterEmployee(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("authServiceImpl.RegisterEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authServiceImpl_LoginEmployee(t *testing.T) {
	ctx := context.TODO()
	pass, _ := bcrypt.GenerateFromPassword([]byte("21e12"), bcrypt.DefaultCost)

	type args struct {
		ctx     context.Context
		request *dto.LoginEmployeeRequest
	}

	type storages struct {
		employeeStorage struct {
			storageArgs struct {
				ctx     context.Context
				request *dto.GetEmployeeRequest
			}
			storageReturn struct {
				employee *model.Employee
				err      error
			}
		}
	}

	employeeMockStorage := mocks.NewEmployeeStorage(t)
	tests := []struct {
		name    string
		a       *authServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect phone number",
			a: &authServiceImpl{
				logger:          nil,
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.LoginEmployeeRequest{
					PhoneNumber: "32423",
					Password:    "gg",
				},
			},
			wantErr: true,

			storages: storages{
				employeeStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.GetEmployeeRequest
					}
					storageReturn struct {
						employee *model.Employee
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.GetEmployeeRequest
					}{
						ctx:     ctx,
						request: &dto.GetEmployeeRequest{PhoneNumber: "32423"},
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
			name: "incorrect password",
			a: &authServiceImpl{
				logger:          nil,
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.LoginEmployeeRequest{
					PhoneNumber: "124",
					Password:    "124",
				},
			},
			wantErr: true,

			storages: storages{
				employeeStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.GetEmployeeRequest
					}
					storageReturn struct {
						employee *model.Employee
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.GetEmployeeRequest
					}{
						ctx:     ctx,
						request: &dto.GetEmployeeRequest{PhoneNumber: "124"},
					},
					storageReturn: struct {
						employee *model.Employee
						err      error
					}{
						employee: &model.Employee{
							ID:          nil,
							FullName:    "",
							PhoneNumber: "",
							CompanyID:   nil,
							Post:        nil,
							Password:    "124",
							DateOfBirth: nil,
						},
						err: nil,
					},
				},
			},
		},
		{
			name: "success",
			a: &authServiceImpl{
				logger:          nil,
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.LoginEmployeeRequest{
					PhoneNumber: "124",
					Password:    "21e12",
				},
			},
			wantErr: false,

			storages: storages{
				employeeStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.GetEmployeeRequest
					}
					storageReturn struct {
						employee *model.Employee
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.GetEmployeeRequest
					}{
						ctx:     ctx,
						request: &dto.GetEmployeeRequest{PhoneNumber: "124"},
					},
					storageReturn: struct {
						employee *model.Employee
						err      error
					}{
						employee: &model.Employee{
							ID:          nil,
							FullName:    "",
							PhoneNumber: "",
							CompanyID:   nil,
							Post:        nil,
							Password:    string(pass),
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
			if err := tt.a.LoginEmployee(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("authServiceImpl.LoginEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
