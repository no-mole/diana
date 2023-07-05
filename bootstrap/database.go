package bootstrap

import (
	"context"
	"fmt"
	"github.com/no-mole/neptune/utils"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/no-mole/neptune/config"
	"github.com/no-mole/neptune/database"
	"github.com/no-mole/neptune/env"
)

var (
	dbNames = map[string]*database.Config{}
)

func InitDatabase(ctx context.Context) error {
	baseConfigDir := fmt.Sprintf("%s/config/%s/mysql.yaml", config.GlobalConfig.BasePath, config.GlobalConfig.Env.Mode)
	exist := utils.FileExist(baseConfigDir)
	if exist {
		body, err := os.ReadFile(baseConfigDir)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(body, &dbNames)
		if err != nil {
			return err
		}
	}

	for dbName, dbValue := range dbNames {

		err := initDatabaseDrive(dbName, dbValue)
		if err != nil {
			return err
		}

	}
	return nil
}

func initDatabaseDrive(dbName string, conf *database.Config) error {
	return database.Init(dbName, conf, env.GetEnvDebug())
}
