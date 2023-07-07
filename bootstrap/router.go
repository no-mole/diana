package bootstrap

import (
	"context"
	"diana/controller"

	"github.com/gin-gonic/gin"
	"github.com/no-mole/neptune/app"
	"github.com/no-mole/neptune/middlewares"
)

func InitRouter(router *gin.Engine) app.HookFunc {
	return func(ctx context.Context) error {
		//using middleware
		router.Use(middleware.Recover, middleware.StartTracing, middleware.AccessLog)

		r := router.Group("/")
		r.GET("auth", controller.Auth)
		return nil
	}
}
