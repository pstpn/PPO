package service

import (
	"context"
	"course/internal/storage/mocks"
	"fmt"
	"reflect"
	"testing"

	"course/internal/model"
	"course/internal/service/dto"
)

func Test_fieldServiceImpl_CreateCardField(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.CreateDocumentFieldRequest
	}

	type storages struct {
		fieldStorage struct {
			storageArgs   args
			storageReturn struct {
				field *model.Field
				err   error
			}
		}
	}

	fieldMockStorage := mocks.NewFieldStorage(t)
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		want    *model.Field
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect document ID",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreateDocumentFieldRequest{
					DocumentID: -1,
					Type:       1,
					Value:      "ok",
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						field *model.Field
						err   error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreateDocumentFieldRequest{
							DocumentID: -1,
							Type:       1,
							Value:      "ok",
						},
					},
					storageReturn: struct {
						field *model.Field
						err   error
					}{
						field: nil,
						err:   fmt.Errorf("incorrect documentID"),
					},
				},
			},
		},
		{
			name: "success",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreateDocumentFieldRequest{
					DocumentID: 1,
					Type:       1,
					Value:      "ok",
				},
			},
			want: &model.Field{
				ID:         model.ToFieldID(1),
				DocumentID: model.ToDocumentID(1),
				Type:       model.ToFieldTypeFromInt(1),
				Value:      "ok",
			},
			wantErr: false,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						field *model.Field
						err   error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreateDocumentFieldRequest{
							DocumentID: 1,
							Type:       1,
							Value:      "ok",
						},
					},
					storageReturn: struct {
						field *model.Field
						err   error
					}{
						field: &model.Field{
							ID:         model.ToFieldID(1),
							DocumentID: model.ToDocumentID(1),
							Type:       model.ToFieldTypeFromInt(1),
							Value:      "ok",
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		fieldMockStorage.
			On("Create",
				tt.storages.fieldStorage.storageArgs.ctx,
				tt.storages.fieldStorage.storageArgs.request,
			).
			Return(
				tt.storages.fieldStorage.storageReturn.field,
				tt.storages.fieldStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.CreateDocumentField(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.CreateDocumentField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldServiceImpl.CreateDocumentField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldServiceImpl_GetCardField(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.GetDocumentFieldRequest
	}

	type storages struct {
		fieldStorage struct {
			storageArgs   args
			storageReturn struct {
				field *model.Field
				err   error
			}
		}
	}

	fieldMockStorage := mocks.NewFieldStorage(t)
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		want    *model.Field
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect document ID",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.GetDocumentFieldRequest{
					DocumentID: -1,
					FieldType:  1,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						field *model.Field
						err   error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.GetDocumentFieldRequest{
							DocumentID: -1,
							FieldType:  1,
						},
					},
					storageReturn: struct {
						field *model.Field
						err   error
					}{
						field: nil,
						err:   fmt.Errorf("incorrect documentID"),
					},
				},
			},
		},
		{
			name: "success",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.GetDocumentFieldRequest{
					DocumentID: 1,
					FieldType:  1,
				},
			},
			want: &model.Field{
				ID:         model.ToFieldID(1),
				DocumentID: model.ToDocumentID(1),
				Type:       model.ToFieldTypeFromInt(1),
				Value:      "ok",
			},
			wantErr: false,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						field *model.Field
						err   error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.GetDocumentFieldRequest{
							DocumentID: 1,
							FieldType:  1,
						},
					},
					storageReturn: struct {
						field *model.Field
						err   error
					}{
						field: &model.Field{
							ID:         model.ToFieldID(1),
							DocumentID: model.ToDocumentID(1),
							Type:       model.ToFieldTypeFromInt(1),
							Value:      "ok",
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		fieldMockStorage.
			On("Get",
				tt.storages.fieldStorage.storageArgs.ctx,
				tt.storages.fieldStorage.storageArgs.request,
			).
			Return(
				tt.storages.fieldStorage.storageReturn.field,
				tt.storages.fieldStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetDocumentField(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.GetDocumentField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldServiceImpl.GetDocumentField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldServiceImpl_ListCardFields(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.ListDocumentFieldsRequest
	}

	type storages struct {
		fieldStorage struct {
			storageArgs   args
			storageReturn struct {
				fields []*model.Field
				err    error
			}
		}
	}

	fieldMockStorage := mocks.NewFieldStorage(t)
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		want    []*model.Field
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect document ID",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListDocumentFieldsRequest{DocumentID: -1},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						fields []*model.Field
						err    error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListDocumentFieldsRequest{DocumentID: -1},
					},
					storageReturn: struct {
						fields []*model.Field
						err    error
					}{
						fields: nil,
						err:    fmt.Errorf("incorrect documentID"),
					},
				},
			},
		},
		{
			name: "success",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListDocumentFieldsRequest{DocumentID: 1},
			},
			want: []*model.Field{
				{
					ID:         model.ToFieldID(1),
					DocumentID: model.ToDocumentID(1),
					Type:       model.ToFieldTypeFromInt(1),
					Value:      "ok",
				},
			},
			wantErr: false,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						fields []*model.Field
						err    error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListDocumentFieldsRequest{DocumentID: 1},
					},
					storageReturn: struct {
						fields []*model.Field
						err    error
					}{
						fields: []*model.Field{
							{
								ID:         model.ToFieldID(1),
								DocumentID: model.ToDocumentID(1),
								Type:       model.ToFieldTypeFromInt(1),
								Value:      "ok",
							},
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		fieldMockStorage.
			On("ListDocumentFields",
				tt.storages.fieldStorage.storageArgs.ctx,
				tt.storages.fieldStorage.storageArgs.request,
			).
			Return(
				tt.storages.fieldStorage.storageReturn.fields,
				tt.storages.fieldStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.ListDocumentFields(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.ListDocumentFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldServiceImpl.ListDocumentFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldServiceImpl_DeleteCardField(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.DeleteDocumentFieldRequest
	}

	type storages struct {
		fieldStorage struct {
			storageArgs   args
			storageReturn struct {
				err error
			}
		}
	}

	fieldMockStorage := mocks.NewFieldStorage(t)
	tests := []struct {
		name    string
		f       *fieldServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect field ID",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteDocumentFieldRequest{FieldID: -1},
			},
			wantErr: true,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteDocumentFieldRequest{FieldID: -1},
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect fieldID"),
					},
				},
			},
		},
		{
			name: "success",
			f: &fieldServiceImpl{
				logger:       NewMockLogger(),
				fieldStorage: fieldMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteDocumentFieldRequest{FieldID: 1},
			},
			wantErr: false,

			storages: storages{
				fieldStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteDocumentFieldRequest{FieldID: 1},
					},
					storageReturn: struct {
						err error
					}{
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		fieldMockStorage.
			On("Delete",
				tt.storages.fieldStorage.storageArgs.ctx,
				tt.storages.fieldStorage.storageArgs.request,
			).
			Return(
				tt.storages.fieldStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.DeleteDocumentField(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("fieldServiceImpl.DeleteDocumentField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
