package parser

import "errors"

var (
	BlockchainNameError   = errors.New("wrong blockchain name")
	DriverConnectionError = errors.New("error occurred while truing connect driver")
)
