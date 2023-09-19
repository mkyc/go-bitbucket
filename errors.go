package bitbucket

import "errors"

var (
	errorBadRequest = errors.New("400: bad request")
	errorNotFound   = errors.New("404: not found")
)
