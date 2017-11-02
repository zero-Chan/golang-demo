package http

import (
	"fmt"
)

type Processor interface {
	PreProcess()
	PostProcess()
}

type HTTPProcessor interface {
	NewPtr() HTTPProcessor
	Init4HTTP(*HTTPChannel)
	Processor
}

type Handler interface {
	Processor
	Handle()
}

type HTTPHandler interface {
	NewPtr() HTTPHandler
	Init4HTTP(*HTTPChannel)
	Handler
}

type ChannelInit func(ch *HTTPChannel)

type HTTPProcessorT1 struct {
	Chan *HTTPChannel
}

func (this *HTTPProcessorT1) Init4HTTP(httpch *HTTPChannel) {
	this.Chan = httpch
}

func (this *HTTPProcessorT1) NewPtr() HTTPProcessor {
	return new(HTTPProcessorT1)
}

func (this *HTTPProcessorT1) PreProcess() {
	fmt.Printf("Session[%s].HTTPProcessor.PreProcess...\n", this.Chan.SessionID)
}

func (this *HTTPProcessorT1) PostProcess() {
	fmt.Printf("Session[%s].HTTPProcessor.PostProcess...\n", this.Chan.SessionID)
}

type HTTPHandlerT1 struct {
	Chan *HTTPChannel
}

func (this *HTTPHandlerT1) Init4HTTP(httpch *HTTPChannel) {
	this.Chan = httpch
}

func (this *HTTPHandlerT1) NewPtr() HTTPHandler {
	return new(HTTPHandlerT1)
}

func (this *HTTPHandlerT1) PreProcess() {
	fmt.Printf("Session[%s].HTTPHandlerT1.PreProcess...\n", this.Chan.SessionID)
}

func (this *HTTPHandlerT1) PostProcess() {
	fmt.Printf("Session[%s].HTTPHandlerT1.PostProcess...\n", this.Chan.SessionID)
}

func (this *HTTPHandlerT1) Handle() {
	fmt.Printf("Session[%s].HTTPHandlerT1.Handle...\n", this.Chan.SessionID)
}
