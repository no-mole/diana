package service

import (
	app2 "diana/core/app"
	"diana/infras/repository"
)

type AppService struct {
	appRepository repository.AppRepository
}

func (appService *AppService) RegisterApp(appName string, description string) (*app2.App, error) {
	app, err := app2.NewApp(appName, description)
	if err != nil {
		return nil, err
	}
	return appService.appRepository.Store(app)
}
