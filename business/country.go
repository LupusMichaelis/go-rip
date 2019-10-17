package business

import (
	"fmt"
	"lupusmic.org/rip/validation"
)

type Country struct {
	Code       string
	Name       string
	Population uint
}

type CountryList []Country

func (b *Business) GetCountryByCode(code string) (country *Country, err error) {

	for _, c := range b.countryList {

		if code == c.Code {

			country = &c
			return
		}
	}

	err = fmt.Errorf("unknown country code '%s'", code)

	return
}

func (b *Business) GetAllCountries() []Country {

	return b.countryList
}

func (b *Business) ValidateCountry(country Country) (err *validation.Errors) {

	err = validation.New()

	if 2 != len(country.Code) {

		err.Messages.Add("code", fmt.Sprintf("Country code '%s' must be a 2 character string", country.Code))
	}

	if c, _ := b.GetCountryByCode(country.Code); nil != c {

		err.Messages.Add("code", fmt.Sprintf("Country code '%s' must be unique", country.Code))
	}

	if 0 == len(country.Name) {

		err.Messages.Add("name", "Country name must not be empty")
	}

	if 0 == len(err.Messages) {

		err = nil
	}

	return
}

func (b *Business) AddCountry(country Country) (err *validation.Errors) {

	err = b.ValidateCountry(Country{
		Code: country.Code,
		Name: country.Name,
	})

	if nil != err {

		return
	}

	b.countryList = append(b.countryList, country)

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

	var (
		position int
		current  Country
	)

	for position, current = range b.countryList {

		if current.Code == deleteMe.Code {

			break
		}
	}

	if len(b.countryList) == position {

		return fmt.Errorf("Not found")
	}

	b.countryList = append(b.countryList[:position], b.countryList[position+1:]...)

	return
}
