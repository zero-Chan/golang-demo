package main

import (
	"fmt"
	"net/url"

	"github.com/streadway/amqp"
)

func main() {
	info := url.URL{}

	info.Scheme = "amqp"

	info.User = url.UserPassword("guest", "guest")

	info.Host = "127.0.0.1:5672"

	info.Path = "cza"

	fmt.Println(info.String())

	//	conn, err := amqp.Dial(info.String())
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

	//	fmt.Println(conn.)
}
