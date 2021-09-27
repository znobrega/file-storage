package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
)

type ListDirectories interface {
	Execute(ctx context.Context) (*dto.Directories, error)
}

type ListDirectoriesParam struct {
	UserID *uint64
}
