package graphql

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"lupusmic.org/rip/business"
)

const schema = `
    schema {
        query: Query
    }

    type Query {
        country(code: String!): Country
    }

    type Country {
        id: ID
        code: String!
    }
`

type Country struct {
	Fields struct {
		Id *graphql.ID
		business.Country
	}
}

func (c Country) ID() *graphql.ID {

	return c.Fields.Id
}

func (c Country) Code() string {

	return c.Fields.Code
}

type query struct{}

func (r *query) Country(args struct{ Code string }) (c *Country, err error) {

	b := business.Business{}
	found, err := b.GetCountryByCode(args.Code)
	if nil == found {

		return
	}

	c = &Country{Fields: struct {
		Id *graphql.ID
		business.Country
	}{
		Id:      nil,
		Country: business.Country{Code: found.Code},
	}}

	return
}

func MakeEndpoint() (endpoint *relay.Handler, err error) {

	parsedSchema := graphql.MustParseSchema(schema, &query{})
	endpoint = &relay.Handler{Schema: parsedSchema}

	return
}
