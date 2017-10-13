package main

import (
	"encoding/json"
	"fmt"

	"code-lib/notify/rabbitmq"
	proto "subassembly/timer-controller/proto/notify"
)

func main() {
	notify, err := rabbitmq.NewRabbitNotify(&rabbitmq.RabbitNotifyConf{
		RabbitClientConf: &rabbitmq.RabbitClientConf{
			Host:     "localhost",
			Port:     5672,
			UserName: "guest",
			Password: "guest",
			VHost:    "/",
		},
		ConsumerInuse: true,
		Exchange:      "cza.test.hello",
		RoutingKey:    "hello.queue",
		QueueName:     "hello.queue",
		Kind:          "direct",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = notify.Receive()
	if err != nil {
		fmt.Println(err)
		return
	}

	for data := range notify.Pop() {

		val := proto.TimerNotice{}
		err = json.Unmarshal(data, &val)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("val = %+v\n", val)

		err = notify.Ack()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
