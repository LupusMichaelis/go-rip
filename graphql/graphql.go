package graphql

import (
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"lupusmic.org/rip/business"
	"math"
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

func (c Country) Population() (population *int32, err error) {

	if uint(math.MaxInt32) < c.Fields.Population {

		err = fmt.Errorf("Population overflow '%d'", c.Fields.Population)

	} else {

		value := int32(c.Fields.Population)
		population = &value
	}

	return
}

type query struct {
	b *business.Business
}

func (q *query) Country(args struct{ Code string }) (c *Country, err error) {

	found, err := q.b.GetCountryByCode(args.Code)
	if nil == found {

		return
	}

	c = &Country{Fields: struct {
		business.Country
	}{
		Country: business.Country{
			Code:       found.Code,
			Name:       found.Name,
			Population: found.Population,
		},
	}}

	return
}

func MakeEndpoint(b *business.Business) (endpoint *relay.Handler, err error) {

	schema, err := ioutil.ReadFile("graphql/schema.graphql")

	if err != nil {

		log.Fatal(err)
	}

	parsedSchema := graphql.MustParseSchema(string(schema), &query{b})
	endpoint = &relay.Handler{Schema: parsedSchema}

	return
}

func (q *query) Add(args struct {
	Code string
	Name string
}) (c *Country, err error) {

	validation := q.b.AddCountry(business.Country{
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
