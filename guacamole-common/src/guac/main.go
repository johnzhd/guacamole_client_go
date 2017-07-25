package main

import (
	"fmt"
	logger "github.com/Sirupsen/logrus"
	"guacamole_client_go/gnet"
	"net/http"
)

func main() {
	logger.SetLevel(logger.DebugLevel)
	// init servlet
	servlet := NewServletHandleStruct(DemoDoConnect)
	// ...

	// init handle
	myHandler := http.NewServeMux()
	myHandler.Handle(servlet.GetEnterPath(), &servlet)

	myHandler.Handle("/", http.FileServer(http.Dir("./")))

	// init server
	s := &http.Server{
		Addr:           ":4567",
		Handler:        myHandler,
		ReadTimeout:    gnet.SocketTimeout,
		WriteTimeout:   gnet.SocketTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
