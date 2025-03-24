package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/movie-app-crud-gorm/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"log"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func RegisterRoutes(
	lc fx.Lifecycle,
	router *gin.Engine,
	movieHandler *MovieHandler,
	authHandler *AuthHandler,
	middleware JwtMiddleware,
) {
	// Auth endpoints
	noAuth := router.Group("")
	noAuth.POST("/sign-up", authHandler.SignUp)
	noAuth.POST("/login", authHandler.Login)

	// Movie endpoints
	movies := router.Group("", middleware.Middleware())
	movies.POST("/movies", movieHandler.Create)
	movies.GET("/movies/:id", movieHandler.GetByID)
	movies.GET("/movies", movieHandler.GetAll)
	movies.PUT("/movies/:id", movieHandler.Update)
	movies.DELETE("/movies/:id", movieHandler.Delete)

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := router.Run()
				if err != nil {
					log.Fatal(fmt.Sprintf("Fail to start the router: %v", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
