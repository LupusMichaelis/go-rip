package main

import (
	"fmt"

	"log"
	"net/http"

	"./config"
	"./graphql"
	"./rest"
	"lupusmic.org/rip/business"
)

func main() {

	var b *business.Business = business.New()

	restApi, err := rest.MakeApi(b)
	if nil != err {

		log.Fatal(err)
	}

	apiHandler := restApi.MakeHandler()
	http.Handle("/rest/", http.StripPrefix("/rest", apiHandler))

	graphqlHandler, err := graphql.MakeEndpoint(b)
	if nil != err {

		log.Fatal(err)
	}

	http.Handle("/graph", graphqlHandler)

	addr := fmt.Sprintf("[%s]:%d", config.GetConfiguration().Ip, config.GetConfiguration().Port)
	srv := &http.Server{
		Addr: addr,
	}

	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(config.GetConfiguration().Certificate, config.GetConfiguration().Key))
}
