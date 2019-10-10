package rest

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"lupusmic.org/rip/business"
)

func MakeApi() (api *rest.Api, err error) {

	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/country", GetAllCountries),
	)

	if err != nil {
		api = nil
		log.Fatal(err)
	}

	api.SetApp(router)

	return
}

func GetAllCountries(w rest.ResponseWriter, req *rest.Request) {

	b := business.Business{}
	all := b.GetAllCountries()
	w.WriteJson(&all)
}
