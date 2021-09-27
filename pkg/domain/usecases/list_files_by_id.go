package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
)

type ListFilesById interface {
	Execute(ctx context.Context, ListFilesById ListFilesByIdParam) ([]dto.FilePublic, error)
}

type ListFilesByIdParam struct {
	Files  []string
	UserID *uint64
}
