// Package kucoinerrors contains errors which may appear when bots work.
package kucoinerrors

import "errors"

var (
	ErrStatusCodeIsNot200 = errors.New("response status code not equal 200")
)

var (
	ErrUnmarshal = errors.New("unmarshal json error")
	ErrRecastDTO = errors.New("recast dto error")
)

var (
	ErrNothingToChange = errors.New("in db no rows have changed")
)
