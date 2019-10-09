package main

import (
	"fmt"

	"log"
	"net/http"

	"./config"
	"./rest"
)

func main() {

	restApi, err := rest.MakeApi()
	if nil != err {

		log.Fatal(err)
	}

	apiHandler := restApi.MakeHandler()
	http.Handle("/rest", apiHandler)

	addr := fmt.Sprintf("[%s]:%d", config.GetConfiguration().Ip, config.GetConfiguration().Port)
	srv := &http.Server{
		Addr: addr,
	}

	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(config.GetConfiguration().Certificate, config.GetConfiguration().Key))
}
