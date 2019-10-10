package rest

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"lupusmic.org/rip/business"
)

func MakeApi() (api *rest.Api, err error) {

	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/country", GetAllCountries),
		rest.Get("/country/:code", GetOneCountry),
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

func GetOneCountry(w rest.ResponseWriter, req *rest.Request) {

	b := business.Business{}
	code := req.PathParam("code")
	one := b.GetCountryByCode(code)

	if nil != one {

		w.WriteJson(&one)
	} else {

		rest.Error(w, fmt.Sprintf("unknown country code '%s'", code), 404)
	}
}
