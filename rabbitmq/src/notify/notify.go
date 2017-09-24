package notify

import (
	"errors"
	"time"

	"heimdall/lib/redis"
)

// 通知底层封装
// 可选择redis等等MQ实现

type INotify interface {
	Name() string
	Receive()
	Pop() <-chan []byte
	Push([]byte) error
	StopPop()
}

// redis >> left>right
type Notify struct {
	*redis.MQ
	name string
	data chan []byte
	stop chan bool
}

func NewNotify(name string) (notify *Notify, err error) {
	mq, ok := redis.RedisClientMap[name]
	if !ok {
		err = errors.New("not exists the mq : " + name)
		return
	}

	notify = &Notify{
		MQ:   mq,
		name: name,
		data: make(chan []byte),
		stop: make(chan bool),
	}
	return
}

func (n *Notify) Name() string {
	return n.name
}

func (n *Notify) Receive() {
	go n.pop()
}

func (n *Notify) pop() {
	tick := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-n.stop:
			return
		case <-tick.C:
			stringSliceCmd := n.MQ.BRPop(1*time.Second, n.MQ.QueueName)
			if stringSliceCmd.Err() != nil {
				break
			}

			if len(stringSliceCmd.Val()) > 1 {
				n.data <- []byte(stringSliceCmd.Val()[1])
			}
		}
	}
}

func (n *Notify) StopPop() {
	n.stop <- true
}

func (n *Notify) Pop() <-chan []byte {
	return n.data
}

func (n *Notify) Push(data []byte) (err error) {
	intCmd := n.MQ.LPush(n.MQ.QueueName, string(data))
	return intCmd.Err()
}
