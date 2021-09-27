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
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"reflect"
	"testing"
)

func TestNewListFilesById(t *testing.T) {
	mockRepository := mocks.FilesRepository{}
	type args struct {
		filesRepository repositories.FilesRepository
	}
	tests := []struct {
		name string
		args args
		want usecases.ListFilesById
	}{
		{
			name: "it should return a file service",
			args: args{
				filesRepository: &mockRepository,
			},
			want: NewListFilesById(&mockRepository),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListFilesById(tt.args.filesRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListFilesById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listFilesById_Execute(t *testing.T) {
	userIdParam := 0
	isPublicTrue := true
	isPublicFalse := false
	userID := uint64(userIdParam)

	type fields struct {
		filesRepository repositories.FilesRepository
	}
	type args struct {
		data usecases.ListFilesByIdParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.FilePublic
		wantErr bool
	}{
		{
			name: "it should return a list of files",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("ListById", mock.Anything, []string{"123e4567-e89b-12d3-a456-426655440000", "223e4567-e89b-12d3-a456-426655440000"}, &userID).Return(FilesTest(userID, isPublicTrue), nil)
					return &repoMock
				}(),
			},
			args: args{
				data: usecases.ListFilesByIdParam{
					UserID: &userID,
					Files:  []string{"123e4567-e89b-12d3-a456-426655440000", "223e4567-e89b-12d3-a456-426655440000"},
				},
			},
			want:    entities.AddUrlToFilesResponse(FilesTest(userID, isPublicTrue)),
			wantErr: false,
		},
		{
			name: "it should return a list of private files",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("ListById", mock.Anything, []string{"123e4567-e89b-12d3-a456-426655440000"}, &userID).Return(FilesTest(userID, isPublicFalse), nil)
					return &repoMock
				}(),
			},
			args: args{
				data: usecases.ListFilesByIdParam{
					UserID: &userID,
					Files:  []string{"123e4567-e89b-12d3-a456-426655440000"},
				},
			},
			want:    entities.AddUrlToFilesResponse(FilesTest(userID, isPublicFalse)),
			wantErr: false,
		},
		{
			name: "it should return error invalid file fileid",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("ListById", mock.Anything, []string{"idtest"}, &userID).Return(nil, errors.New(""))
					return &repoMock
				}(),
			},
			args: args{
				data: usecases.ListFilesByIdParam{
					UserID: &userID,
					Files:  []string{"idtest"},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := listFilesById{
				filesRepository: tt.fields.filesRepository,
			}
			got, err := i.Execute(context.WithValue(context.Background(), helpers.ContextUserKey, uint64(1)), tt.args.data)
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
