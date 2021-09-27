package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
)

type FindFilesById interface {
	Execute(ctx context.Context, data FindFileByIdParam) (*dto.FilePublic, error)
}

type FindFileByIdParam struct {
	FileID string
	UserID uint64
}
