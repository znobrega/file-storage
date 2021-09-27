package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
)

type ListFiles interface {
	Execute(ctx context.Context, listFiles ListFilesParam) (*dto.FileResponse, error)
}

type ListFilesParam struct {
	UserId      *int
	IsPublic    *bool
	Page, Limit int
}
