package services

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
)

type listFilesById struct {
	filesRepository repositories.FilesRepository
}

func NewListFilesById(filesRepository repositories.FilesRepository) usecases.ListFilesById {
	return listFilesById{filesRepository: filesRepository}
}

// Execute list multiple files by multiples ids
func (i listFilesById) Execute(ctx context.Context, data usecases.ListFilesByIdParam) ([]dto.FilePublic, error) {
	files, err := i.filesRepository.ListById(ctx, data.Files, data.UserID)
	if err != nil {
		return nil, err
	}
	filesResponse := entities.AddUrlToFilesResponse(files)
	return filesResponse, nil
}
