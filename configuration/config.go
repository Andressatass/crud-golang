package configuration

import (
	"crud-golang/app"
	"fmt"

	"github.com/iamolegga/enviper"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database Database `env:"CRUDGOLANG_DATABASE"`
	API      API      `env:"CRUDGOLANG_API"`
}

type Database struct {
	Port     string `env:"CRUDGOLANG_DATABASE_PORT"`
	Host     string `env:"CRUDGOLANG_DATABASE_HOST"`
	User     string `env:"CRUDGOLANG_DATABASE_USER"`
	Password string `env:"CRUDGOLANG_DATABASE_PASSWORD"`
	Name     string `env:"CRUDGOLANG_DATABASE_NAME"`
}

type API struct {
	Port string `env:"CRUDGOLANG_API_PORT"`
}

func GetConfig() (Configuration, error) {
	var config Configuration

	e := enviper.New(viper.New())

	e.SetEnvPrefix(app.EnvPrefix)
	e.SetConfigFile(app.ConfigFilePath)
	e.SetConfigName(app.ConfigName)

	err := e.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("could not unmarshal config into the struct: %s", err)
	}

	return config, nil
}
