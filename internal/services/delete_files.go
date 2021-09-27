package services

import (
	"context"
	"errors"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"log"
	"os"
)

type deleteFiles struct {
	filesRepository repositories.FilesRepository
}

func NewDeleteFiles(filesRepository repositories.FilesRepository) usecases.DeleteFiles {
	return deleteFiles{filesRepository: filesRepository}
}

// Execute executes file deletions given the id list received
func (i deleteFiles) Execute(ctx context.Context, data usecases.DeleteFilesParam) ([]dto.FilePublic, error) {
	files, err := i.filesRepository.ListById(ctx, data.Files.Files, &data.UserID)
	if err != nil {
		return nil, err
	}

	fileDoestExistOrUserInvalidPermission := len(files) != len(data.Files.Files)
	if fileDoestExistOrUserInvalidPermission {
		return nil, errors.New("files doesnt exist or user doesn't have permission")
	}

	filesDeleted := 0
	for _, file := range files {
		err = i.filesRepository.Delete(ctx, file)
		if err != nil {
			return nil, err
		}

		err = os.Remove(file.GetFullPath())
		if err != nil {
			return nil, err
		}
		filesDeleted++
	}

	log.Println("deletions asked:", len(data.Files.Files), "deletions completed: ", filesDeleted)

	return nil, nil
}
