// Package config will provide enviroment config strict for gateway microservice
package config

type Config struct {
	GatewayPort string `env:"GATEWAY_PORT" envDefault:"8080"`
	TimeOut     string `env:"TIME_OUT" envDefault:"15s"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
}
