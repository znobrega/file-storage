package repositories

import (
	"context"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
)

type FilesRepository interface {
	Store(ctx context.Context, files []entities.File) error
	Delete(ctx context.Context, file entities.File) error
	List(ctx context.Context, data usecases.ListFilesParam) ([]entities.File, error)
	ListById(ctx context.Context, files []string, userID *uint64) ([]entities.File, error)
	ListByDirectory(ctx context.Context, directory string) ([]entities.File, error)
	ListDirectories(ctx context.Context) ([]entities.File, error)
	Update(ctx context.Context, newFile *entities.File) error
}
