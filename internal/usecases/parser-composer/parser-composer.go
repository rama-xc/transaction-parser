package prsrcmps

import (
	"context"
	"log/slog"
	"transaction-parser/internal/entity"
	"transaction-parser/internal/usecases/parser-composer/parser"
)

type BlockProvider interface {
	Block(ctx context.Context, number int64) (*entity.Block, error)
	LastBlockNum(ctx context.Context) (int64, error)
}

type IParserComposer interface {
	ID() string
	run()
	stop()
}

type ParserComposer struct {
	id      string
	gateway BlockProvider
	log     *slog.Logger

	ready   State
	running State
	stopped State

	currentState State

	parsers map[parser.Type]int
}

func (p *ParserComposer) setState(s State) {
	p.currentState = s
}

func (p *ParserComposer) run() {
	p.currentState.run()
}

func (p *ParserComposer) stop() {
	p.currentState.stop()
}

func (p *ParserComposer) ID() string {
	return p.id
}

func (p *ParserComposer) createParser(tp parser.Type) (parser.IParser, error) {
	switch tp {
	case parser.History:
		return parser.NewHistoryParser(p.gateway), nil
	default:
		return nil, InvalidParserType
	}
}
