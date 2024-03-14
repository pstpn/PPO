package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
)

func Test_photoServiceImpl_CreatePhoto(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.CreatePhotoRequest
	}

	type storages struct {
		photoKeyStorage struct {
			storageArgs struct {
				ctx     context.Context
				request *dto.CreatePhotoKeyRequest
			}
			storageReturn struct {
				err error
			}
		}
		photoStorage struct {
			storageArgs struct {
				ctx  context.Context
				data []byte
			}
			storageReturn struct {
				photoKey *model.PhotoKey
				err      error
			}
		}
	}

	photoMockStorage := storage.NewMockPhotoStorage(t)
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect photo data",
			p: &photoServiceImpl{
				logger:       nil,
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreatePhotoRequest{
					DocumentID: -1,
					Data:       nil,
				},
			},
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoKeyRequest{
							DocumentID: model.ToDocumentID(1),
							Key:        model.ToPhotoKey("soso"),
						},
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect documentID"),
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx  context.Context
						data []byte
					}
					storageReturn struct {
						photoKey *model.PhotoKey
						err      error
					}
				}{
					storageArgs: struct {
						ctx  context.Context
						data []byte
					}{
						ctx:  ctx,
						data: nil,
					},
					storageReturn: struct {
						photoKey *model.PhotoKey
						err      error
					}{
						photoKey: nil,
						err:      fmt.Errorf("incorrect photo data"),
					},
				},
			},
		},
		{
			name: "incorrect document ID",
			p: &photoServiceImpl{
				logger:       nil,
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreatePhotoRequest{
					DocumentID: -1,
					Data:       nil,
				},
			},
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoKeyRequest{
							DocumentID: model.ToDocumentID(-1),
							Key:        model.ToPhotoKey("soso"),
						},
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect documentID"),
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx  context.Context
						data []byte
					}
					storageReturn struct {
						photoKey *model.PhotoKey
						err      error
					}
				}{
					storageArgs: struct {
						ctx  context.Context
						data []byte
					}{
						ctx:  ctx,
						data: nil,
					},
					storageReturn: struct {
						photoKey *model.PhotoKey
						err      error
					}{
						photoKey: model.ToPhotoKey("soso"),
						err:      nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		photoMockStorage.
			On("Save",
				tt.storages.photoStorage.storageArgs.ctx,
				tt.storages.photoStorage.storageArgs.data,
			).
			Return(
				tt.storages.photoStorage.storageReturn.photoKey,
				tt.storages.photoStorage.storageReturn.err,
			).
			Once()
		photoMockStorage.
			On("SaveKey",
				tt.storages.photoKeyStorage.storageArgs.ctx,
				tt.storages.photoKeyStorage.storageArgs.request,
			).
			Return(
				tt.storages.photoKeyStorage.storageReturn.err,
			).
			Once()
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
