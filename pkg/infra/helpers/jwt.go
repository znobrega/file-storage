package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"time"
)

type TokenResponse struct {
	UserId     uint64    `json:"userId"`
	AcessToken string    `json:"acessToken"`
	ExpiryDate time.Time `json:"expiryDate"`
}

func CreateJWT(user *entities.User) (*TokenResponse, error) {
	expiryDate := time.Now().Add(time.Minute * 120)
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.UserID,
		"expiry":     expiryDate.Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		UserId:     user.UserID,
		AcessToken: token,
		ExpiryDate: expiryDate,
	}, nil
}
