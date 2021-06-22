package validation

import (
	"fmt"
	"lupusmic.org/rip/registry"
)

// Inspired by https://evilmartians.com/chronicles/errors-in-go-from-denial-to-acceptance

func New() (err *Errors) {

	err = &Errors{}
	err.Messages = registry.New()
	return
}

type Errors struct {
	Messages registry.Map
}

func (e *Errors) Error() (out string) {

	out = fmt.Sprintf("%d validation errors", len(e.Messages))
	return
}
