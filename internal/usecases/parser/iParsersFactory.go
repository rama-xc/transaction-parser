package parser

import (
	"context"
	"log/slog"
	ethgateway "transaction-parser/internal/adapters/gateways/eth"
	"transaction-parser/internal/app/common/config"
	ethdriver "transaction-parser/internal/app/drivers/eth"
	"transaction-parser/internal/entity"
)

type BlockProvider interface {
	Block(ctx context.Context, number int64) (*entity.Block, error)
	LastBlockNum(ctx context.Context) (int64, error)
}

type IParsersFactory interface {
	getHistoryParser(fromBlk, toBlk int64, wrks int, log *slog.Logger) IHistory
}

func GetParsersFactory(cfg config.ParserConfig) (IParsersFactory, error) {
	switch cfg.Blockchain {
	case config.Ethereum:
		drv, err := ethdriver.New(cfg.ProviderUrl)
		if err != nil {
			return nil, DriverConnectionError
		}

		gtw := ethgateway.New(drv)

		return &Ethereum{gateway: gtw}, nil
	default:
		return nil, BlockchainNameError
	}
}

type IParserBase interface {
	Listen()
}
