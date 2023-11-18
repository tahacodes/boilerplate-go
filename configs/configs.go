package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/multierr"
)

type Configuration struct {
	App struct {
		Version string `env:"VERSION" env-default:"development"`
	} `env-prefix:"APPLICATION_"`
}

var Config *Configuration

func InitConfigs() error {
	configuration := &Configuration{}

	if errConfig := cleanenv.ReadConfig(".env", configuration); errConfig != nil {
		if errEnv := cleanenv.ReadEnv(configuration); errEnv != nil {
			return multierr.Combine(errConfig, errEnv)
		}
	}

	Config = configuration

	return nil
}
