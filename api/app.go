package api

import (
	service2 "diana/service"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/neptune/enum"
	"github.com/no-mole/neptune/output"
)

type registerAppReq struct {
	AppName     string `json:"app_name" form:"app_name" binding:"required"`
	Description string `json:"description" form:"description"`
}

type registerAppResp struct {
	AppName     string `json:"app_name"`
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	Description string `json:"description"`
}

func RegisterApp(ctx *gin.Context) {
	req := registerAppReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {

	}
	service := service2.AppService{}
	app, err := service.RegisterApp(req.AppName, req.Description)
	if err != nil {

	}
	output.Json(ctx, enum.Success, &registerAppResp{
		AppId:       app.AppId,
		AppName:     app.Name,
		AppSecret:   app.Secret,
		Description: app.Description,
	})
}
