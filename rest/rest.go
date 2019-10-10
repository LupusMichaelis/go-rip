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

	code := req.PathParam("code")

	b := business.Business{}
	one, err := b.GetCountryByCode(code)

	if nil == one {

		rest.Error(w, err.Error(), 404)
		return
	}

	w.WriteJson(&one)
}
