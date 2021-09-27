package services

import (
	"context"
	"fmt"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"github.com/znobrega/file-storage/pkg/fileid"
)

type saveFiles struct {
	filesDatabaseRepository         repositories.FilesRepository
	saveFilesOnFileSystemRepository repositories.SaveFilesOnFileSystem
	fileIdGenerator                 fileid.GenerateId
}

func NewSaveFiles(filesRepository repositories.FilesRepository, saveFilesOnFileSystem repositories.SaveFilesOnFileSystem, fileIdGenerator fileid.GenerateId) usecases.SaveFiles {
	return saveFiles{
		filesDatabaseRepository:         filesRepository,
		saveFilesOnFileSystemRepository: saveFilesOnFileSystem,
		fileIdGenerator:                 fileIdGenerator,
	}
}

// Execute saves a file
func (i saveFiles) Execute(ctx context.Context, fileData usecases.SaveFilesParam) ([]dto.FilePublic, error) {
	err := i.validateData(fileData)
	if err != nil {
		return nil, err
	}

	fileId, err := i.fileIdGenerator.Generate()
	fileData.FileId = fileId

	filesSavedOnFileSystem, err := i.saveFilesOnFileSystemRepository.Execute(fileData)
	if err != nil {
		return nil, err
	}

	err = i.filesDatabaseRepository.Store(ctx, filesSavedOnFileSystem)
	if err != nil {
		return nil, err
	}

	filesResponse := entities.AddUrlToFilesResponse(filesSavedOnFileSystem)

	return filesResponse, nil
}

func (i saveFiles) validateData(fileData usecases.SaveFilesParam) error {
	isFileMissing := fileData.Files == nil
	if isFileMissing {
		return fmt.Errorf("there is no files")
	}

	isInvalidUserId := fileData.UserID < 0
	if isInvalidUserId {
		return ErrUserIDMustBePositive
	}

	isMultipleFileRequest := len(fileData.Files) > 1
	if isMultipleFileRequest {
		return fmt.Errorf("only 1 file per request is allowed")
	}

	isFileSizeExceeded := fileData.Files[0].Size > configs.Viper.GetInt64("files.sizeLimit")
	if isFileSizeExceeded {
		return fmt.Errorf("file is too large")
	}

	if fileData.Filename == "" {
		return fmt.Errorf("filename is required")
	}

	if fileData.Directory == "" {
		return fmt.Errorf("directory is required")
	}

	return nil
}
