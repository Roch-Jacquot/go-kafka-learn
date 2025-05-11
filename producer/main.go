package main

import (
	"bufio"
	"fmt"
	"os"

	client "github.com/Roch-Jacquot/go-kafka-learn/producer/internal/kafka_client"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	brokers := []string{"localhost:9092"}
	kafkaclient, err := client.NewKafkaClient(brokers)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer kafkaclient.Close()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(line, " ", err.Error())
			os.Exit(1)
		}

		err = kafkaclient.ProduceMessage("first_topic", line)
		if err != nil {
			fmt.Println(line, " kafka error :", err.Error())
			os.Exit(1)
		}

		fmt.Println("message: ", line)
	}

}
