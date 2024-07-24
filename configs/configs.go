package configs

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/multierr"
)

type Configs struct {
	Application struct {
		Name    string `env:"NAME"`
		Version string `env:"VERSION" env-default:"development"`
	} `env-prefix:"APPLICATION_"`

	Log struct {
		LogLevel  string `env:"LEVEL" env-default:"debug"`
		SentryDSN string `env:"SENTRY_DSN"`
	} `env-prefix:"LOG_"`
}

// Global configs
var C Configs

func init() {
	configs := &Configs{}

	if errConfig := cleanenv.ReadConfig(".env", configs); errConfig != nil {
		if errEnv := cleanenv.ReadEnv(configs); errEnv != nil {
			log.Fatal(multierr.Combine(errConfig, errEnv))
		}
	}

	C = *configs
}
