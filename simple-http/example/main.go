package main

import (
	"fmt"

	"zero-Chan/godemo/simple-http"
)

func main() {
	svr := http.NewHTTPServer()
	svr.RegisterHandler("/test/t1", new(http.HTTPHandlerT1))
	svr.SetProcessor(new(http.HTTPProcessorT1))

	err := svr.Serve("localhost:6666")
	if err != nil {
		fmt.Println(err)
		return
	}
}
