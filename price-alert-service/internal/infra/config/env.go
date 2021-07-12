package config

import (
	"github.com/Netflix/go-env"
	"log"
)

type Environment struct {
	Port     int `env:"PORT"`
	GrpcPort int `env:"GRPC_PORT"`
	DbConfig DbConfig
	Extras   env.EnvSet
}

type DbConfig struct {
	User     string `env:"DB_USER"`
	Port     string `env:"DB_PORT"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
}

var Env Environment

func init() {
	es, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}
	// Remaining environment variables.
	Env.Extras = es
}
