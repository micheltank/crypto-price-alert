package infra

import (
	"github.com/Netflix/go-env"
	"log"
)

type Environment struct {
	EmailSender             string `env:"EMAIL_SENDER"`
	EmailPassword           string `env:"EMAIL_PASSWORD"`
	Extras                  env.EnvSet
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
