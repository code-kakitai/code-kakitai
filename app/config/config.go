package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server Server
	DB     DBConfig
	ReadDB ReadDBConfig
	Redis  Redis
}

type DBConfig struct {
	Name     string `envconfig:"DB_DATABASE" default:"code_kakitai"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASS" default:""`
	Port     string `envconfig:"DB_PORT" default:"3306"`
	Host     string `envconfig:"DB_HOST" default:"db"`
}

type ReadDBConfig struct {
	Name     string `envconfig:"DB_READ_DATABASE" default:"code_kakitai"`
	User     string `envconfig:"DB_READ_USER" default:"root"`
	Password string `envconfig:"DB_READ_PASS" default:""`
	Port     string `envconfig:"DB_READ_PORT" default:"3306"`
	Host     string `envconfig:"DB_READ_HOST" default:"db"`
}

type Redis struct {
	Host string `envconfig:"REDIS_HOST" default:"redis"`
	Port string `envconfig:"REDIS_PORT" default:"6379"`
}

type Server struct {
	Address string `envconfig:"ADDRESS" default:"0.0.0.0"`
	Port    string `envconfig:"PORT" default:"8080"`
}

var (
	once   sync.Once
	config Config
)

func GetConfig() *Config {
	// goroutine実行中でも一度だけ実行される
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			panic(err)
		}
	})
	return &config
}
