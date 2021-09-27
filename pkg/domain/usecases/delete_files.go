package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
)

type DeleteFiles interface {
	Execute(ctx context.Context, deleteFiles DeleteFilesParam) ([]dto.FilePublic, error)
}

type DeleteFilesParam struct {
	Files  dto.FileList
	UserID uint64
}
