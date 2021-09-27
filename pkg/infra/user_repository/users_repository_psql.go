package user_repository

import (
	"errors"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/dto"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) repositories.UsersRepository {
	return UserRepository{db: db}
}

func (u UserRepository) FindByEmail(email string) (*entities.User, error) {
	var dbRecord *entities.User
	err := u.db.Where("email = ?", email).First(&dbRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return dbRecord, err
}

func (u UserRepository) ListAll() ([]dto.User, error) {
	var dbRecords []dto.User
	return dbRecords, u.db.Find(&dbRecords).Error
}

func (u UserRepository) UpdateUser(user *entities.User) error {
	return u.db.Updates(&user).Error
}

func (u UserRepository) Store(user *entities.User) error {
	return u.db.Create(&user).Error
}

func (u UserRepository) FindById(userID int) (*entities.User, error) {
	var dbRecord *entities.User
	err := u.db.Where("user_id = ?", userID).First(&dbRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return dbRecord, err
}
