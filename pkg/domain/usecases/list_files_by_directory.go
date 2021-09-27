package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
)

type ListFilesByDirectory interface {
	Execute(ctx context.Context, data ListFilesByDirectoryParam) ([]dto.FilePublic, error)
}

type ListFilesByDirectoryParam struct {
	Directory string `json:"directory"`
}
