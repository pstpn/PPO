package service

import (
	"context"
	"course/internal/model"
	"reflect"
	"testing"
)

func Test_documentServiceImpl_CreateDocument(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *CreateDocumentRequest
	}
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.CreateDocument(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("documentServiceImpl.CreateDocument() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_documentServiceImpl_GetDocument(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *GetDocumentRequest
	}
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		want    *model.Document
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
	type args struct {
		ctx     context.Context
		request *ListEmployeeDocumentsRequest
	}
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		want    []*model.Document
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
	type args struct {
		ctx     context.Context
		request *DeleteDocumentRequest
	}
	tests := []struct {
		name    string
		d       *documentServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.DeleteDocument(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("documentServiceImpl.DeleteDocument() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
