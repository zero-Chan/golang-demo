package main

import (
	"encoding/json"
	"fmt"
	"time"

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
		Exchange:       "cza.test.timer",
		RoutingKey:     "controller_1s",
		Kind:           "direct",
		PublisherInuse: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := proto.TimerNotice{
		Destination: proto.RabbitmqDestination{
			Exchange:   "cza.test.hello",
			RoutingKey: "hello.queue",
		},

		SendUnixTime: time.Now(),
		Expire:       time.Second * 10,
		Target: json.RawMessage([]byte(`
			{
				"key": 1
			}
		`)),
	}

	senddata, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = notify.Push(senddata)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("hh")
}
