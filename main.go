package main

import (
	"fmt"

	"log"
	"net/http"

	"lupusmic.org/rip/business"
	"lupusmic.org/rip/config"
	"lupusmic.org/rip/graphql"
	"lupusmic.org/rip/rest"
)

func main() {

	var b *business.Business = business.New()

	restApi, err := rest.MakeApi(b)
	if nil != err {

		log.Fatal(err)
	}

	apiHandler := restApi.MakeHandler()
	http.Handle("/r/", http.StripPrefix("/r", apiHandler))

	graphqlHandler, err := graphql.MakeEndpoint(b)
	if nil != err {

		log.Fatal(err)
	}

	http.Handle("/g", graphqlHandler)

	addr := fmt.Sprintf("[%s]:%d", config.GetConfiguration().Ip, config.GetConfiguration().Port)
	srv := &http.Server{
		Addr: addr,
	}

	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(config.GetConfiguration().Certificate, config.GetConfiguration().Key))
}
