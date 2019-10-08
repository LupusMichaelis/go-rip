package main

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

type Configuration struct {
	Ip          string
	Port        int16
	Certificate string
	Key         string
}

type httpHandler func(rest.ResponseWriter, *rest.Request)

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

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(compose([]httpHandler{logger, kevin})))
	apiHandler := api.MakeHandler()

	addr := fmt.Sprintf("[%s]:%d", getConfiguration().Ip, getConfiguration().Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: apiHandler,
	}
	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(getConfiguration().Certificate, getConfiguration().Key))
}

func logger(w rest.ResponseWriter, r *rest.Request) {

	log.Printf("Got connection: '%s'", r.Proto)
}

func kevin(w rest.ResponseWriter, r *rest.Request) {

	w.WriteJson(map[string]string{"Booty": "Where is it?!!"})
}

func compose(handlerList []httpHandler) httpHandler {

	return func(w rest.ResponseWriter, r *rest.Request) {

		for _, handler := range handlerList {

			handler(w, r)
		}
	}
}
