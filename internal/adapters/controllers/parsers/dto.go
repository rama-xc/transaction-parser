package prscontroller

import "transaction-parser/internal/usecases/parser-composer/parser"

type ParserExistDto struct {
	ID string `param:"id"`
}

type RunDto struct {
	ID      string       `json:"id" validate:"required"`
	Parsers []*RunDtoOpt `json:"parsers" validate:"required,dive,required"`
}

type RunDtoOpt struct {
	Type    parser.Type `json:"type"  validate:"oneof=history support real-time"`
	Workers int         `json:"workers"`
}
