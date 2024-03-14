package service

import (
	"context"
	"course/internal/model"
	"course/internal/service/dto"
	"reflect"
	"testing"
)

func Test_companyServiceImpl_GetCompany(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.GetCompanyRequest
	}
	tests := []struct {
		name    string
		c       *companyServiceImpl
		args    args
		want    *model.Company
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
