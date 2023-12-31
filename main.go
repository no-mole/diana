package main

import (
	"context"

	_ "go.uber.org/automaxprocs"

	"diana/bootstrap"
	neptune "github.com/no-mole/neptune/app"
	"github.com/no-mole/neptune/config"
)

func main() {
	ctx := context.Background()

	app := neptune.NewApp(ctx)

	neptune.AddHook(
		config.Init,          //初始化配置
		bootstrap.InitLogger, //初始化日志
		//bootstrap.InitDatabase,
		bootstrap.InitRouter(app.HttpEngine), //初始化http router
	)

	if err := neptune.Start(); err != nil {
		panic(err)
	}

	err := <-neptune.ErrorCh()
	neptune.Stop()
	panic(err)
}
