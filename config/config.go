package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server server
	App    app
}

type server struct {
	Port string `envconfig:"PORT" default:"8080"`
}

type app struct {
}

var cfg Config

func init() {
	_ = godotenv.Load()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() Config {
	return cfg
}
