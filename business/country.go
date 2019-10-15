package business

import (
	"fmt"
	"lupusmic.org/rip/validation"
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

func (b *Business) ValidateCountry(country Country) (err *validation.Errors) {

	err = validation.New()

	if 2 != len(country.Code) {

		if 0 < len(err.Messages["code"]) {

			err.Messages["code"] = make([]string, 1)
		}

		err.Messages["code"] = append(err.Messages["code"], fmt.Sprintf("Country code '%s' must be a 2 character string", country.Code))
	}

	if 0 == len(country.Name) {

		if 0 < len(err.Messages["name"]) {

			err.Messages["name"] = make([]string, 1)
		}

		err.Messages["name"] = append(err.Messages["name"], "Country name must not be empty")
	}

	if 0 == len(err.Messages) {

		err = nil
	}

	return
}

func (b *Business) AddCountry(country Country) (err *validation.Errors) {

	lock.Lock()
	defer lock.Unlock()

	err = b.ValidateCountry(Country{
		Code: country.Code,
		Name: country.Name,
	})

	if nil != err {

		return
	}

	countryList = append(countryList, country)

	return
}

func (b *Business) UpdateCountry(newValue Country) (err error) {

	currentValue, err := b.GetCountryByCode(newValue.Code)

	if nil != err {

		return
	}

	currentValue.Name = newValue.Name
	return
}

func (b *Business) DeleteCountry(deleteMe Country) (err error) {

	lock.Lock()
	defer lock.Unlock()

	var (
		position int
		current  Country
	)

	for position, current = range countryList {

		if current.Code == deleteMe.Code {

			break
		}
	}

	if len(countryList) == position {

		return fmt.Errorf("Not found")
	}

	countryList = append(countryList[:position], countryList[position+1:]...)

	return
}
