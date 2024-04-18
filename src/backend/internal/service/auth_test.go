package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"reflect"
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
		want    *model.Employee
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect company ID",
			a: &authServiceImpl{
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.RegisterEmployeeRequest{
					PhoneNumber: "3123124",
					FullName:    "Stepa Stepan Stepanovich",
					CompanyID:   -1,
					Post:        1,
					Password: &model.Password{
						Value:    "123",
						IsHashed: false,
					},
					DateOfBirth: nil,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				employeeStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request string
					}
					storageReturn struct {
						employee *model.Employee
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request string
					}{ctx: ctx, request: mock.Anything},
					storageReturn: struct {
						employee *model.Employee
						err      error
					}{
						employee: nil,
						err:      fmt.Errorf("incorrect companyID"),
					},
				},
			},
		},
		{
			name: "success",
			a: &authServiceImpl{
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.RegisterEmployeeRequest{
					PhoneNumber: "3123124",
					FullName:    "Stepa Stepan Stepanovich",
					CompanyID:   1,
					Post:        1,
					Password: &model.Password{
						Value:    "123",
						IsHashed: false,
					},
					DateOfBirth: nil,
				},
			},
			want: &model.Employee{
				ID:          model.ToEmployeeID(1),
				FullName:    "Stepa Stepan Stepanovich",
				PhoneNumber: "3123124",
				CompanyID:   model.ToCompanyID(1),
				Post:        model.ToPostTypeFromInt(1),
				Password: &model.Password{
					Value:    "123",
					IsHashed: true,
				},
				DateOfBirth: nil,
			},
			wantErr: false,

			storages: storages{
				employeeStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request string
					}
					storageReturn struct {
						employee *model.Employee
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request string
					}{ctx: ctx, request: mock.Anything},
					storageReturn: struct {
						employee *model.Employee
						err      error
					}{
						employee: &model.Employee{
							ID:          model.ToEmployeeID(1),
							FullName:    "Stepa Stepan Stepanovich",
							PhoneNumber: "3123124",
							CompanyID:   model.ToCompanyID(1),
							Post:        model.ToPostTypeFromInt(1),
							Password: &model.Password{
								Value:    "123",
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
			On("Register",
				tt.storages.employeeStorage.storageArgs.ctx,
				tt.storages.employeeStorage.storageArgs.request).
			Return(
				tt.storages.employeeStorage.storageReturn.employee,
				tt.storages.employeeStorage.storageReturn.err).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.RegisterEmployee(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("authServiceImpl.RegisterEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authServiceImpl.RegisterEmployee() = %v, want %v", got, tt.want)
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
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.LoginEmployeeRequest{
					PhoneNumber: "32423",
					Password: &model.Password{
						Value:    "gg",
						IsHashed: false,
					},
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
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.LoginEmployeeRequest{
					PhoneNumber: "124",
					Password: &model.Password{
						Value:    "124",
						IsHashed: false,
					},
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
							Password: &model.Password{
								Value:    "3d5f",
								IsHashed: true,
							},
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
				logger:          NewMockLogger(),
				employeeStorage: employeeMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.LoginEmployeeRequest{
					PhoneNumber: "124",
					Password: &model.Password{
						Value:    "21e12",
						IsHashed: false,
					},
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
							Password: &model.Password{
								Value:    string(pass),
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
			if err := tt.a.LoginEmployee(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("authServiceImpl.LoginEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
