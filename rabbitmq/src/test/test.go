package main

import (
	"fmt"
	"time"

	"conf"
	"notify"
	"rabbitmq"
)

func main() {
	config := conf.Conf{
		MyMq1: &conf.RabbitmqConf{
			Host:       "localhost",
			Port:       5672,
			UserName:   "guest",
			Password:   "guest",
			BindKind:   "direct",
			Exchange:   "myfirstExchange",
			RoutingKey: "myfirstRoutingKey",
			QueueName:  "myfirstqueue",
		},
	}

	rabbitmq.InitRMQ(config)

	mq, err := notify.NewRabbitNotify("myfirstqueue")
	if err != nil {
		fmt.Println(err)
		return
	}

	mq.Receive()

	count := 0

	go func() {
		data := mq.Pop()
		for {
			select {
			case <-data:
				err := mq.Ack()
				if err != nil {
					fmt.Println(err)
					return
				}
				count++
			}
		}

		//		for msg := range data {
		//			//			fmt.Println(string(msg))
		//			err := mq.Ack()
		//			if err != nil {
		//				fmt.Println(err)
		//				return
		//			}
		//			count++
		//		}
	}()

	go func() {
		tick := time.NewTicker(time.Second * 1)
		sec := 1

		for {
			select {
			case <-tick.C:
				fmt.Printf("%d sec can consumer count: %d\n", sec, count)
				sec++
			}
		}

	}()

	select {}

}
