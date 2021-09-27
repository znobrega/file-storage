package services

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
)

type listFilesByDirectory struct {
	filesRepository repositories.FilesRepository
}

func NewListFilesByDirectory(filesRepository repositories.FilesRepository) usecases.ListFilesByDirectory {
	return listFilesByDirectory{filesRepository: filesRepository}
}

// Execute list files by directory
func (i listFilesByDirectory) Execute(ctx context.Context, data usecases.ListFilesByDirectoryParam) ([]dto.FilePublic, error) {
	files, err := i.filesRepository.ListByDirectory(ctx, data.Directory)
	if err != nil {
		return nil, err
	}

	filesResponse := entities.AddUrlToFilesResponse(files)

	return filesResponse, err
}
