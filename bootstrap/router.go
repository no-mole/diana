package bootstrap

import (
	"context"

	"diana/controller/bar"

	"github.com/gin-gonic/gin"
	"github.com/no-mole/neptune/app"
	"github.com/no-mole/neptune/middlewares"
)

func InitRouter(router *gin.Engine) app.HookFunc {
	return func(ctx context.Context) error {
		//using middleware
		router.Use(middleware.Recover, middleware.StartTracing, middleware.AccessLog)

		groupBar := router.Group("/bar")
		groupBar.GET("/say_hello", bar.SayHello)

		return nil
	}
}