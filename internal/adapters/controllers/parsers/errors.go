package prscontroller

import "errors"

var (
	ParserNotFoundErr = errors.New("parser not found")
	ValidationErr     = errors.New("validation error")
)
