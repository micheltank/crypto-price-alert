package infra

import (
	"github.com/Netflix/go-env"
	"log"
)

type Environment struct {
	CryptoCompareApiHost    string `env:"CRYPTO_COMPARE_API_HOST"`
	CryptoCompareApiKey     string `env:"CRYPTO_COMPARE_API_KEY"`
	PriceAlertServiceHost   string `env:"PRICE_ALERT_SERVICE_HOST"`
	NotificationServiceHost string `env:"NOTIFICATION_SERVICE_HOST"`
	Periodicity             string `env:"PERIODICITY"`
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
