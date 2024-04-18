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

func Test_companyServiceImpl_GetCompany(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.GetCompanyRequest
	}

	type storages struct {
		companyStorage struct {
			storageArgs   args
			storageReturn struct {
				company *model.Company
				err     error
			}
		}
	}

	companyMockStorage := mocks.NewCompanyStorage(t)
	tests := []struct {
		name    string
		c       *companyServiceImpl
		args    args
		want    *model.Company
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect company ID",
			c: &companyServiceImpl{
				logger:         NewMockLogger(),
				companyStorage: companyMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetCompanyRequest{CompanyID: -1},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				companyStorage: struct {
					storageArgs   args
					storageReturn struct {
						company *model.Company
						err     error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetCompanyRequest{CompanyID: -1},
					},
					storageReturn: struct {
						company *model.Company
						err     error
					}{
						company: nil,
						err:     fmt.Errorf("incorrect companyID"),
					},
				},
			},
		},
		{
			name: "success",
			c: &companyServiceImpl{
				logger:         NewMockLogger(),
				companyStorage: companyMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetCompanyRequest{CompanyID: 1},
			},
			want: &model.Company{
				ID:   model.ToCompanyID(1),
				Name: "123",
				City: "123",
			},
			wantErr: false,

			storages: storages{
				companyStorage: struct {
					storageArgs   args
					storageReturn struct {
						company *model.Company
						err     error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetCompanyRequest{CompanyID: 1},
					},
					storageReturn: struct {
						company *model.Company
						err     error
					}{
						company: &model.Company{
							ID:   model.ToCompanyID(1),
							Name: "123",
							City: "123",
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		companyMockStorage.
			On("GetByID",
				tt.storages.companyStorage.storageArgs.ctx,
				tt.storages.companyStorage.storageArgs.request,
			).
			Return(
				tt.storages.companyStorage.storageReturn.company,
				tt.storages.companyStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetCompany(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("companyServiceImpl.GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("companyServiceImpl.GetCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}
