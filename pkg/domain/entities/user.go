package entities

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var (
	ErrNameIsRequired     = errors.New("name is required")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrEmailIsRequired    = errors.New("email is required")
)

type User struct {
	UserID    uint64 `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameIsRequired
	}

	if u.Password == "" {
		return ErrPasswordIsRequired
	}

	if u.Email == "" {
		return ErrEmailIsRequired
	}

	return nil
}

func (u *User) ValidateLogin() error {
	if u.Email == "" {
		return ErrEmailIsRequired
	}

	if u.Password == "" {
		return ErrPasswordIsRequired
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	if tx.Statement.Changed("Password") {
		tx.Statement.SetColumn("Password", u.Password)
	}
	return nil
}
