package business

import (
	"sync"
)

type Business struct {
	lock        sync.RWMutex
	countryList CountryList
}

func (b *Business) Lock() {
	b.lock.Lock()
}

func (b *Business) Unlock() {
	b.lock.Unlock()
}

func New() (b *Business) {

	b = &Business{}
	b.lock = sync.RWMutex{}

	b.Lock()
	defer b.Unlock()

	b.countryList = append(b.countryList,
		Country{Code: "fr", Name: "France"},
		Country{Code: "de", Name: "Germany"},
		Country{Code: "en", Name: "England"},
	)

	return
}
