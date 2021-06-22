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
		Country{Code: "fr", Name: "France", Population: 67372000},
		Country{Code: "de", Name: "Germany", Population: 82887000},
		Country{Code: "uk", Name: "United Kingdom", Population: 66435600},
		Country{Code: "cn", Name: "China", Population: 1399580000},
		Country{Code: "zz", Name: "Zorg Zo", Population: 239958123456},
	)

	return
}
