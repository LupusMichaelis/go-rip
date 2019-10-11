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
		rest.Put("/country/:code", PutOneCountry),
		rest.Delete("/country/:code", deleteOneCountry),
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

/*
func (c * Country) hydrate(from business.Country) {

    c.Code = from.Code
    c.Name = from.Name
}

type CountryList []Country

*/
func GetAllCountries(out rest.ResponseWriter, in *rest.Request) {

	b := business.Business{}
	all := b.GetAllCountries()
	out.WriteJson(&all)
}

func GetOneCountry(out rest.ResponseWriter, in *rest.Request) {

	code := in.PathParam("code")

	b := business.Business{}
	one, err := b.GetCountryByCode(code)

	if nil == one {

		rest.Error(out, err.Error(), http.StatusNotFound)
		return
	}

	out.WriteJson(&one)
}

/*
	var payload []Country = make([]Country, len(all))
    for index, from := range all {

        payload[index].hydrate(from)
    }

    out.WriteJson(&payload)
}
*/

func PostOneCountry(out rest.ResponseWriter, in *rest.Request) {

	payload := Country{}
	err := in.DecodeJsonPayload(&payload)

	if nil != err {

		rest.Error(out, err.Error(), 400)
		return
	}

	b := business.Business{}
	validationErrorList := b.ValidateCountry(business.Country{
		Code: payload.Code,
		Name: payload.Name,
	})

	if 0 < len(validationErrorList) {

		out.WriteHeader(http.StatusBadRequest)
		out.WriteJson(validationErrorList)
		return
	}

	err = b.AddCountry(business.Country{
		Code: payload.Code,
		Name: payload.Name,
	})

	if nil != err {

		out.WriteHeader(http.StatusInternalServerError)
		out.WriteJson(map[string]string{"error": "Couldn't add the country"})
		return
	}

	out.Header().Set("Location", fmt.Sprintf("/country/%s", payload.Code))
	out.WriteHeader(http.StatusCreated)

	return
}

func PutOneCountry(out rest.ResponseWriter, in *rest.Request) {

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
