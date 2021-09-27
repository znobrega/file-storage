package services

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"os"
)

type updateFileBlob struct {
	filesRepository                 repositories.FilesRepository
	saveFilesOnFileSystemRepository repositories.SaveFilesOnFileSystem
}

func NewUpdateFileBlob(filesRepository repositories.FilesRepository, saveFilesOnFileSystemRepository repositories.SaveFilesOnFileSystem) usecases.UpdateFileBlob {
	return updateFileBlob{filesRepository: filesRepository, saveFilesOnFileSystemRepository: saveFilesOnFileSystemRepository}
}

func (i updateFileBlob) Execute(ctx context.Context, data usecases.UpdateFileBlobParam) (*dto.FileResponse, error) {
	userId := helpers.GetUserIdFromContext(ctx)
	files, err := i.filesRepository.ListById(ctx, []string{data.FileID}, &userId)
	if err != nil {
		return nil, err
	}
	file := files[0]

	filesSaved, err := i.saveFilesOnFileSystemRepository.Execute(usecases.SaveFilesParam{
		FileId:    file.FileID,
		Files:     data.Files,
		Filename:  file.Name,
		UserID:    userId,
		IsPublic:  file.IsPublic,
		Directory: file.Directory,
	})
	if err != nil {
		return nil, err
	}

	for _, fileSaved := range filesSaved {
		err := i.clearOldFile(file, fileSaved)
		if err != nil {
			return nil, err
		}

		err = i.filesRepository.Update(ctx, &fileSaved)
		if err != nil {
			return nil, err
		}
	}

	filesResponse := entities.AddUrlToFilesResponse(filesSaved)
	return &dto.FileResponse{
		Files: filesResponse,
	}, nil
}

func (i updateFileBlob) clearOldFile(file entities.File, fileSaved entities.File) error {
	isFileReplaced := file.Extension == fileSaved.Extension
	if isFileReplaced {
		return nil
	}

	err := os.Remove(file.GetFullPath())
	if err != nil {
		return err
	}
	return nil
}
