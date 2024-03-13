package service

import (
	"context"
	"course/internal/model"
	"reflect"
	"testing"
)

func Test_fieldServiceImpl_CreateCardField(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *CreateCardFieldRequest
	}
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.CreateCardField(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.CreateCardField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fieldServiceImpl_GetCardField(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *GetCardFieldRequest
	}
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		want    *model.Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetCardField(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.GetCardField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldServiceImpl.GetCardField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldServiceImpl_ListCardFields(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *ListCardFieldsRequest
	}
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		want    []*model.Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.ListCardFields(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.ListCardFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldServiceImpl.ListCardFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldServiceImpl_DeleteCardField(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *DeleteCardFieldRequest
	}
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.DeleteCardField(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.DeleteCardField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
