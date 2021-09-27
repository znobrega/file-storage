package dto

import "time"

type User struct {
	UserID    uint64     `json:"userId" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type UserRequest struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}
