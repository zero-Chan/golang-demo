package http

import (
	"net/http"
)

type HTTPServer struct {
	HProcessor HTTPProcessor
	mux        *http.ServeMux
}

func NewHTTPServer() *HTTPServer {
	svr := &HTTPServer{
		mux: http.NewServeMux(),
	}

	return svr
}

func (this *HTTPServer) SetProcessor(pros HTTPProcessor) {
	this.HProcessor = pros
}

func (this *HTTPServer) RegisterHandler(path string, handler HTTPHandler) {
	ctl := NewHTTPController()
	ctl.svr = this
	//	ctl.HProcessor = this.HProcessor
	ctl.Handler = handler

	this.mux.Handle(path, ctl)
}

func (this *HTTPServer) Serve(addr string) error {
	return http.ListenAndServe(addr, this.mux)
}
