package http

import (
	"net/http"

	"github.com/pborman/uuid"
)

type HTTPChannel struct {
	ctl *HTTPController

	SessionID string
	request   *http.Request
	responsew http.ResponseWriter

	Processor Processor
	Handler   Handler
}

func NewHTTPChannel() *HTTPChannel {
	hc := &HTTPChannel{
		SessionID: uuid.New(),
	}

	return hc
}

//func (this *HTTPChannel)
