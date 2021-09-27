package services

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
)

type listDirectories struct {
	filesRepository repositories.FilesRepository
}

func NewListDirectories(filesRepository repositories.FilesRepository) usecases.ListDirectories {
	return listDirectories{filesRepository: filesRepository}
}

// Execute list directories
func (i listDirectories) Execute(ctx context.Context) (*dto.Directories, error) {
	files, err := i.filesRepository.ListDirectories(ctx)
	if err != nil {
		return nil, err
	}

	var directories dto.Directories
	for _, file := range files {
		directories.Directories = append(directories.Directories, dto.Directory{Path: file.GetPublicFullPath()})
	}
	return &directories, err
}
