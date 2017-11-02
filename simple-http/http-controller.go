package http

import (
	"net/http"
)

type HTTPController struct {
	svr *HTTPServer
	//	HProcessor HTTPProcessor
	Handler HTTPHandler
}

func NewHTTPController() *HTTPController {
	hc := &HTTPController{}
	return hc
}

func (this *HTTPController) Channel(initfunc ChannelInit) (httpch *HTTPChannel) {
	httpch = NewHTTPChannel()
	httpch.ctl = this

	hp := this.svr.HProcessor.NewPtr()
	hp.Init4HTTP(httpch)
	httpch.Processor = hp

	hh := this.Handler.NewPtr()
	hh.Init4HTTP(httpch)
	httpch.Handler = hh

	initfunc(httpch)

	return
}

func (this *HTTPController) ServeHTTP(respw http.ResponseWriter, req *http.Request) {

	var initfunc ChannelInit = func(ch *HTTPChannel) {
		ch.request = req
		ch.responsew = respw
	}

	httpch := this.Channel(initfunc)
	httpch.Processor.PreProcess()
	httpch.Handler.PreProcess()
	httpch.Handler.Handle()
	httpch.Handler.PostProcess()
	httpch.Processor.PostProcess()

	respw.WriteHeader(200)
}
