package repositories

import (
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/dto"
)

type UsersRepository interface {
	Store(users *entities.User) error
	FindById(userID int) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	ListAll() ([]dto.User, error)
	UpdateUser(user *entities.User) error
}
