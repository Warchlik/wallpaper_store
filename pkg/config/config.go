// Package config
package config

import (
	"fmt"
	"os"
	"strings"

	env "github.com/caarlos0/env/v11"
	godotenv "github.com/joho/godotenv"
)

func Load[T any]() (*T, error) {
	envMap, err := godotenv.Read()

	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error with loading .env file; %w", err)
	}

	if envMap == nil {
		envMap = map[string]string{}
	}

	for _, kv := range os.Environ() {
		key, value, _ := strings.Cut(kv, "=")
		envMap[key] = value
	}

	cfg := new(T)

	if err := env.ParseWithOptions(cfg, env.Options{Environment: envMap}); err != nil {
		return nil, fmt.Errorf("can not parse environment into config; %w", err)
	}

	return cfg, nil
}
