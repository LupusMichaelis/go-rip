package main

import (
	"fmt"

	"log"
	"net/http"

	"./config"
	"./rest"
)

func main() {

	restApi, err := rest.MakeRouter()
	if nil != err {

		log.Fatal(err)
	}

	apiHandler := restApi.MakeHandler()

	addr := fmt.Sprintf("[%s]:%d", config.GetConfiguration().Ip, config.GetConfiguration().Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: apiHandler,
	}

	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(config.GetConfiguration().Certificate, config.GetConfiguration().Key))
}
