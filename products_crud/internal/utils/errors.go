package utils

import "errors"

var (
	// ErrRequestContentTypeNotJSON is used when the request content type is not application/json.
	ErrRequestContentTypeNotJSON = errors.New("request content type is not application/json")
	// ErrRequestJSONInvalid is used when the request json is invalid.
	ErrRequestJSONInvalid = errors.New("request json invalid")
)
