package postgres

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"course/internal/model"
	"course/internal/service/dto"
	"course/pkg/storage/postgres"
)

func Test_checkpointStorageImpl_CreatePassage(t *testing.T) {
	testDB := NewTestStorage()
	defer testDB.Close()
	checkpointStorage := NewCheckpointStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	request := &dto.CreatePassageRequest{
		CheckpointID: 1,
		DocumentID:   1,
		Type:         0,
		Time:         &tm,
	}

	err := checkpointStorage.CreatePassage(context.TODO(), request)

	require.NoError(t, err)

	passage, err := checkpointStorage.GetPassage(context.TODO(), &dto.GetPassageRequest{PassageID: 1})
}

func Test_checkpointStorageImpl_GetPassage(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.GetPassageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Passage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &checkpointStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := c.GetPassage(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkpointStorageImpl.GetPassage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkpointStorageImpl.GetPassage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkpointStorageImpl_ListPassages(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.ListPassagesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Passage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &checkpointStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			got, err := c.ListPassages(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkpointStorageImpl.ListPassages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkpointStorageImpl.ListPassages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkpointStorageImpl_DeletePassage(t *testing.T) {
	type fields struct {
		Postgres *postgres.Postgres
	}
	type args struct {
		ctx     context.Context
		request *dto.DeletePassageRequest
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
			c := &checkpointStorageImpl{
				Postgres: tt.fields.Postgres,
			}
			if err := c.DeletePassage(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("checkpointStorageImpl.DeletePassage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
