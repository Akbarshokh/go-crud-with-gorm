package bootstrap

import (
	"github.com/movie-app-crud-gorm/internal/config"
	"github.com/movie-app-crud-gorm/internal/domain"
	"github.com/movie-app-crud-gorm/internal/drivers/dbstore"
	"github.com/movie-app-crud-gorm/internal/drivers/dbstore/movies"
	"github.com/movie-app-crud-gorm/internal/drivers/dbstore/user"
	"github.com/movie-app-crud-gorm/internal/pkg/logger"
	"github.com/movie-app-crud-gorm/internal/rest"
	"github.com/movie-app-crud-gorm/internal/usecases/auth"
	"github.com/movie-app-crud-gorm/internal/usecases/movies"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	infraModule,
	repoModule,
	usecaseModule,
	rest.Module,
)

var infraModule = fx.Options(
	fx.Provide(config.NewConfig),
	fx.Provide(func(cfg config.Config) logger.Logger {
		return logger.New(cfg.LogLevel, "movie-app")
	}),
	fx.Provide(dbstore.NewGormDB),
)

var repoModule = fx.Options(
	fx.Provide(movies_repo.New),
	fx.Provide(func(r *movies_repo.Repo) domain.MovieRepository {
		return r
	}),
	fx.Provide(user.New),
	fx.Provide(func(a *user.Repo) domain.UserRepository {
		return a
	}),
)

var usecaseModule = fx.Options(
	fx.Provide(
		movies.New,
		func(uc *movies.UseCase) domain.MovieUseCase {
			return uc
		},

		auth.New,
		func(uc *auth.UseCase) domain.AuthUseCase {
			return uc
		},
	),
)
