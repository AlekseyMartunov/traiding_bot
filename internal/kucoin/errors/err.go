// Package kucoinerrors contains errors which may appear when bots work.
package kucoinerrors

import "errors"

var (
	StatusCodeIsNot200 = errors.New("response status code not equal 200")
)

var (
	NothingToChange = errors.New("in db no rows have changed")
)
