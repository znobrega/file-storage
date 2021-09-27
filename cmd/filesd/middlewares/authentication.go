package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

var (
	authorization           = "Authorization"
	ErrTokenIsRequired      = errors.New("token is required")
	ErrAuthorizationInvalid = errors.New("authorization invalid")
	ErrInvalidToken         = errors.New("token invalid")
)

func Authentication(httpHandler http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get(authorization)
		if authorizationHeader == "" {
			helpers.ReturnHttpError(w, 403, ErrTokenIsRequired)
			return
		}

		BearerAndToken := strings.Split(authorizationHeader, " ")
		if len(BearerAndToken) < 2 {
			helpers.ReturnHttpError(w, 403, ErrAuthorizationInvalid)
		}

		tokenString := BearerAndToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.secret")), nil
		})

		if err != nil {
			helpers.ReturnHttpError(w, 403, err)
			return
		}

		if !token.Valid {
			helpers.ReturnHttpError(w, 403, ErrInvalidToken)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helpers.ReturnHttpError(w, 403, ErrInvalidToken)
			return
		}

		userId, ok := claims["user_id"].(float64)
		if !ok {
			helpers.ReturnHttpError(w, 403, ErrInvalidToken)
			return
		}
		ctx := context.WithValue(r.Context(), helpers.ContextUserKey, uint64(userId))

		httpHandler.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)

}

func getTokenFromHeader(r *http.Request) string {
	authorizationHeader := r.Header.Get(authorization)
	if authorizationHeader == "" {
		return ""
	}
	return strings.Split(authorizationHeader, " ")[1]
}
