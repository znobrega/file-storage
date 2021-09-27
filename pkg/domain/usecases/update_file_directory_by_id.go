package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
)

type UpdateFileDirectoryById interface {
	Execute(ctx context.Context, UpdateFileDirectoryById UpdateFileDirectoryByIdParam) (*dto.FileResponse, error)
}

type UpdateFileDirectoryByIdParam struct {
	FileID       string `json:"fileId"`
	UserID       uint64 `json:"userId"`
	NewDirectory string `json:"newDirectory"`
}
