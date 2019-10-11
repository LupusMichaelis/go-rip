package business

import (
	"fmt"
	"sync"
)

type Business struct{}

type Country struct {
	Code string
	Name string
}

type CountryList []Country

var lock = sync.RWMutex{}
var countryList CountryList

func init() {
	lock.Lock()
	defer lock.Unlock()

	countryList = append(countryList,
		Country{Code: "fr", Name: "France"},
		Country{Code: "de", Name: "Germany"},
		Country{Code: "en", Name: "England"},
	)
}

func (b *Business) GetCountryByCode(code string) (country *Country, err error) {

	lock.RLock()
	defer lock.RUnlock()

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

func (b *Business) ValidateCountry(country Country) (err map[string]string) {

	err = make(map[string]string, 0)

	if 2 != len(country.Code) {

		err["code"] = fmt.Sprintf("Country code '%s' must be a 2 character string", country.Code)
	}

	if 0 == len(country.Name) {

		err["name"] = "Country name must not be empty"
	}

	if 0 == len(err) {

		err = nil
	}

	return
}

func (b *Business) AddCountry(country Country) (err error) {

	lock.Lock()
	defer lock.Unlock()

	countryList = append(countryList, country)

	return
}
