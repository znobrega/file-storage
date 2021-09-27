package routes

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
	"github.com/znobrega/file-storage/docs"
	"github.com/znobrega/file-storage/internal/container"
	"log"
	"net/http"
)

type AppRouter interface {
	InitializeHttpServer()
}

type Server struct {
}

type Router interface {
	Routes(router chi.Router)
}

func (s Server) InitializeHttpServer(dependencies container.Dependencies) {
	r := chi.NewRouter()
	r.Use(s.cors())

	docs.SwaggerInfo.Title = "carlos"
	r.Get("/swagger/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		s, _ := swag.ReadDoc()
		_, _ = w.Write([]byte(s))
	})
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/swagger.json")))

	routes := make([]Router, 0)
	routes = append(routes,
		NewUtilsRouter(),
		NewUsersRouter(dependencies.ServicesContainer.UsersService),
		NewFilesRouter(dependencies),
	)

	for _, route := range routes {
		route.Routes(r)
	}

	port := viper.GetString("server.port")
	log.Println("server running on port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func (s Server) cors() func(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposedHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type", "Content-Disposition"},
		AllowCredentials:   true,
		MaxAge:             300,
		OptionsPassthrough: false,
	}).Handler
}
