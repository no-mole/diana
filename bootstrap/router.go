package bootstrap

import (
	"context"
	"diana/api"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/neptune/app"
	"github.com/no-mole/neptune/middlewares"
)

func InitRouter(router *gin.Engine) app.HookFunc {
	return func(ctx context.Context) error {
		//using middleware
		router.Use(middleware.Recover, middleware.StartTracing, middleware.AccessLog)

		rg := router.Group("/app")
		rg.POST("", api.RegisterApp)
		return nil
	}
}
