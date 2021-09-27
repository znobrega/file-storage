package services

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"os"
)

type updateFileDirectoryById struct {
	filesRepository repositories.FilesRepository
}

func NewUpdateFileDirectoryById(filesRepository repositories.FilesRepository) usecases.UpdateFileDirectoryById {
	return updateFileDirectoryById{filesRepository: filesRepository}
}

func (i updateFileDirectoryById) Execute(ctx context.Context, data usecases.UpdateFileDirectoryByIdParam) (*dto.FileResponse, error) {
	files, err := i.filesRepository.ListById(ctx, []string{data.FileID}, &data.UserID)
	if err != nil {
		return nil, err
	}
	file := &files[0]

	oldPath := file.GetFullPath()

	file.Directory = data.NewDirectory

	file.SetPath()
	newPath := file.GetFullPath()
	newDir := file.GetDir()

	err = i.filesRepository.Update(ctx, file)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(newDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	err = os.Rename(oldPath, newPath)
	if err != nil {
		return nil, err
	}

	filesResponse := entities.AddUrlToFilesResponse(files)
	return &dto.FileResponse{Files: filesResponse}, nil
}
