package config

import (
	"context"
	"github.com/guilhermefbarbosa/travel-planner/api/travel-planner-go/pkg/database"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Database database.Config
}

var globalConfig *Config

func NewConfig() Config {
	if globalConfig != nil {
		return *globalConfig
	}

	config := Config{}
	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		panic(err)
	}

	globalConfig = &config
	return config
}
