package rest

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ant0ine/go-json-rest/rest"

	"lupusmic.org/rip/business"
)

func Contains(haystack []string, needle string) (found bool) {

	for _, straw := range haystack {

		if straw == needle {

			found = true
			break
		}
	}

	return
}

func MakeApi(b *business.Business) (api *rest.Api, err error) {

	api = rest.NewApi()

	api.Use(rest.DefaultDevStack...)

	allowedOrigins := []string{
		"https://172.17.0.3",
	}

	api.Use(&rest.CorsMiddleware{
		OriginValidator: func(
			origin string,
			request *rest.Request,
		) bool {
			return Contains(allowedOrigins, origin)
		},
		RejectNonCorsRequests: false,
		AllowedMethods:        []string{"OPTIONS", "GET"},
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
			"Origin",
		},
		AccessControlAllowCredentials: true,
	})

	router, err := rest.MakeRouter(
		rest.Options("/country", corsOptions(allowedOrigins, []string{"GET", "POST"})),
		rest.Options("/country/:code", corsOptions(allowedOrigins, []string{"DELETE", "GET", "POST", "PUT"})),

		rest.Get("/country", lockBusiness(b, getAllCountries)),
		rest.Delete("/country/:code", lockBusiness(b, deleteOneCountry)),
		rest.Get("/country/:code", lockBusiness(b, getOneCountry)),
		rest.Post("/country", lockBusiness(b, postOneCountry)),
		rest.Put("/country/:code", lockBusiness(b, putOneCountry)),
	)

	if err != nil {
		api = nil
		log.Fatal(err)
	}

	api.SetApp(router)

	return
}

func corsOptions(
	allowedOrigins []string,
	allowedMethods []string,
) (wrapped rest.HandlerFunc) {

	wrapped = func(out rest.ResponseWriter, in *rest.Request) {

		out.Header().Set("Access-Control-Allow-Origin", strings.Join(allowedOrigins, ","))
		out.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ","))
		out.WriteHeader(http.StatusOK)
	}

	return
}

type Country struct {
	Code       string `json:string`
	Name       string `json:string`
	Population uint64 `json:integer`
}

type businessHandlerFunc func(b *business.Business) (h rest.HandlerFunc)

func lockBusiness(
	b *business.Business,
	wrapped businessHandlerFunc,
) (wrap rest.HandlerFunc) {

	allowedOrigins := []string{
		"https://172.17.0.3",
	}

	wrap = func(out rest.ResponseWriter, in *rest.Request) {
		b.Lock()
		defer b.Unlock()

		out.Header().Set("Access-Control-Allow-Origin", strings.Join(allowedOrigins, ","))
		wrapped(b)(out, in)
	}

	return
}

func getAllCountries(b *business.Business) (h rest.HandlerFunc) {

	h = func(out rest.ResponseWriter, in *rest.Request) {

		all := b.GetAllCountries()
		out.WriteJson(&all)
	}

	return
}

func getOneCountry(b *business.Business) (h rest.HandlerFunc) {

	h = func(out rest.ResponseWriter, in *rest.Request) {

		code := in.PathParam("code")

		one, err := b.GetCountryByCode(code)

		if nil == one {

			rest.Error(out, err.Error(), http.StatusNotFound)
			return
		}

		out.WriteJson(&one)
	}

	return
}

func postOneCountry(b *business.Business) (h rest.HandlerFunc) {

	h = func(out rest.ResponseWriter, in *rest.Request) {

		payload := Country{}
		err := in.DecodeJsonPayload(&payload)

		if nil != err {

			rest.Error(out, err.Error(), http.StatusNotFound)
			return
		}

		validation := b.AddCountry(business.Country{
			Code:       payload.Code,
			Name:       payload.Name,
			Population: payload.Population,
		})

		if nil != validation {

			out.WriteHeader(http.StatusBadRequest)
			out.WriteJson(validation.Messages)
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

	return
}

func putOneCountry(b *business.Business) (h rest.HandlerFunc) {

	h = func(out rest.ResponseWriter, in *rest.Request) {

		payload := Country{}
		err := in.DecodeJsonPayload(&payload)

		if nil != err {

			rest.Error(out, err.Error(), http.StatusBadRequest)
			return
		}

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

	return
}

func deleteOneCountry(b *business.Business) (h rest.HandlerFunc) {

	h = func(out rest.ResponseWriter, in *rest.Request) {

		code := in.PathParam("code")
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
		return
	}

	return
}
