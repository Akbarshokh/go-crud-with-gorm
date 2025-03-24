package rest

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMovieHandler),
	fx.Provide(NewRouter),
	fx.Invoke(RegisterRoutes))
