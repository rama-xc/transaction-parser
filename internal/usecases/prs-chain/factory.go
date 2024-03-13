package prs_chain

import (
	"context"
	"fmt"
	ethgateway "transaction-parser/internal/adapters/gateways/eth"
	"transaction-parser/internal/app/common/config"
	ethdriver "transaction-parser/internal/app/drivers/eth"
	"transaction-parser/internal/entity"
)

type IWeb3Gateway interface {
	Block(ctx context.Context, number int64) (*entity.Block, error)
	LastBlockNum(ctx context.Context) (int64, error)
}

type IParser interface {
}

type Parser struct {
	gateway IWeb3Gateway
}

type ETHParser struct {
	Parser
}

func NewETHParser(gateway IWeb3Gateway) *ETHParser {
	return &ETHParser{Parser: Parser{gateway: gateway}}
}

func MustCreateParser(psrCfg config.ParserConfig) IParser {
	switch psrCfg.Blockchain {
	case config.Ethereum:
		drv, err := ethdriver.New(psrCfg.ProviderUrl)
		if err != nil {
			panic(fmt.Sprintf("can't create parser #%s. cause: %s", psrCfg.ID, err.Error()))
		}

		gtw := ethgateway.New(drv)

		return NewETHParser(gtw)
	default:
		panic("not allowed blockchain.")
	}
}
