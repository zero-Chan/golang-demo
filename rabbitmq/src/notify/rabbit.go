package notify

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"rabbitmq"
)

type RabbitNotify struct {
	*rabbitmq.MQ

	stop    chan bool
	data    chan []byte
	deliver <-chan amqp.Delivery

	curDeliver amqp.Delivery
	isAck      chan bool
}

func NewRabbitNotify(name string) (n *RabbitNotify, err error) {
	mq, exist := rabbitmq.RMQMarker.Get(name)
	if !exist {
		err = fmt.Errorf("rabbitmq queue[%s] no exist.", name)
		return
	}

	n = &RabbitNotify{
		MQ:    mq,
		stop:  make(chan bool),
		data:  make(chan []byte),
		isAck: make(chan bool),
	}

	return
}

func (this *RabbitNotify) Name() string {
	return this.QueueName
}

func (this *RabbitNotify) StopPop() {
	this.stop <- true
}

func (this *RabbitNotify) Pop() <-chan []byte {
	return this.data
}

func (this *RabbitNotify) Ack() (err error) {
	err = this.curDeliver.Ack(false)
	if err != nil {
		err = fmt.Errorf("rabbitmq consumer ack error: %s", err)
		return
	}
	this.isAck <- true

	return
}

func (this *RabbitNotify) Receive() {
	go this.pop()
}

func (this *RabbitNotify) pop() {
	var err error
	this.deliver, err = this.AmqpChan.Consume(this.QueueName, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("warn: can not new rabbitmq consumer to queue[%s], err: %s", this.QueueName, err)
		return
	}

	var ok bool
	for {
		select {
		case <-this.stop:
			return
		case this.curDeliver, ok = <-this.deliver:
			if !ok {
				continue
			}

			this.data <- this.curDeliver.Body
			<-this.isAck
		}
	}
}

func (this *RabbitNotify) Push(data []byte) (err error) {
	err = this.AmqpChan.Publish(this.Exchange, this.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})

	return err
}
