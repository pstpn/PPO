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

func Test_checkpointServiceImpl_CreatePassage(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.CreatePassageRequest
	}

	type storages struct {
		checkpointStorage struct {
			storageArgs   args
			storageReturn struct {
				passage *model.Passage
				err     error
			}
		}
	}

	checkpointMockStorage := mocks.NewCheckpointStorage(t)
	tests := []struct {
		name    string
		c       *checkpointServiceImpl
		args    args
		want    *model.Passage
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect checkpoint ID",
			c: &checkpointServiceImpl{
				logger:            NewMockLogger(),
				checkpointStorage: checkpointMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreatePassageRequest{
					CheckpointID: -1,
					DocumentID:   1,
					Type:         1,
					Time:         nil,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				checkpointStorage: struct {
					storageArgs   args
					storageReturn struct {
						passage *model.Passage
						err     error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreatePassageRequest{
							CheckpointID: -1,
							DocumentID:   1,
							Type:         1,
							Time:         nil,
						},
					},
					storageReturn: struct {
						passage *model.Passage
						err     error
					}{
						passage: nil,
						err:     fmt.Errorf("incorrect checkpointID"),
					},
				},
			},
		},
		{
			name: "success",
			c: &checkpointServiceImpl{
				logger:            NewMockLogger(),
				checkpointStorage: checkpointMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreatePassageRequest{
					CheckpointID: 1,
					DocumentID:   1,
					Type:         1,
					Time:         nil,
				},
			},
			want: &model.Passage{
				ID:           model.ToPassageID(1),
				CheckpointID: model.ToCheckpointID(1),
				DocumentID:   model.ToDocumentID(1),
				Type:         model.ToPassageTypeFromInt(1),
				Time:         nil,
			},
			wantErr: false,

			storages: storages{
				checkpointStorage: struct {
					storageArgs   args
					storageReturn struct {
						passage *model.Passage
						err     error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreatePassageRequest{
							CheckpointID: 1,
							DocumentID:   1,
							Type:         1,
							Time:         nil,
						},
					},
					storageReturn: struct {
						passage *model.Passage
						err     error
					}{
						passage: &model.Passage{
							ID:           model.ToPassageID(1),
							CheckpointID: model.ToCheckpointID(1),
							DocumentID:   model.ToDocumentID(1),
							Type:         model.ToPassageTypeFromInt(1),
							Time:         nil,
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		checkpointMockStorage.
			On("CreatePassage",
				tt.storages.checkpointStorage.storageArgs.ctx,
				tt.storages.checkpointStorage.storageArgs.request,
			).
			Return(
				tt.storages.checkpointStorage.storageReturn.passage,
				tt.storages.checkpointStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreatePassage(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkpointServiceImpl.CreatePassage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkpointServiceImpl.CreatePassage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkpointServiceImpl_ListPassages(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.ListPassagesRequest
	}

	type storages struct {
		checkpointStorage struct {
			storageArgs   args
			storageReturn struct {
				passages []*model.Passage
				err      error
			}
		}
	}

	checkpointMockStorage := mocks.NewCheckpointStorage(t)
	tests := []struct {
		name    string
		c       *checkpointServiceImpl
		args    args
		want    []*model.Passage
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect checkpoint ID",
			c: &checkpointServiceImpl{
				logger:            NewMockLogger(),
				checkpointStorage: checkpointMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListPassagesRequest{InfoCardID: -1},
			},
			wantErr: true,

			storages: storages{
				checkpointStorage: struct {
					storageArgs   args
					storageReturn struct {
						passages []*model.Passage
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListPassagesRequest{InfoCardID: -1},
					},
					storageReturn: struct {
						passages []*model.Passage
						err      error
					}{
						passages: nil,
						err:      fmt.Errorf("incorrect InfoCardID"),
					},
				},
			},
		},
		{
			name: "success",
			c: &checkpointServiceImpl{
				logger:            NewMockLogger(),
				checkpointStorage: checkpointMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListPassagesRequest{InfoCardID: 1},
			},
			want: []*model.Passage{
				{
					ID:           model.ToPassageID(1),
					CheckpointID: model.ToCheckpointID(1),
					DocumentID:   model.ToDocumentID(1),
					Type:         model.ToPassageTypeFromInt(1),
					Time:         nil,
				},
			},
			wantErr: false,

			storages: storages{
				checkpointStorage: struct {
					storageArgs   args
					storageReturn struct {
						passages []*model.Passage
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListPassagesRequest{InfoCardID: 1},
					},
					storageReturn: struct {
						passages []*model.Passage
						err      error
					}{
						passages: []*model.Passage{
							{
								ID:           model.ToPassageID(1),
								CheckpointID: model.ToCheckpointID(1),
								DocumentID:   model.ToDocumentID(1),
								Type:         model.ToPassageTypeFromInt(1),
								Time:         nil,
							},
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		checkpointMockStorage.
			On("ListPassages",
				tt.storages.checkpointStorage.storageArgs.ctx,
				tt.storages.checkpointStorage.storageArgs.request,
			).
			Return(
				tt.storages.checkpointStorage.storageReturn.passages,
				tt.storages.checkpointStorage.storageReturn.err,
			).
			Once()
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
