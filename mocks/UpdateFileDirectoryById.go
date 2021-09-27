// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	dto "github.com/znobrega/file-storage/pkg/dto"

	usecases "github.com/znobrega/file-storage/pkg/domain/usecases"
)

// UpdateFileDirectoryById is an autogenerated mock type for the UpdateFileDirectoryById type
type UpdateFileDirectoryById struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, UpdateFileDirectoryById
func (_m *UpdateFileDirectoryById) Execute(ctx context.Context, UpdateFileDirectoryById usecases.UpdateFileDirectoryByIdParam) (*dto.FileResponse, error) {
	ret := _m.Called(ctx, UpdateFileDirectoryById)

	var r0 *dto.FileResponse
	if rf, ok := ret.Get(0).(func(context.Context, usecases.UpdateFileDirectoryByIdParam) *dto.FileResponse); ok {
		r0 = rf(ctx, UpdateFileDirectoryById)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.FileResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, usecases.UpdateFileDirectoryByIdParam) error); ok {
		r1 = rf(ctx, UpdateFileDirectoryById)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}