// Code generated by mockery v2.25.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "course/internal/service/dto"

	mock "github.com/stretchr/testify/mock"

	model "course/internal/model"
)

// DocumentStorage is an autogenerated mock type for the DocumentStorage type
type DocumentStorage struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) Create(ctx context.Context, request *dto.CreateDocumentRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateDocumentRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) Delete(ctx context.Context, request *dto.DeleteDocumentRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.DeleteDocumentRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) GetByID(ctx context.Context, request *dto.GetDocumentRequest) (*model.Document, error) {
	ret := _m.Called(ctx, request)

	var r0 *model.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentRequest) (*model.Document, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentRequest) *model.Document); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetDocumentRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) List(ctx context.Context, request *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error) {
	ret := _m.Called(ctx, request)

	var r0 []*model.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListEmployeeDocumentsRequest) ([]*model.Document, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListEmployeeDocumentsRequest) []*model.Document); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListEmployeeDocumentsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDocumentStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewDocumentStorage creates a new instance of DocumentStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDocumentStorage(t mockConstructorTestingTNewDocumentStorage) *DocumentStorage {
	mock := &DocumentStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
