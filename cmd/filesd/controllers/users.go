package controllers

import (
	"encoding/json"
	"github.com/znobrega/file-storage/internal/services"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/dto"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UsersController interface {
	HandleUserCreation() http.HandlerFunc
	HandleUserList() http.HandlerFunc
	HandleUserLogin() http.HandlerFunc
	HandleUserUpdate() http.HandlerFunc
	HandleFindOneUser() http.HandlerFunc
}

type usersController struct {
	usersService services.UsersService
}

func NewUsersController(usersService services.UsersService) UsersController {
	return usersController{usersService: usersService}
}

// @Title user create
// @Tags Users
// @Summary Creates a new user
// @Description Creates a new user
// @Param  content body    dto.UserRequest  true "Object for creating the user"
// @Success 200 {object} dto.User
// @Accept json
// @Router /users/ [post]
func (u usersController) HandleUserCreation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestJson dto.UserRequest
		body, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &requestJson)
		userResponse, err := u.usersService.Create(entities.User{
			Name:     requestJson.Name,
			Email:    requestJson.Email,
			Password: requestJson.Password,
		})
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, userResponse)
	}
}

// @Title user create
// @Tags Users
// @Security userIdAuthentication
// @Summary Updates a user
// @Description Updates a user
// @Param  content body    dto.UserRequest  true "Object for update the user"
// @Success 200 {object} dto.User
// @Accept json
// @Router /users/ [put]
func (u usersController) HandleUserUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestJson dto.UserRequest
		body, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &requestJson)
		userResponse, err := u.usersService.Update(entities.User{
			UserID:   helpers.GetUserIdFromContext(r.Context()),
			Name:     requestJson.Name,
			Email:    requestJson.Email,
			Password: requestJson.Password,
		})
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, userResponse)
	}
}

// @Title user login
// @Tags Users
// @Summary Sign in a user
// @Description Sign in a user
// @Param  content body    dto.LoginRequest  true "Object for sign in the user"
// @Success 200 {object} helpers.TokenResponse
// @Accept json
// @Router /users/login [post]
func (u usersController) HandleUserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestJson dto.LoginRequest
		body, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &requestJson)
		tokenResponse, err := u.usersService.Login(entities.User{
			Password: requestJson.Password,
			Email:    requestJson.Email,
		})
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, tokenResponse)
	}
}

// @Title list users
// @Tags Users
// @Summary list users
// @Description list users
// @Success 200 {object} dto.Users
// @Accept json
// @Router /users/list [get]
func (u usersController) HandleUserList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usersResponse, err := u.usersService.ListAll()
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, usersResponse)
	}
}

// @Title Find user
// @Tags Users
// @Summary Find specific user
// @Description Find specific user
// @Param   user_id 				query    string   false "Attribute to get specific user"
// @Success 200 {object} dto.User
// @Accept json
// @Router /users/findone [get]
func (u usersController) HandleFindOneUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		userID, exists := query["user_id"]
		if !exists {
			http.Error(w, "user_id param page is required", 400)
			return
		}

		id, _ := strconv.Atoi(userID[0])

		usersResponse, err := u.usersService.FindById(id)
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, usersResponse)
	}
}
