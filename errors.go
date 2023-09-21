package bitbucket

import "errors"

var (
	errorBadRequest   = errors.New("400: bad request")
	errorUnauthorized = errors.New("401: unauthorized")
	errorForbidden    = errors.New("403: forbidden")
	errorNotFound     = errors.New("404: not found")
	errorConflict     = errors.New("409: conflict")
)
