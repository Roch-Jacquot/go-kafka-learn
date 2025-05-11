package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "first_topic"
	msgCount := 0

	worker, err := connectKafkaBroker([]string{"localhost:9092"})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("consumer started")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
				message := string(msg.Value)
				fmt.Printf("message: %s\n", message)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed ", msgCount, " messages")

	err = worker.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func connectKafkaBroker(brokers []string) (sarama.Consumer, error) {
	fmt.Println("preparing consumer")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}
