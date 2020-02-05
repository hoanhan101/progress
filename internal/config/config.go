package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerOptions `json:"server"`
	DB     DBOptions     `json:"db"`
}

type ServerOptions struct {
	Address string `json:"address"`
}

type DBOptions struct {
	User     string `json:"user" pg:"user"`
	Password string `json:"password" pg:"password"`
	Host     string `json:"host" pg:"host"`
	Port     string `json:"port" pg:"port"`
	SSLMode  string `json:"sslmode" pg:"sslmode"`
}

// Load loads the configuration.
func Load() (*Config, error) {
	v := viper.New()

	// Set config name and path.
	v.SetConfigName("config")
	v.AddConfigPath(".")

	// Read in config file.
	err := v.ReadInConfig()
	if err != nil {
		// FIXME - does not require a config file for now.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, errors.Wrap(err, "failed to read in config")
		}
	}

	// Load in the default config values
	v.SetDefault("server.address", ":8000")

	// Marshall config options into Config struct.
	var cfg = Config{}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config options")
	}

	return &cfg, nil
}
