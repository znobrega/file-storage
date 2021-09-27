package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/znobrega/file-storage/cmd/filesd/controllers"
	"github.com/znobrega/file-storage/cmd/filesd/middlewares"
	"github.com/znobrega/file-storage/internal/container"
)

type filesRouter struct {
	dependencies container.Dependencies
}

func NewFilesRouter(dependencies container.Dependencies) Router {
	return filesRouter{
		dependencies: dependencies,
	}
}

func (i filesRouter) Routes(router chi.Router) {
	filesController := controllers.NewFileController(
		i.dependencies.ServicesContainer.SaveFilesService,
		i.dependencies.ServicesContainer.DeleteFilesService,
		i.dependencies.ServicesContainer.ListFilesService,
		i.dependencies.ServicesContainer.ListFilesByIdService,
		i.dependencies.ServicesContainer.FindFilesByIdService,
		i.dependencies.ServicesContainer.UpdateFileDirectoryByIdService,
		i.dependencies.ServicesContainer.ListDirectoriesService,
		i.dependencies.ServicesContainer.ListFilesByDirectoriesService,
		i.dependencies.ServicesContainer.UpdateFileBlobService,
	)

	router.Route("/files", func(r chi.Router) {
		r.Route("/list", func(r chi.Router) {
			r.Get("/public", filesController.HandleListPublicFiles())
			r.With(middlewares.Authentication).Get("/private", filesController.HandleListPrivateFiles())
		})

		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.Authentication)
			r.Post("/", filesController.HandleFileUpload())
			r.Delete("/delete", filesController.HandleFileDelete())
			r.Delete("/{id}", filesController.HandleFileDeleteById())
			r.Get("/ls", filesController.HandleListDirectories())
			r.Get("/grep", filesController.HandleListByDirectory())
			r.Get("/{id}", filesController.HandleFindById())
			r.Patch("/{id}", filesController.HandleUpdateDirectoryById())
			r.Patch("/{id}/replace", filesController.HandleFileReplace())
		})
	})
}
