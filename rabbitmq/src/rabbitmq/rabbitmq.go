package rabbitmq

import (
	"log"
	"sync"

	"github.com/streadway/amqp"

	"conf"
)

var (
	RMQConnMarker rmqConnMarker = rmqConnMarker{
		// map[amqpUrl]*amqp.Connection
		conn: map[string]*amqp.Connection{},
	}

	RMQMarker rmqMarker = rmqMarker{
		// map[queueName]*MQ
		MQ: map[string]*MQ{},
	}
)

type rmqConnMarker struct {
	conn  map[string]*amqp.Connection
	mutex sync.Mutex
}

func (this *rmqConnMarker) Set(key string, val *amqp.Connection) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.conn[key] = val
}

func (this *rmqConnMarker) Get(key string) (val *amqp.Connection, exist bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	val, exist = this.conn[key]
	return
}

func (this *rmqConnMarker) GetSet(key string, val *amqp.Connection) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if _, exist := this.conn[key]; !exist {
		this.conn[key] = val
	}
}

type rmqMarker struct {
	MQ    map[string]*MQ
	mutex sync.Mutex
}

func (this *rmqMarker) Set(key string, val *MQ) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.MQ[key] = val
}

func (this *rmqMarker) Get(key string) (val *MQ, exist bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	val, exist = this.MQ[key]
	return
}

func (this *rmqMarker) GetSet(key string, val *MQ) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if _, exist := this.MQ[key]; !exist {
		this.MQ[key] = val
	}
}

type MQ struct {
	Vhost      string
	Exchange   string
	RoutingKey string
	QueueName  string
	conn       *amqp.Connection
	AmqpChan   *amqp.Channel
}

func NewMQ(conn *amqp.Connection, vhost string, exchange string, bindKind string, routingKey string, queueName string) (mq *MQ, err error) {
	ch, err := conn.Channel()
	if err != nil {
		return
	}

	err = ch.ExchangeDeclare(exchange, bindKind, false, false, false, false, nil)
	if err != nil {
		return
	}

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return
	}

	err = ch.QueueBind(queueName, routingKey, exchange, false, nil)
	if err != nil {
		return
	}

	mq = &MQ{
		Vhost:      vhost,
		Exchange:   exchange,
		RoutingKey: routingKey,
		QueueName:  queueName,
		conn:       conn,
		AmqpChan:   ch,
	}

	return
}

func (this *MQ) Close() (err error) {
	this.Close()
	return this.conn.Close()
}

func InitRMQ(c conf.Conf) {
	var err error

	if c.MyMq1 != nil {
		url := c.MyMq1.String()
		if url == "" {
			log.Panicf("Amqp can not dial empty url.")
		}

		conn, exist := RMQConnMarker.Get(url)
		if !exist {
			conn, err = amqp.Dial(url)
			if err != nil {
				log.Panicf("Amqp dial [%s] fail: %s", url, err)
				return
			}
			RMQConnMarker.Set(url, conn)
		}

		queueName := c.MyMq1.QueueName
		mq, exist := RMQMarker.Get(queueName)
		if !exist {
			mq, err = NewMQ(conn, c.MyMq1.VHost, c.MyMq1.Exchange, c.MyMq1.BindKind, c.MyMq1.RoutingKey, queueName)
			if err != nil {
				log.Panicf("NewMQ[%s] fail: %s", queueName, err)
				return
			}
			RMQMarker.Set(queueName, mq)
		}
	}
}
