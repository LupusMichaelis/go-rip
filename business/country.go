package business

import (
	"fmt"
)

type Business struct{}

type Country struct {
	Code string
	Name string
}

type CountryList []Country

var countryList CountryList

func init() {
	countryList = append(countryList,
		Country{Code: "fr", Name: "France"},
		Country{Code: "de", Name: "Germany"},
		Country{Code: "en", Name: "England"},
	)
}

func (b *Business) GetCountryByCode(code string) (country *Country, err error) {

	for _, c := range countryList {

		if code == c.Code {

			country = &c
			return
		}
	}

	err = fmt.Errorf("unknown country code '%s'", code)

	return
}

func (b *Business) GetAllCountries() []Country {

	return countryList
}
