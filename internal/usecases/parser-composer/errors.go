package prsrcmps

import "errors"

var (
	InvalidParserType = errors.New("parser type doesn't exist")
	NotImplemented    = errors.New("not implemented")
)
