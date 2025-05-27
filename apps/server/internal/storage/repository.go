package repository

import "errors"

var ErrNotFound = errors.New("resource not found")
var ErrInternalServerError = errors.New("error on query execution")
