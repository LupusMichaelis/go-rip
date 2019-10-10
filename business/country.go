package business

import ()

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

func (b *Business) GetCountryByCode(code string) *Country {

	for _, country := range countryList {

		if code == country.Code {

			return &country
		}
	}

	return nil
}

func (b *Business) GetAllCountries() []Country {

	return countryList
}
