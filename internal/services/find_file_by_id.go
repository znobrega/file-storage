package services

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
)

type findFilesById struct {
	filesRepository repositories.FilesRepository
}

func NewFindFilesById(filesRepository repositories.FilesRepository) usecases.FindFilesById {
	return findFilesById{filesRepository: filesRepository}
}

// Execute find file by id
func (i findFilesById) Execute(ctx context.Context, data usecases.FindFileByIdParam) (*dto.FilePublic, error) {
	files, err := i.filesRepository.ListById(ctx, []string{data.FileID}, &data.UserID)
	if err != nil {
		return nil, err
	}
	// TODO THIS MUST BE A METHOD
	fileResponse := entities.AddUrlToFilesResponse(files)[0]
	return &fileResponse, nil
}
