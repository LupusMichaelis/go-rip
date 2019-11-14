package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"lupusmic.org/rip/business"
	"lupusmic.org/rip/config"
	"lupusmic.org/rip/graphql"
	"lupusmic.org/rip/rest"
)

func main() {

	// Configuration
	configfile := flag.String("config", "config.json", "configuration file")
	flag.Parse()

	fmt.Println(*configfile)
	err := config.Load(*configfile)
	if nil != err {

		log.Fatal(err)
	}

	// Define the model
	var b *business.Business = business.New()

	// Mount REST API
	restApi, err := rest.MakeApi(b)
	if nil != err {

		log.Fatal(err)
	}

	apiHandler := restApi.MakeHandler()
	http.Handle("/r/", http.StripPrefix("/r", apiHandler))

	// Mount GraphQL API
	graphqlHandler, err := graphql.MakeEndpoint(b)
	if nil != err {

		log.Fatal(err)
	}

	http.Handle("/g", graphqlHandler)

	// Launch HTTP server
	addr := fmt.Sprintf("[%s]:%d", config.GetConfiguration().Ip, config.GetConfiguration().Port)
	srv := &http.Server{
		Addr: addr,
	}

	log.Printf("Serving on '%s'", addr)
	err = srv.ListenAndServeTLS(
		config.GetConfiguration().Certificate,
		config.GetConfiguration().Key,
	)
	log.Fatal(err)
}
