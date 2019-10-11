package rest

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"lupusmic.org/rip/business"
	"net/http"
)

func MakeApi() (api *rest.Api, err error) {

	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/country", GetAllCountries),
		rest.Post("/country", PostOneCountry),
		rest.Get("/country/:code", GetOneCountry),
	)

	if err != nil {
		api = nil
		log.Fatal(err)
	}

	api.SetApp(router)

	return
}

type Country struct {
	Code string `json:string`
	Name string `json:string`
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

func PostOneCountry(w rest.ResponseWriter, req *rest.Request) {

	payload := Country{}
	err := req.DecodeJsonPayload(&payload)

	if nil != err {

		rest.Error(w, err.Error(), 400)
		return
	}

	b := business.Business{}
	validationErrorList := b.ValidateCountry(business.Country{
		Code: payload.Code,
		Name: payload.Name,
	})

	if 0 < len(validationErrorList) {

		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(validationErrorList)
		return
	}

	err = b.AddCountry(business.Country{
		Code: payload.Code,
		Name: payload.Name,
	})

	if nil != err {

		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"error": "Couldn't add the country"})
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/country/%s", payload.Code))
	w.WriteHeader(http.StatusCreated)

	return
}
