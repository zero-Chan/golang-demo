package conf

import (
	"net/url"
	"strconv"
)

type Conf struct {
	MyMq1 *RabbitmqConf
}

type RabbitmqConf struct {
	Host       string
	Port       int64
	UserName   string
	Password   string
	VHost      string
	Exchange   string
	RoutingKey string
	QueueName  string
	BindKind   string // direct, fanout, topic
}

func (this *RabbitmqConf) Addr() string {
	if len(this.Host) == 0 || this.Port <= 0 {
		return ""
	}

	return this.Host + ":" + strconv.FormatInt(this.Port, 10)
}

// amqp_URI = "amqp:// amqp_authority ["/" vhost] ["?" query]
// amqp_authority = [amqp_userinfo "@"] host [":" port]
// amqp_userinfo = username [":" password]
// username = *(unreserved / pct-encoded / sub-delims)
// password = *(unreserved / pct-encoded / sub-delims)
// vhost = segment

func (this *RabbitmqConf) String() string {
	amqpUri := url.URL{}

	if this.Addr() == "" {
		return ""
	}

	amqpUri.Scheme = "amqp"
	amqpUri.User = url.UserPassword(this.UserName, this.Password)
	amqpUri.Host = this.Addr()

	if this.VHost != "" {
		amqpUri.Path = this.VHost
	}

	return amqpUri.String()
}

// QueueNameNS = "/vhost/exchange/queueName"
func (this RabbitmqConf) QueueNameNS() string {
	return "/" + this.VHost + "/" + this.Exchange + "/" + this.QueueName
}
