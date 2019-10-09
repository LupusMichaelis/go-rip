package rest

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
)

type HttpHandler func(rest.ResponseWriter, *rest.Request)

func logger(w rest.ResponseWriter, r *rest.Request) {

	log.Printf("Got connection: '%s'", r.Proto)
}

func kevin(w rest.ResponseWriter, r *rest.Request) {

	w.WriteJson(map[string]string{"Booty": "Where is it?!!"})
}

func compose(handlerList []HttpHandler) HttpHandler {

	return func(w rest.ResponseWriter, r *rest.Request) {

		for _, handler := range handlerList {

			handler(w, r)
		}
	}
}

func MakeApi() (api *rest.Api, err error) {

	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(compose([]HttpHandler{logger, kevin})))

	return
}
