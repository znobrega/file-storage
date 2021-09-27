package services

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/znobrega/file-storage/mocks"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"reflect"
	"testing"
	"time"
)

func TestNewListFiles(t *testing.T) {
	mockRepository := mocks.FilesRepository{}
	type args struct {
		filesRepository repositories.FilesRepository
	}
	tests := []struct {
		name string
		args args
		want usecases.ListFiles
	}{
		{
			name: "it should return a file service",
			args: args{
				filesRepository: &mockRepository,
			},
			want: NewListFiles(&mockRepository),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListFiles(tt.args.filesRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FilesTest(userId uint64, isPublic bool) []entities.File {
	return []entities.File{{
		FileID:    "d173faa6-5c5e-1809-0b63-0fb27e500e19",
		Name:      "testfile",
		Extension: ".png",
		UserID:    userId,
		IsPublic:  false,
		CreatedAt: time.Time{},
		DeletedAt: nil,
	}, {
		FileID:    "d173faa6-5c5e-1809-0b63-0fb27e500e17",
		Name:      "testfile",
		Extension: ".png",
		UserID:    userId,
		IsPublic:  false,
		CreatedAt: time.Time{},
		DeletedAt: nil,
	}}
}

func Test_listFiles_Execute(t *testing.T) {
	userIdParam := 0
	isPublicTrue := true
	isPublicFalse := false
	page := 1
	limit := 10
	userID := uint64(userIdParam)

	type fields struct {
		filesRepository repositories.FilesRepository
	}
	type args struct {
		listFiles usecases.ListFilesParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.FileResponse
		wantErr bool
	}{
		{
			name: "it should return a list of public files",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("List", mock.Anything, usecases.ListFilesParam{
						UserId:   func() *int { id := int(userID); return &id }(),
						IsPublic: &isPublicTrue,
						Page:     page,
						Limit:    limit,
					}).Return(FilesTest(userID, isPublicTrue), nil)
					return &repoMock
				}(),
			},
			args: args{
				listFiles: usecases.ListFilesParam{
					UserId:   func() *int { id := int(userID); return &id }(),
					IsPublic: &isPublicTrue,
					Page:     page,
					Limit:    limit,
				},
			},
			want: &dto.FileResponse{
				Files: entities.AddUrlToFilesResponse(FilesTest(userID, isPublicTrue)),
			},
			wantErr: false,
		},
		{
			name: "it should return a list of private files",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("List", mock.Anything, usecases.ListFilesParam{
						UserId:   func() *int { id := int(userID); return &id }(),
						IsPublic: &isPublicTrue,
						Page:     page,
						Limit:    limit,
					}).Return(FilesTest(userID, isPublicFalse), nil)
					return &repoMock
				}(),
			},
			args: args{
				listFiles: usecases.ListFilesParam{
					UserId:   func() *int { id := int(userID); return &id }(),
					IsPublic: &isPublicTrue,
					Page:     page,
					Limit:    limit,
				},
			},
			want: &dto.FileResponse{
				Files: entities.AddUrlToFilesResponse(FilesTest(userID, isPublicFalse)),
			},
			wantErr: false,
		},
		{
			name: "it should return a list of private files",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("List",mock.Anything, usecases.ListFilesParam{
						UserId:   func() *int { id := int(userID); return &id }(),
						IsPublic: &isPublicTrue,
						Page:     page,
						Limit:    limit,
					}).Return(nil, errors.New(""))
					return &repoMock
				}(),
			},
			args: args{
				listFiles: usecases.ListFilesParam{
					UserId:   func() *int { id := int(userID); return &id }(),
					IsPublic: &isPublicTrue,
					Page:     page,
					Limit:    limit,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return error user if must be valid",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					return &mocks.FilesRepository{}
				}(),
			},
			args: args{
				listFiles: usecases.ListFilesParam{
					UserId:   nil,
					IsPublic: &isPublicFalse,
					Page:     page,
					Limit:    limit,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := listFiles{
				filesRepository: tt.fields.filesRepository,
			}
			got, err := i.Execute(context.Background(), tt.args.listFiles)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
