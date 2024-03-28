package prscontroller

import "errors"

var (
	ParserNotFoundErr = errors.New("block not found")
	ValidationErr     = errors.New("validation error")
)
