package main

import (
	"fmt"

	"log"
	"net/http"

	"./config"
	"./graphql"
	"./rest"
)

func main() {

	restApi, err := rest.MakeApi()
	if nil != err {

		log.Fatal(err)
	}

	apiHandler := restApi.MakeHandler()
	http.Handle("/rest", apiHandler)

	graphqlHandler, err := graphql.MakeEndpoint()
	if nil != err {

		log.Fatal(err)
	}

	http.Handle("/graph", graphqlHandler) // XXX

	addr := fmt.Sprintf("[%s]:%d", config.GetConfiguration().Ip, config.GetConfiguration().Port)
	srv := &http.Server{
		Addr: addr,
	}

	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(config.GetConfiguration().Certificate, config.GetConfiguration().Key))
}
