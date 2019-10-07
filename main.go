package main

import (
	"fmt"

	"log"
	"net/http"
)

func getConf(key string) (value string) {

	config := map[string]string{
		"ip":          "::1",
		"port":        "4343",
		"certificate": "server.crt",
		"key":         "server.key",
	}

	value = config[key]

	return
}

func main() {

	addr := fmt.Sprintf("[%s]:%s", getConf("ip"), getConf("port"))

	srv := &http.Server{Addr: addr, Handler: http.HandlerFunc(handle)}
	log.Printf("Serving on '%s'", addr)
	log.Fatal(srv.ListenAndServeTLS(getConf("certificate"), getConf("key")))
}

func handle(w http.ResponseWriter, r *http.Request) {

	log.Printf("Got connection: '%s'", r.Proto)
	w.Write([]byte("YOLO!!!!!!!"))
}
