package main

import (
	"fmt"
	"os"

	client "github.com/Roch-Jacquot/go-kafka-learn/producer/internal/kafka_client"
)

func main() {

	//reader := bufio.NewReader(os.Stdin)
	brokers := []string{"kafka:9094"}
	kafkaclient, err := client.NewKafkaClient(brokers)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer kafkaclient.Close()
	line := "test"
	fmt.Println("line is ", line)
	err = kafkaclient.ProduceMessage("first_topic", line)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
