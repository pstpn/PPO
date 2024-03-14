package service

import (
	"context"
	"course/internal/model"
	"course/internal/service/dto"
	"reflect"
	"testing"
)

func Test_infoCardServiceImpl_CreateInfoCard(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.CreateInfoCardRequest
	}
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.CreateInfoCard(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.CreateInfoCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_infoCardServiceImpl_ValidateInfoCard(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.ValidateInfoCardRequest
	}
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.ValidateInfoCard(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.ValidateInfoCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_infoCardServiceImpl_GetInfoCard(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.GetInfoCardRequest
	}
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		want    *model.InfoCard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetInfoCard(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.GetInfoCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infoCardServiceImpl.GetInfoCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardServiceImpl_ListInfoCards(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.ListInfoCardsRequest
	}
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		want    []*model.InfoCard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.ListInfoCards(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.ListInfoCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infoCardServiceImpl.ListInfoCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardServiceImpl_DeleteInfoCard(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.DeleteInfoCardRequest
	}
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.DeleteInfoCard(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.DeleteInfoCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
