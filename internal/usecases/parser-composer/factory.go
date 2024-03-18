package prsrcmps

import (
	"fmt"
	ethgateway "transaction-parser/internal/adapters/gateways/eth"
	"transaction-parser/internal/app/common/config"
	ethdriver "transaction-parser/internal/app/drivers/eth"
)

type ETHParserComposer struct {
	ParserComposer
}

func superParserComposer(id string, gateway BlockProvider) ParserComposer {
	p := &ParserComposer{
		id:      id,
		gateway: gateway,
	}

	readyState := &ReadyState{parserComposer: p}
	runningState := &RunningState{parserComposer: p}
	stoppedState := &StoppedState{parserComposer: p}

	p.setState(readyState)

	p.ready = readyState
	p.running = runningState
	p.stopped = stoppedState

	return *p
}

func NewETHParserComposer(id string, gateway BlockProvider) *ETHParserComposer {
	return &ETHParserComposer{ParserComposer: superParserComposer(
		id,
		gateway,
	)}
}

func mustCreateParserComposer(psrCfg config.ParserConfig) IParserComposer {
	switch psrCfg.Blockchain {
	case config.Ethereum:
		drv, err := ethdriver.New(psrCfg.ProviderUrl)
		if err != nil {
			panic(fmt.Sprintf("can't create parser #%s. cause: %s", psrCfg.ID, err.Error()))
		}

		gtw := ethgateway.New(drv)

		return NewETHParserComposer(psrCfg.ID, gtw)
	default:
		panic("not allowed blockchain.")
	}
}

func MustLoadParsers(cfg []config.ParserConfig) map[string]IParserComposer {
	parsers := make(map[string]IParserComposer)

	for _, prsCfg := range cfg {
		parser := mustCreateParserComposer(prsCfg)
		parsers[parser.ID()] = parser
	}

	return parsers
}
