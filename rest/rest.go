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
		rest.Get("/country", getAllCountries),

		rest.Delete("/country/:code", deleteOneCountry),
		rest.Get("/country/:code", getOneCountry),
		rest.Post("/country", postOneCountry),
		rest.Put("/country/:code", putOneCountry),
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

func getAllCountries(out rest.ResponseWriter, in *rest.Request) {

	b := business.Business{}
	all := b.GetAllCountries()
	out.WriteJson(&all)
}

func getOneCountry(out rest.ResponseWriter, in *rest.Request) {

	code := in.PathParam("code")

	b := business.Business{}
	one, err := b.GetCountryByCode(code)

	if nil == one {

		rest.Error(out, err.Error(), http.StatusNotFound)
		return
	}

	out.WriteJson(&one)
}

func postOneCountry(out rest.ResponseWriter, in *rest.Request) {

	payload := Country{}
	err := in.DecodeJsonPayload(&payload)

	if nil != err {

		rest.Error(out, err.Error(), http.StatusNotFound)
		return
	}

	b := business.Business{}
	validation := b.AddCountry(business.Country{
		Code: payload.Code,
		Name: payload.Name,
	})

	if nil != validation {

		out.WriteHeader(http.StatusBadRequest)
		out.WriteJson(err)
		return
	}

	if nil != err {

		out.WriteHeader(http.StatusInternalServerError)
		out.WriteJson(map[string]string{"error": "Couldn't add the country"})
		return
	}

	out.Header().Set("Location", fmt.Sprintf("/country/%s", payload.Code))
	out.WriteHeader(http.StatusCreated)

	return
}

func putOneCountry(out rest.ResponseWriter, in *rest.Request) {

	payload := Country{}
	err := in.DecodeJsonPayload(&payload)

	if nil != err {

		rest.Error(out, err.Error(), http.StatusBadRequest)
		return
	}

	b := business.Business{}
	one, err := b.GetCountryByCode(payload.Code)
	if nil != err {

		rest.Error(out, err.Error(), http.StatusNotFound)
		return
	}

	if one.Code != payload.Code && one.Name != payload.Name {

		one.Name = payload.Name
		err = b.UpdateCountry(*one)
	}

	if nil != err {

		rest.Error(out, err.Error(), http.StatusNotFound)
		return
	}

	out.Header().Set("Location", fmt.Sprintf("/country/%s", payload.Code))
	out.WriteHeader(http.StatusNoContent)

	return
}

func deleteOneCountry(out rest.ResponseWriter, in *rest.Request) {

	code := in.PathParam("code")

	b := business.Business{}
	one, err := b.GetCountryByCode(code)

	if nil == one {

		rest.Error(out, err.Error(), http.StatusNotFound)
		return
	}

	err = b.DeleteCountry(*one)
	if nil != err {

		rest.Error(out, err.Error(), http.StatusInternalServerError)
		return
	}

	out.WriteHeader(http.StatusNoContent)
}
