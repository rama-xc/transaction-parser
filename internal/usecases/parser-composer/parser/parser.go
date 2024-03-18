package parser

import (
	"context"
	"transaction-parser/internal/entity"
)

type BlockProvider interface {
	Block(ctx context.Context, number int64) (*entity.Block, error)
	LastBlockNum(ctx context.Context) (int64, error)
}

type Type string

const (
	History  Type = "history"
	Support  Type = "support"
	RealTime Type = "real-time"
)

type IParser interface {
	Start()
	SetWorkers(wrks int)
}

type Parser struct {
	jobs chan *Job
	free chan bool
	stop chan bool

	wrks int

	queue *JobQueue

	gateway BlockProvider
}

func (p *Parser) Start() {

}

func (p *Parser) SetWorkers(wrks int) {
	p.wrks = wrks
}

type HistoryParser struct {
	Parser
}

func NewHistoryParser(
	gateway BlockProvider,
) *HistoryParser {

	return &HistoryParser{
		Parser: Parser{
			gateway: gateway,
		},
	}

}

type RealTimeParser struct {
	Parser
}

func NewRealTimeParser(
	gateway BlockProvider,
) IParser {

	return &RealTimeParser{
		Parser: Parser{
			gateway: gateway,
		},
	}

}

type SupportParser struct {
	Parser
}

//func NewSupportParser(
//	wrks int,
//	gateway prschain.BlockProvider,
//) *SupportParser {
//
//	return &SupportParser{
//		Parser: Parser{
//			wrks:    wrks,
//			gateway: gateway,
//		},
//	}
//
//}
