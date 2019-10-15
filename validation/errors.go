package validation

import (
	"fmt"
	"lupusmic.org/rip/util"
)

// Inspired by https://evilmartians.com/chronicles/errors-in-go-from-denial-to-acceptance

func New() (err *Errors) {

	err = &Errors{}
	err.Messages = util.MakeRegistry()
	return
}

type Errors struct {
	Messages util.Registry
}

func (e *Errors) Error() (out string) {

	out = fmt.Sprintf("%d validation errors", len(e.Messages))
	return
}
