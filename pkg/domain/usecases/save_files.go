package usecases

import (
	"context"
	"github.com/znobrega/file-storage/pkg/dto"
	"mime/multipart"
)

type SaveFiles interface {
	Execute(ctx context.Context, saveFiles SaveFilesParam) ([]dto.FilePublic, error)
}

type SaveFilesParam struct {
	FileId    string
	Files     []*multipart.FileHeader
	Filename  string
	UserID    uint64
	IsPublic  bool
	Directory string
}
