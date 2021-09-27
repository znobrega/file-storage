// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	dto "github.com/znobrega/file-storage/pkg/dto"

	usecases "github.com/znobrega/file-storage/pkg/domain/usecases"
)

// SaveFiles is an autogenerated mock type for the SaveFiles type
type SaveFiles struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, saveFiles
func (_m *SaveFiles) Execute(ctx context.Context, saveFiles usecases.SaveFilesParam) ([]dto.FilePublic, error) {
	ret := _m.Called(ctx, saveFiles)

	var r0 []dto.FilePublic
	if rf, ok := ret.Get(0).(func(context.Context, usecases.SaveFilesParam) []dto.FilePublic); ok {
		r0 = rf(ctx, saveFiles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.FilePublic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, usecases.SaveFilesParam) error); ok {
		r1 = rf(ctx, saveFiles)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
