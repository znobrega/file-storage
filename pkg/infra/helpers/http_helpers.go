package helpers

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"
)

var (
	ErrIsPublicRequired   = errors.New("body param isPublic is required")
	ErrFilenameIsRequired = errors.New("body param filename is required")
	ErrLimitRequired      = errors.New("query param limit is required")
	ErrPageRequired       = errors.New("query param page is required")
)

type HttpError struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
}

type HttpResponse struct {
	Data interface{}
}

func ReturnHttpError(w http.ResponseWriter, status int, err error) {
	log.Println(err)
	w.WriteHeader(status)
	WriteResponseAsJson(w, HttpError{
		Status:      status,
		Description: err.Error(),
	})
}

func WriteResponseAsJson(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(response)
}

func GetBooleanFromHeader(r *http.Request, key string) (bool, error) {
	value := r.Header.Get(key)
	if value == "" {
		return true, nil
	}
	boolean, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return boolean, nil
}

func GetIntegerFromHeader(r *http.Request, key string) (*int, error) {
	value := r.Header.Get(key)
	if value == "" {
		return nil, nil
	}
	integer, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &integer, nil
}

func GetUserIdFromContext(context context.Context) uint64 {
	return context.Value(ContextUserKey).(uint64)
}
