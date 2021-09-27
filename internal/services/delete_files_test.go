package services

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/mocks"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewDeleteFiles(t *testing.T) {
	mockRepository := mocks.FilesRepository{}
	type args struct {
		filesRepository repositories.FilesRepository
	}
	tests := []struct {
		name string
		args args
		want usecases.DeleteFiles
	}{
		{
			name: "it should return a file service",
			args: args{
				filesRepository: &mockRepository,
			},
			want: NewDeleteFiles(&mockRepository),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeleteFiles(tt.args.filesRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FilesTestDelete(userId uint64, isPublic bool) []entities.File {
	return []entities.File{{
		FileID:    "d173faa6-5c5e-1809-0b63-0fb27e500e19",
		Name:      "testfile",
		Extension: ".png",
		UserID:    userId,
		IsPublic:  false,
		CreatedAt: time.Time{},
		DeletedAt: nil,
	}}
}

func Test_deleteFiles_Execute(t *testing.T) {
	configs.LoadConfig("../../resources")
	userIdParam := 0
	isPublicTrue := true
	userID := uint64(userIdParam)
	type fields struct {
		filesRepository repositories.FilesRepository
	}
	type args struct {
		data usecases.DeleteFilesParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.FilePublic
		wantErr bool
	}{
		{
			name: "it should successfully delete a file",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					file, err := os.Create(fmt.Sprintf("%s/%d/%s%s", "files", userID,  "testfile", ".png"))
					if err != nil {
						t.Errorf(err.Error())
					}
					defer file.Close()

					repoMock := mocks.FilesRepository{}
					repoMock.On("ListById", mock.Anything, []string{"d173faa6-5c5e-1809-0b63-0fb27e500e19"}, &userID).Return(FilesTestDelete(userID, isPublicTrue), nil)
					repoMock.On("Delete", mock.Anything, FilesTest(userID, isPublicTrue)[0]).Return(nil)
					return &repoMock
				}(),
			},
			args: args{
				data: usecases.DeleteFilesParam{
					Files: dto.FileList{
						Files: []string{"d173faa6-5c5e-1809-0b63-0fb27e500e19"},
					},
					UserID: userID,
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "it should occur an error files doesnt exist or user doesn't have permission",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("ListById", mock.Anything, []string{"d173faa6-5c5e-1809-0b63-0fb27e500e19", "d173faa6-5c5e-1809-0b63-0fb27e500e39", "d173faa6-5c5e-1809-0b63-0fb27e500e29"}, &userID).Return(FilesTest(userID, isPublicTrue), nil)
					repoMock.On("Delete", mock.Anything, FilesTest(userID, isPublicTrue)[0]).Return(nil)
					return &repoMock
				}(),
			},
			args: args{
				data: usecases.DeleteFilesParam{
					Files: dto.FileList{
						Files: []string{"d173faa6-5c5e-1809-0b63-0fb27e500e19", "d173faa6-5c5e-1809-0b63-0fb27e500e39", "d173faa6-5c5e-1809-0b63-0fb27e500e29"},
					},
					UserID: userID,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should occur an error file does not exist ",
			fields: fields{
				filesRepository: func() repositories.FilesRepository {
					repoMock := mocks.FilesRepository{}
					repoMock.On("ListById", mock.Anything, []string{"123e4567-e89b-12d3-a456-426655440000", "223e4567-e89b-12d3-a456-426655440000"}, &userID).Return([]entities.File{{
						FileID:    "d173faa6-5c5e-1809-0b63-0fb27e500e19",
						Name:      "testfile",
						Extension: ".png",
						UserID:    userID,
						IsPublic:  false,
						CreatedAt: time.Time{},
						DeletedAt: nil,
					}}, nil)
					repoMock.On("Delete", mock.Anything, FilesTest(userID, isPublicTrue)[0]).Return(nil)
					return &repoMock
				}(),
			},
			args: args{
				data: usecases.DeleteFilesParam{
					Files: dto.FileList{
						Files: []string{"123e4567-e89b-12d3-a456-426655440000", "223e4567-e89b-12d3-a456-426655440000"},
					},
					UserID: userID,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := deleteFiles{
				filesRepository: tt.fields.filesRepository,
			}
			got, err := i.Execute(context.Background(), tt.args.data)
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
