package kucoinerrors

import "errors"

var (
	StatusCodeIsNot200 = errors.New("response status code not equal 200")
)
