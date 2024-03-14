package service

import (
	"context"
	"course/internal/model"
	"course/internal/service/dto"
	"reflect"
	"testing"
)

func Test_checkpointServiceImpl_CreatePassage(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.CreatePassageRequest
	}
	tests := []struct {
		name    string
		c       *checkpointServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.CreatePassage(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("checkpointServiceImpl.CreatePassage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkpointServiceImpl_ListPassages(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *dto.ListPassagesRequest
	}
	tests := []struct {
		name    string
		c       *checkpointServiceImpl
		args    args
		want    []*model.Passage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ListPassages(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkpointServiceImpl.ListPassages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkpointServiceImpl.ListPassages() = %v, want %v", got, tt.want)
			}
		})
	}
}
