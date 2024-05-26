package settings

import "github.com/caarlos0/env/v11"

type AppEnvironment struct {
	Port string `env:"APP_PORT" envDefault:"3000"`
}

type DatabaseEnvironment struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Database string `env:"POSTGRES_DB"`
	Password string `env:"POSTGRES_PASSWORD"`
}

type Environment struct {
	App      AppEnvironment
	Database DatabaseEnvironment
}

var envs Environment

func LoadEnvs() error {
	db := DatabaseEnvironment{}
	if err := env.Parse(&db); err != nil {
		return err
	}

	app := AppEnvironment{}
	if err := env.Parse(&app); err != nil {
		return err
	}

	envs = Environment{
		Database: db,
		App:      app,
	}

	return nil
}

func GetEnvs() Environment {
	return envs
}
