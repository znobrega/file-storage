package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
	"mime/multipart"
)

type UpdateFileBlob interface {
	Execute(ctx context.Context, UpdateFileBlob UpdateFileBlobParam) (*dto.FileResponse, error)
}

type UpdateFileBlobParam struct {
	FileID string                  `json:"fileId"`
	Files  []*multipart.FileHeader `json:"Files"`
}
