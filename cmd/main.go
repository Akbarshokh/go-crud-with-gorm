package main

import (
	"github.com/joho/godotenv"
	"github.com/movie-app-crud-gorm/api/docs"
	"github.com/movie-app-crud-gorm/internal/bootstrap"
	"github.com/movie-app-crud-gorm/internal/rest"
	"go.uber.org/fx"
)

func main() {
	_ = godotenv.Load()

	docs.SwaggerInfo.Title = "Movie CRUD API"
	docs.SwaggerInfo.Description = "Simple REST API using Gin, GORM, UberFx"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	fx.New(
		bootstrap.Modules,
		rest.Module).Run()
}
