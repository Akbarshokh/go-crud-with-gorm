package rest

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewMovieHandler,
		NewAuthHandler,
		NewJwtMiddleware,
		NewRouter,
	),
	fx.Invoke(RegisterRoutes),
)
