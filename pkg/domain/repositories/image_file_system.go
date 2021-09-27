package repositories

import (
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
)

type SaveFilesOnFileSystem interface {
	Execute(fileData usecases.SaveFilesParam) ([]entities.File, error)
}
