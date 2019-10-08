package main

import (
	"fmt"

	"log"
	"net/http"
)

type Configuration struct {
	Ip          string
	Port        int16
	Certificate string
	Key         string
}

type httpHandler func(http.ResponseWriter, *http.Request)

var getConfiguration = (func() func() Configuration {
	configuration := Configuration{
		Ip:          "::1",
		Port:        4343,
		Certificate: "server.crt",
		Key:         "server.key",
	}

	return (func() Configuration { return configuration })
})()

func main() {

	addr := fmt.Sprintf("[%s]:%d", getConfiguration().Ip, getConfiguration().Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(compose([]httpHandler{logger, kevin})),
	}
	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(getConfiguration().Certificate, getConfiguration().Key))
}

func logger(w http.ResponseWriter, r *http.Request) {

	log.Printf("Got connection: '%s'", r.Proto)
}

func kevin(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("YOLO!!!!!!!"))
}

func compose(handlerList []httpHandler) httpHandler {

	return func(w http.ResponseWriter, r *http.Request) {

		for _, handler := range handlerList {

			handler(w, r)
		}
	}
}
