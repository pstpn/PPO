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

func TestNewPhotoMetaStorage(t *testing.T) {
	type args struct {
		db *postgres.Postgres
	}
	tests := []struct {
		name string
		args args
		want storage.PhotoMetaStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPhotoMetaStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhotoMetaStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_photoMetaStorageImpl_SaveKey(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.CreatePhotoKeyRequest
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
			p := &photoMetaStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := p.SaveKey(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("photoMetaStorageImpl.SaveKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_photoMetaStorageImpl_GetKey(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.GetPhotoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.PhotoMeta
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &photoMetaStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := p.GetKey(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("photoMetaStorageImpl.GetKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("photoMetaStorageImpl.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_photoMetaStorageImpl_DeleteKey(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.DeletePhotoRequest
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
			p := &photoMetaStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := p.DeleteKey(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("photoMetaStorageImpl.DeleteKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
