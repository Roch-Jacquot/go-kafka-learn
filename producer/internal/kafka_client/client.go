package kafkaclient

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type kafkaProducer interface {
	Close() error
	SendMessage(*sarama.ProducerMessage) (partition int32, offset int64, err error)
}

type KafkaClient struct {
	Client kafkaProducer
}

func NewKafkaClient(brokers []string) (KafkaClient, error) {
	producer, err := connectKafkaBroker(brokers)
	if err != nil {
		return KafkaClient{}, err
	}
	return KafkaClient{Client: producer}, nil
}

func connectKafkaBroker(brokers []string) (sarama.SyncProducer, error) {
	fmt.Println("preparing syncProducer")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func (k KafkaClient) ProduceMessage(topic string, message string) error {

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder([]byte(message)),
	}
	partition, offset, err := k.Client.SendMessage(msg)
	if err != nil {
		fmt.Println("failure to send message")
		return err
	}
	log.Printf("message is stored in topic (%s)/partition(%d)/offset(%d)\n",
		topic, partition, offset)
	return nil
}

func (k KafkaClient) Close() {
	k.Client.Close()
}
