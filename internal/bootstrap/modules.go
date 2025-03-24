package bootstrap

import (
	"github.com/movie-app-crud-gorm/internal/config"
	"github.com/movie-app-crud-gorm/internal/drivers/dbstore"
	"github.com/movie-app-crud-gorm/internal/drivers/dbstore/movies"
	"github.com/movie-app-crud-gorm/internal/drivers/dbstore/user"
	"github.com/movie-app-crud-gorm/internal/pkg/logger"
	"github.com/movie-app-crud-gorm/internal/usecases/movies"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		config.NewConfig,
		logger.New("Debug", "movie-app-crud-gorm"),
		//db
		dbstore.NewGormDB,

		//repo
		movies_repo.New,
		user.New,

		//useCase
		movies.New),
)
