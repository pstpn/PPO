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

func Test_documentServiceImpl_CreateDocument(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.CreateDocumentRequest
	}

	type storages struct {
		documentStorage struct {
			storageArgs   args
			storageReturn struct {
				document *model.Document
				err      error
			}
		}
	}

	documentMockStorage := mocks.NewDocumentStorage(t)
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		want    *model.Document
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect info card ID",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreateDocumentRequest{
					SerialNumber: "123",
					InfoCardID:   -1,
					DocumentType: 1,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						document *model.Document
						err      error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreateDocumentRequest{
							SerialNumber: "123",
							InfoCardID:   -1,
							DocumentType: 1,
						},
					},
					storageReturn: struct {
						document *model.Document
						err      error
					}{
						document: nil,
						err:      fmt.Errorf("incorrect infoardID"),
					},
				},
			},
		},
		{
			name: "success",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreateDocumentRequest{
					SerialNumber: "123",
					InfoCardID:   1,
					DocumentType: 1,
				},
			},
			want: &model.Document{
				ID:         model.ToDocumentID(1),
				InfoCardID: model.ToInfoCardID(1),
				Type:       model.ToDocumentTypeFromInt(1),
			},
			wantErr: false,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						document *model.Document
						err      error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreateDocumentRequest{
							SerialNumber: "123",
							InfoCardID:   1,
							DocumentType: 1,
						},
					},
					storageReturn: struct {
						document *model.Document
						err      error
					}{
						document: &model.Document{
							ID:         model.ToDocumentID(1),
							InfoCardID: model.ToInfoCardID(1),
							Type:       model.ToDocumentTypeFromInt(1),
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		documentMockStorage.
			On("Create",
				tt.storages.documentStorage.storageArgs.ctx,
				tt.storages.documentStorage.storageArgs.request,
			).
			Return(
				tt.storages.documentStorage.storageReturn.document,
				tt.storages.documentStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.CreateDocument(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("documentServiceImpl.CreateDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("documentServiceImpl.CreateDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_documentServiceImpl_GetDocument(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.GetDocumentByIDRequest
	}

	type storages struct {
		documentStorage struct {
			storageArgs   args
			storageReturn struct {
				document *model.Document
				err      error
			}
		}
	}

	documentMockStorage := mocks.NewDocumentStorage(t)
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		want    *model.Document
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect document ID",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetDocumentByIDRequest{DocumentID: -1},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						document *model.Document
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetDocumentByIDRequest{DocumentID: -1},
					},
					storageReturn: struct {
						document *model.Document
						err      error
					}{
						document: nil,
						err:      fmt.Errorf("incorrect documentID"),
					},
				},
			},
		},
		{
			name: "success",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetDocumentByIDRequest{DocumentID: 1},
			},
			want: &model.Document{
				ID:         model.ToDocumentID(1),
				InfoCardID: model.ToInfoCardID(1),
				Type:       model.ToDocumentTypeFromInt(1),
			},
			wantErr: false,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						document *model.Document
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetDocumentByIDRequest{DocumentID: 1},
					},
					storageReturn: struct {
						document *model.Document
						err      error
					}{
						document: &model.Document{
							ID:         model.ToDocumentID(1),
							InfoCardID: model.ToInfoCardID(1),
							Type:       model.ToDocumentTypeFromInt(1),
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		documentMockStorage.
			On("GetByID",
				tt.storages.documentStorage.storageArgs.ctx,
				tt.storages.documentStorage.storageArgs.request,
			).
			Return(
				tt.storages.documentStorage.storageReturn.document,
				tt.storages.documentStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetDocument(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("documentServiceImpl.GetDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("documentServiceImpl.GetDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_documentServiceImpl_ListEmployeeDocuments(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.ListEmployeeDocumentsRequest
	}

	type storages struct {
		documentStorage struct {
			storageArgs   args
			storageReturn struct {
				documents []*model.Document
				err       error
			}
		}
	}

	documentMockStorage := mocks.NewDocumentStorage(t)
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		want    []*model.Document
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect employee ID",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListEmployeeDocumentsRequest{EmployeeID: -1},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						documents []*model.Document
						err       error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListEmployeeDocumentsRequest{EmployeeID: -1},
					},
					storageReturn: struct {
						documents []*model.Document
						err       error
					}{
						documents: nil,
						err:       fmt.Errorf("incorrect employeeID"),
					},
				},
			},
		},
		{
			name: "success",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListEmployeeDocumentsRequest{EmployeeID: 1},
			},
			want: []*model.Document{
				{
					ID:         model.ToDocumentID(1),
					InfoCardID: model.ToInfoCardID(1),
					Type:       model.ToDocumentTypeFromInt(1),
				},
			},
			wantErr: false,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						documents []*model.Document
						err       error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListEmployeeDocumentsRequest{EmployeeID: 1},
					},
					storageReturn: struct {
						documents []*model.Document
						err       error
					}{
						documents: []*model.Document{
							{
								ID:         model.ToDocumentID(1),
								InfoCardID: model.ToInfoCardID(1),
								Type:       model.ToDocumentTypeFromInt(1),
							},
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		documentMockStorage.
			On("List",
				tt.storages.documentStorage.storageArgs.ctx,
				tt.storages.documentStorage.storageArgs.request,
			).
			Return(
				tt.storages.documentStorage.storageReturn.documents,
				tt.storages.documentStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.ListEmployeeDocuments(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("documentServiceImpl.ListEmployeeDocuments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("documentServiceImpl.ListEmployeeDocuments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_documentServiceImpl_DeleteDocument(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.DeleteDocumentRequest
	}

	type storages struct {
		documentStorage struct {
			storageArgs   args
			storageReturn struct {
				err error
			}
		}
	}

	documentMockStorage := mocks.NewDocumentStorage(t)
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect document ID",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteDocumentRequest{DocumentID: -1},
			},
			wantErr: true,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteDocumentRequest{DocumentID: -1},
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect documentID"),
					},
				},
			},
		},
		{
			name: "success",
			d: &documentServiceImpl{
				logger:          NewMockLogger(),
				documentStorage: documentMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteDocumentRequest{DocumentID: 1},
			},
			wantErr: false,

			storages: storages{
				documentStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteDocumentRequest{DocumentID: 1},
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
		documentMockStorage.
			On("Delete",
				tt.storages.documentStorage.storageArgs.ctx,
				tt.storages.documentStorage.storageArgs.request,
			).
			Return(
				tt.storages.documentStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.DeleteDocument(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("documentServiceImpl.DeleteDocument() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
