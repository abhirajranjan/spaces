package db

import "errors"

var ErrCql = errors.New("failed creating record")
var ErrZeroQuery = errors.New("zero queries passed")
