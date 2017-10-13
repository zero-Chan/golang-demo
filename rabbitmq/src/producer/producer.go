package main

import (
	"fmt"
	"time"

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
		fmt.Println("amqp new channel fail:", err)
		return
	}
	defer rmqch.Close()

	//	rmqQueue, err := rmqch.QueueDeclare("myfirstqueue", false, false, false, false, nil)
	//	if err != nil {
	//		fmt.Println("amqp new queue fail:", err)
	//		return
	//	}

	//	err = rmqch.ExchangeDeclare("myfirstExchange", "direct", false, false, false, false, nil)
	//	if err != nil {
	//		fmt.Println("amqp new exchange fail:", err)
	//		return
	//	}

	var count int = 0

	go func() {
		tick := time.NewTicker(time.Second * 1)
		sec := 1

		for {
			select {
			case <-tick.C:
				fmt.Printf("%d sec can publish count: %d\n", sec, count)
				sec++
			}
		}

	}()

	for {
		msgbody := `{"content": "hello world", "name": "Cza"}`
		err = rmqch.Publish("myfirstExchange", "myfirstRoutingKey", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msgbody),
		})
		if err != nil {
			fmt.Println("amqp publish message fail:", err)
			return
		}

		count++
	}

	//	fmt.Println("publish message:", msgbody)

}
