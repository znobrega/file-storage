package main

import (
	"github.com/znobrega/file-storage/cmd/filesd/routes"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/internal/container"
)

// @title File API
// @version 1.0.0
// @description API documentation for File API
// @termsOfService http://swagger.io/terms/

// @contact.name Carlos Nóbrega
// @contact.email nobreqacarlosjr@gmail.com

// @host localhost:8090

// @securityDefinitions.apikey userIdAuthentication
// @in header
// @name Authorization
func main() {

	configs.LoadConfig("./resources")

	dependencies := container.Injector(configs.Viper)

	routes.Server{}.InitializeHttpServer(dependencies)
}
