package infra

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	errors "github.com/pkg/errors"
)

type IKafkaProducer interface {
	Publish(msg []byte, topic string, key []byte) error
}

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (IKafkaProducer, error) {
	producer, err := newKafkaProducer()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create kafka producer")
	}
	return &KafkaProducer{
		producer: producer,
	}, nil
}

func newKafkaProducer() (*kafka.Producer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (k *KafkaProducer) Publish(msg []byte, topic string, key []byte) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
		Key:            key,
	}
	err := k.producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}
