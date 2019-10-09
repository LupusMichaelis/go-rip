package main

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"

	"./config"
)

type httpHandler func(rest.ResponseWriter, *rest.Request)

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(compose([]httpHandler{logger, kevin})))
	apiHandler := api.MakeHandler()

	addr := fmt.Sprintf("[%s]:%d", config.GetConfiguration().Ip, config.GetConfiguration().Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: apiHandler,
	}
	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(config.GetConfiguration().Certificate, config.GetConfiguration().Key))
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
