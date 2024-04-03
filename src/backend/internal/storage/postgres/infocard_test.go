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

func TestNewInfoCardStorage(t *testing.T) {
	type args struct {
		db *postgres.Postgres
	}
	tests := []struct {
		name string
		args args
		want storage.InfoCardStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInfoCardStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInfoCardStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardStorageImpl_Create(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.CreateInfoCardRequest
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
			i := &infoCardStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := i.Create(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardStorageImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_infoCardStorageImpl_Validate(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.ValidateInfoCardRequest
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
			i := &infoCardStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := i.Validate(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardStorageImpl.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_infoCardStorageImpl_GetByID(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.GetInfoCardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.InfoCard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &infoCardStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := i.GetByID(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoCardStorageImpl.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infoCardStorageImpl.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardStorageImpl_List(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.ListInfoCardsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.InfoCard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &infoCardStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := i.List(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoCardStorageImpl.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infoCardStorageImpl.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardStorageImpl_Delete(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.DeleteInfoCardRequest
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
			i := &infoCardStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := i.Delete(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardStorageImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
