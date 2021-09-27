package file_repository

import (
	"context"
	"fmt"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFilesRepository(db *gorm.DB) repositories.FilesRepository {
	return FileRepository{db: db}
}

func (i FileRepository) ListByDirectory(ctx context.Context, directory string) ([]entities.File, error) {
	var dbRecords []entities.File
	userId := helpers.GetUserIdFromContext(ctx)
	err := i.db.Where("user_id = ?", userId).Where("directory like ?", fmt.Sprintf("%%%s%%", directory)).Find(&dbRecords).Error
	return dbRecords, err
}

func (i FileRepository) ListDirectories(ctx context.Context) ([]entities.File, error) {
	var dbRecords []entities.File
	userId := helpers.GetUserIdFromContext(ctx)
	err := i.db.Where("user_id = ?", userId).Order("directory").Find(&dbRecords).Error
	return dbRecords, err
}

func (i FileRepository) Update(ctx context.Context, newFile *entities.File) error {
	return i.db.Model(newFile).Where("file_id", newFile.FileID).Updates(*newFile).Error
}

func (i FileRepository) ListById(ctx context.Context, files []string, userID *uint64) ([]entities.File, error) {
	var dbRecords []entities.File
	dbCopy := i.db
	if userID != nil {
		dbCopy = i.db.Where("user_id = ?", *userID)
	}
	err := dbCopy.Where("file_id in ?", files).Find(&dbRecords).Error
	if recordNotFound := dbCopy.RowsAffected == 0; recordNotFound {
		return nil, fmt.Errorf("user record not found")
	}
	return dbRecords, err
}

func (i FileRepository) Store(ctx context.Context, files []entities.File) error {
	return i.db.Create(files).Error
}

func (i FileRepository) Delete(ctx context.Context, file entities.File) error {
	return i.db.Where("file_id = ?", file.FileID).Delete(&file).Error
}

func (i FileRepository) ListPublicFiles(ctx context.Context, page, limit int) ([]entities.File, error) {
	var dbRecords []entities.File
	offset := helpers.CalculateOffset(page, limit)
	err := i.db.Where("is_public = ?", true).Limit(limit).Offset(offset).Find(&dbRecords).Error
	return dbRecords, err
}

func (i FileRepository) List(ctx context.Context, data usecases.ListFilesParam) ([]entities.File, error) {
	dbCopy := i.db
	var dbRecords []entities.File

	if data.UserId != nil {
		dbCopy = dbCopy.Where("user_id = ?", *data.UserId)
	}

	offset := helpers.CalculateOffset(data.Page, data.Limit)
	err := dbCopy.Where("is_public = ?", data.IsPublic).Limit(data.Limit).Offset(offset).Find(&dbRecords).Error
	return dbRecords, err
}
