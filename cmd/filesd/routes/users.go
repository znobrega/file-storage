package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/znobrega/file-storage/cmd/filesd/controllers"
	"github.com/znobrega/file-storage/cmd/filesd/middlewares"
	"github.com/znobrega/file-storage/internal/services"
)

type usersRouter struct {
	userService services.UsersService
}

func NewUsersRouter(userService services.UsersService) Router {
	return usersRouter{
		userService: userService,
	}
}

func (i usersRouter) Routes(router chi.Router) {
	usersController := controllers.NewUsersController(i.userService)
	router.Route("/users", func(r chi.Router) {
		r.Post("/", usersController.HandleUserCreation())
		r.Post("/login", usersController.HandleUserLogin())
		r.Get("/list", usersController.HandleUserList())
		r.Get("/findone", usersController.HandleFindOneUser())
		r.With(middlewares.Authentication).Put("/", usersController.HandleUserUpdate())
	})
}
