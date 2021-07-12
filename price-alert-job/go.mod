module github.com/micheltank/crypto-price-alert/price-alert-job

go 1.16

replace github.com/micheltank/crypto-price-alert/price-alert-service => ../price-alert-service

require (
	github.com/Netflix/go-env v0.0.0-20210215222557-e437a7e7f9fb
	github.com/confluentinc/confluent-kafka-go v1.7.0
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.1.2
	github.com/micheltank/crypto-price-alert/price-alert-service v0.1.0
	github.com/onsi/gomega v1.14.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	google.golang.org/grpc v1.33.1
)
