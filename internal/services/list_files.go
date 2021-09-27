package services

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
)

type listFiles struct {
	filesRepository repositories.FilesRepository
}

func NewListFiles(filesRepository repositories.FilesRepository) usecases.ListFiles {
	return listFiles{filesRepository: filesRepository}
}

// Execute list files
func (i listFiles) Execute(ctx context.Context, listFiles usecases.ListFilesParam) (*dto.FileResponse, error) {
	isPrivateAndHasUserId := listFiles.IsPublic != nil && !*listFiles.IsPublic && listFiles.UserId == nil
	if isPrivateAndHasUserId {
		return nil, ErrUserIDMustBeValid
	}

	files, err := i.filesRepository.List(ctx, listFiles)
	if err != nil {
		return nil, err
	}

	filesResponse := entities.AddUrlToFilesResponse(files)

	return &dto.FileResponse{
		Files: filesResponse,
	}, err
}
