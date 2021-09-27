package routes

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/znobrega/file-storage/cmd/filesd/middlewares"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"io"
	"net/http"
)

type utilsRouter struct{}

func NewUtilsRouter() Router {
	return utilsRouter{}
}

func (U utilsRouter) Routes(router chi.Router) {
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api up and running"))
	})

	router.With(middlewares.Authentication).Get("/static/*", U.provideStaticFile())
}

type file io.Writer

// @Title getStatic file
// @Tags Files
// @Security userIdAuthentication
// @Param Authorization header string true "Bearer <token> is required to download/view"
// @Param fullfilepath path string true "router example: http://localhost:8090/static/example/updated/fix/splunk.pdf"
// @Summary Download static file (doesnt working on swagger)
// @Description Download static file (doesnt working on swagger)
// @Success 200 {file} file
// @Failure 400 "Bad request"
// @Accept json
// @Router /static/{fullfilepath} [get]
func (U utilsRouter) provideStaticFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(helpers.ContextUserKey).(uint64)
		fileServerDir := fmt.Sprintf("./%s/%d", configs.Viper.GetString("files.directory"), userID)
		fs := http.FileServer(http.Dir(fileServerDir))

		http.StripPrefix("/static", fs).ServeHTTP(w, r)
	}
}
