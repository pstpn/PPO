package service

import (
	"context"
	"course/internal/model"
	"course/internal/service/dto"
	"reflect"
	"testing"
)

func Test_photoServiceImpl_CreatePhoto(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.CreatePhotoRequest
	}
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.CreatePhoto(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("photoServiceImpl.CreatePhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_photoServiceImpl_GetPhoto(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.GetPhotoRequest
	}
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		want    *model.Photo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetPhoto(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("photoServiceImpl.GetPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("photoServiceImpl.GetPhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_photoServiceImpl_UpdatePhoto(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.UpdatePhotoRequest
	}
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.UpdatePhoto(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("photoServiceImpl.UpdatePhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_photoServiceImpl_DeletePhoto(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.DeletePhotoRequest
	}
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.DeletePhoto(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("photoServiceImpl.DeletePhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
