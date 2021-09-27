package services

import (
	"errors"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/dto"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
)

var (
	UserAlreadyExists              = errors.New("user already exists")
	UserAlreadyExistsWithThisEmail = errors.New("user already exists with this email")
	ErrPasswordIsInvalid           = errors.New("password is invalid")
)

type UsersService interface {
	FindById(userID int) (*dto.User, error)
	FindByEmail(email string) (*entities.User, error)
	Create(user entities.User) (*dto.User, error)
	Update(user entities.User) (*dto.User, error)
	ListAll() (*dto.Users, error)
	Login(user entities.User) (*helpers.TokenResponse, error)
}

type usersService struct {
	usersRepository repositories.UsersRepository
}

func NewUsersService(usersRepository repositories.UsersRepository) UsersService {
	return usersService{usersRepository: usersRepository}
}

func (u usersService) FindById(userID int) (*dto.User, error) {
	user, err := u.usersRepository.FindById(userID)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}

func (u usersService) FindByEmail(email string) (*entities.User, error) {
	user, err := u.usersRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u usersService) ListAll() (*dto.Users, error) {
	users, err := u.usersRepository.ListAll()
	if err != nil {
		return nil, err
	}

	return &dto.Users{Users: users}, nil
}

func (u usersService) Create(user entities.User) (*dto.User, error) {

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	userExists, err := u.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, UserAlreadyExists
	}

	err = u.usersRepository.Store(&user)
	if err != nil {
		return nil, err
	}

	userResponse := dto.User{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return &userResponse, err
}

func (u usersService) Update(user entities.User) (*dto.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	userLogged, err := u.FindById(int(user.UserID))
	if err != nil {
		return nil, err
	}

	userExists, err := u.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, UserAlreadyExistsWithThisEmail
	}

	inputUpdate := entities.User{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: userLogged.CreatedAt,
	}
	err = u.usersRepository.UpdateUser(&inputUpdate)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		UserID:    inputUpdate.UserID,
		Name:      inputUpdate.Name,
		Email:     inputUpdate.Email,
		CreatedAt: inputUpdate.CreatedAt,
		UpdatedAt: inputUpdate.UpdatedAt,
		DeletedAt: inputUpdate.DeletedAt,
	}, err
}

func (u usersService) Login(user entities.User) (*helpers.TokenResponse, error) {
	err := user.ValidateLogin()
	if err != nil {
		return nil, err
	}

	userExists, err := u.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if userExists == nil {
		return nil, ErrInvalidUserOrPassword
	}

	passwordIsValid, err := helpers.CheckPassowordHash(user.Password, userExists.Password)
	if err != nil {
		if !passwordIsValid {
			return nil, ErrInvalidUserOrPassword
		}

		return nil, err
	}

	if !passwordIsValid {
		return nil, ErrPasswordIsInvalid
	}

	return helpers.CreateJWT(userExists)

}
