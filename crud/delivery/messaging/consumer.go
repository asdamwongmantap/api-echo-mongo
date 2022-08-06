package main

import (
	"fmt"
	"log"

	"github.com/asdamwongmantap/api-echo-mongo/lib/queue"
)

func main() {
	rabbitMq, err := queue.Connect(queue.QueueConfig{
		Username: "guest",
		Password: "guest",
		Host:     "0.0.0.0",
		Port:     "5672",
	})
	if err != nil {
		log.Println(err)
		return
	}

	ch, err := rabbitMq.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"MsgRabbitGo",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("failed to consume")
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
