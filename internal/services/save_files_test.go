package services

import (
	"context"
	"errors"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/mocks"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"github.com/znobrega/file-storage/pkg/fileid"
	"mime/multipart"
	"reflect"
	"testing"
)

func TestNewSaveFiles(t *testing.T) {
	mockRepository := mocks.FilesRepository{}
	mockFileIdGenerator := mocks.GenerateId{}
	mockSaveFileOnFileSystem := mocks.SaveFilesOnFileSystem{}
	type args struct {
		filesRepository       repositories.FilesRepository
		saveFilesOnFileSystem repositories.SaveFilesOnFileSystem
		fileIdGenerator       fileid.GenerateId
	}
	tests := []struct {
		name string
		args args
		want usecases.SaveFiles
	}{
		{
			name: "it should return a file service",
			args: args{
				filesRepository:       &mockRepository,
				saveFilesOnFileSystem: &mockSaveFileOnFileSystem,
				fileIdGenerator:       &mockFileIdGenerator,
			},
			want: NewSaveFiles(&mockRepository, &mockSaveFileOnFileSystem, &mockFileIdGenerator),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSaveFiles(tt.args.filesRepository, tt.args.saveFilesOnFileSystem, tt.args.fileIdGenerator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSaveFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func MakeMockFilesRepository(ctx context.Context, userID uint64, isPublic bool, page, limit int, fileId, filename, fileExtension string, throwError bool) repositories.FilesRepository {
	var err error
	if throwError {
		err = errors.New("error")
	}
	repoMock := mocks.FilesRepository{}
	repoMock.On("List", ctx, usecases.ListFilesParam{
		UserId:   func() *int { id := int(userID); return &id }(),
		IsPublic: &isPublic,
		Page:     page,
		Limit:    limit,
	}).Return(FilesTest(userID, isPublic), err)

	repoMock.On("Store", ctx, MakeFilesUnderTest(userID, isPublic, fileId, filename, fileExtension)).Return(err)
	return &repoMock
}

func MakeFilesUnderTest(userId uint64, isPublic bool, fileId, filename, fileExtension string) []entities.File {
	return []entities.File{{
		FileID:    fileId,
		Name:      filename,
		Extension: fileExtension,
		UserID:    userId,
		IsPublic:  isPublic,
	}}
}

func MakeMockFilesSystemRepository(fileId string, userID uint64, isPublic bool, page, limit int, files []*multipart.FileHeader, filename, directory, fileExtension string, throwError bool) repositories.SaveFilesOnFileSystem {
	var err error
	if throwError {
		err = errors.New("error")
	}
	repoMock := mocks.SaveFilesOnFileSystem{}
	repoMock.On("Execute", usecases.SaveFilesParam{
		FileId:    fileId,
		Files:     files,
		Filename:  filename,
		UserID:    userID,
		IsPublic:  isPublic,
		Directory: directory,
	}).Return(MakeFilesUnderTest(userID, isPublic, fileId, filename, fileExtension), err)
	return &repoMock
}

func MakeFileUnderTest(filename string, size int) []*multipart.FileHeader {
	var files []*multipart.FileHeader
	files = append(files, &multipart.FileHeader{
		Filename: filename,
		Header:   nil,
		Size:     int64(size),
	})
	return files
}

func MakeMockFileIdGenerator(fileId string) fileid.GenerateId {
	fileIdGenerator := &mocks.GenerateId{}
	fileIdGenerator.On("Generate").Return(fileId, nil)
	return fileIdGenerator
}

func Test_saveFiles_Execute(t *testing.T) {

	configs.LoadConfig("../../resources")
	userIDUnderTest := uint64(0)
	isPublicTrue := true
	isPublicFalse := false
	page, limit := 1, 1
	filename := "filetest"
	fileSize := 1
	directory := "directorytest"
	fileId := "123123"
	fileExtension := ".png"

	type fields struct {
		filesDatabaseRepository         repositories.FilesRepository
		saveFilesOnFileSystemRepository repositories.SaveFilesOnFileSystem
		fileIdGenerator                 fileid.GenerateId
	}
	type args struct {
		ctx      context.Context
		fileData usecases.SaveFilesParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.FilePublic
		wantErr bool
	}{
		{
			name: "it should return error when save on database",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, true),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx:      context.Background(),
				fileData: usecases.SaveFilesParam{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return error when save on file system",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, false),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicFalse, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx:      context.Background(),
				fileData: usecases.SaveFilesParam{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return error when there is no files",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, false),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx:      context.Background(),
				fileData: usecases.SaveFilesParam{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return error when there is no filename",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, false),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx: context.Background(),
				fileData: usecases.SaveFilesParam{
					FileId:    "",
					Files:     MakeFileUnderTest("", 10),
					Filename:  "",
					UserID:    0,
					IsPublic:  false,
					Directory: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return error when file size is too large",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, false),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx: context.Background(),
				fileData: usecases.SaveFilesParam{
					FileId:    "",
					Files:     MakeFileUnderTest("", configs.Viper.GetInt("files.sizeLimit")+10),
					Filename:  "",
					UserID:    0,
					IsPublic:  false,
					Directory: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return error when directory is empty",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, false),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx: context.Background(),
				fileData: usecases.SaveFilesParam{
					FileId:    "",
					Files:     MakeFileUnderTest("", 10),
					Filename:  "test",
					UserID:    0,
					IsPublic:  false,
					Directory: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should sucessfully save a file",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, false),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx: context.Background(),
				fileData: usecases.SaveFilesParam{
					FileId:    fileId,
					Files:     MakeFileUnderTest(filename, fileSize),
					Filename:  filename,
					UserID:    0,
					IsPublic:  isPublicTrue,
					Directory: directory,
				},
			},
			want:    entities.AddUrlToFilesResponse(MakeFilesUnderTest(userIDUnderTest, isPublicTrue, fileId, filename, fileExtension)),
			wantErr: false,
		},
		{
			name: "it should occur a database error",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, true),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, false),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx: context.Background(),
				fileData: usecases.SaveFilesParam{
					FileId:    fileId,
					Files:     MakeFileUnderTest(filename, fileSize),
					Filename:  filename,
					UserID:    0,
					IsPublic:  isPublicTrue,
					Directory: directory,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should should occur a file system error",
			fields: fields{
				filesDatabaseRepository:         MakeMockFilesRepository(context.Background(), userIDUnderTest, isPublicTrue, page, limit, fileId, filename, fileExtension, false),
				saveFilesOnFileSystemRepository: MakeMockFilesSystemRepository(fileId, userIDUnderTest, isPublicTrue, page, limit, MakeFileUnderTest(filename, fileSize), filename, directory, fileExtension, true),
				fileIdGenerator:                 MakeMockFileIdGenerator(fileId),
			},
			args: args{
				ctx: context.Background(),
				fileData: usecases.SaveFilesParam{
					FileId:    fileId,
					Files:     MakeFileUnderTest(filename, fileSize),
					Filename:  filename,
					UserID:    0,
					IsPublic:  isPublicTrue,
					Directory: directory,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := saveFiles{
				filesDatabaseRepository:         tt.fields.filesDatabaseRepository,
				saveFilesOnFileSystemRepository: tt.fields.saveFilesOnFileSystemRepository,
				fileIdGenerator:                 tt.fields.fileIdGenerator,
			}
			got, err := i.Execute(tt.args.ctx, tt.args.fileData)
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
