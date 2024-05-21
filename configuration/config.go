package configuration

import (
	"crud-golang/app"
	"fmt"

	"github.com/iamolegga/enviper"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database Database
}

type Database struct {
	Port     string `env:"CRUDGOLANG_DATABASE_PORT"`
	Host     string `env:"CRUDGOLANG_DATABASE_HOST"`
	User     string `env:"CRUDGOLANG_DATABASE_USER"`
	Password string `env:"CRUDGOLANG_DATABASE_PASSWORD"`
	Name     string `env:"CRUDGOLANG_DATABASE_NAME"`
}

func GetConfig(config Configuration) error {
	e := enviper.New(viper.New())

	e.SetEnvPrefix(app.EnvPrefix)
	e.SetConfigFile(app.ConfigFilePath)
	e.SetConfigName(app.ConfigName)

	err := e.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("could not unmarshal config into the struct: %s", err)
	}

	return nil
}
