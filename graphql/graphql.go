package graphql

import (
	"fmt"
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
        code: String
    }
`

type Country struct {
	Code_ *string
	Id    *graphql.ID
}

func (c Country) ID() *graphql.ID {

	return c.Id
}

func (c Country) Code() *string {

	return c.Code_
}

type query struct{}

func (r *query) Country(args struct{ Code string }) (c *Country, err error) {

	b := business.Business{}
	found := b.GetCountryByCode(args.Code)
	if nil == found {

		err = fmt.Errorf("unknown country code '%s'", args.Code)
		return
	}

	c = &Country{Code_: &found.Code}

	return
}

func MakeEndpoint() (endpoint *relay.Handler, err error) {

	parsedSchema := graphql.MustParseSchema(schema, &query{})
	endpoint = &relay.Handler{Schema: parsedSchema}

	return
}
