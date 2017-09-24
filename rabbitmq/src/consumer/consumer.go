package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("amqp dial fail:", err)
		return
	}
	defer conn.Close()

	rmqch, err := conn.Channel()
	if err != nil {
		fmt.Println("amqp get channel fail:", err)
		return
	}
	defer rmqch.Close()

	rmqQueue, err := rmqch.QueueDeclare("myfirstqueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println("amqp get queue fail:", err)
		return
	}

	err = rmqch.QueueBind(rmqQueue.Name, "myfirstRoutingKey", "myfirstExchange", false, nil)
	if err != nil {
		fmt.Println("amqp bind queue to exchange fail:", err)
		return
	}

	msgSource, err := rmqch.Consume("myfirstqueue", "myfirstconsumer", false, false, false, false, nil)
	if err != nil {
		fmt.Println("amqp new consumer fail:", err)
		return
	}

	forever := make(chan bool)

	go func(msgSource <-chan amqp.Delivery) {
		for data := range msgSource {
			fmt.Println("consumer get data:", data.Body)
			err := data.Ack(false)
			if err != nil {
				fmt.Println("amqp ack fail:", err)
				continue
			}
		}
	}(msgSource)

	<-forever
}
