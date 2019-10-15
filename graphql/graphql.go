package graphql

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"lupusmic.org/rip/business"
)

type Country struct {
	Fields struct {
		business.Country
	}
}

func (c Country) Code() string {

	return c.Fields.Code
}

func (c Country) Name() *string {

	name := c.Fields.Name
	return &name
}

type query struct{}

func (r *query) Country(args struct{ Code string }) (c *Country, err error) {

	b := business.Business{}
	found, err := b.GetCountryByCode(args.Code)
	if nil == found {

		return
	}

	c = &Country{Fields: struct {
		business.Country
	}{
		Country: business.Country{
			Code: found.Code,
			Name: found.Name,
		},
	}}

	return
}

func MakeEndpoint() (endpoint *relay.Handler, err error) {

	schema, err := ioutil.ReadFile("graphql/schema.graphql")

	if err != nil {

		log.Fatal(err)
	}

	parsedSchema := graphql.MustParseSchema(string(schema), &query{})
	endpoint = &relay.Handler{Schema: parsedSchema}

	return
}

func (q *query) Add(args struct {
	Code string
	Name string
}) (c *Country, err error) {

	b := business.Business{}
	validation := b.AddCountry(business.Country{
		Code: args.Code,
		Name: args.Name,
	})

	if nil != validation {

		err = validation
		return
	}

	c, err = q.Country(struct{ Code string }{args.Code})

	return
}
