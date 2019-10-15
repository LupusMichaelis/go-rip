package validation

import (
	"fmt"
)

// Inspired by https://evilmartians.com/chronicles/errors-in-go-from-denial-to-acceptance

func New() (err *Errors) {

	err = &Errors{}
	err.Messages = make(map[string][]string, 0)
	return
}

type Errors struct {
	Messages map[string][]string
}

func (e *Errors) Error() (out string) {

	out = fmt.Sprintf("%d validation errors", len(e.Messages))
	return
}
