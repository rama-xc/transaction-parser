package parser

import (
	"context"
	"fmt"
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
	GetHistoryParser(fromBlk, toBlk int64, wrks int, log *slog.Logger) IHistory
}

func GetParsersFactory(
	blockchain config.BlockchainName,
	providerUrl string,
) (IParsersFactory, error) {
	switch blockchain {
	case config.Ethereum:
		drv, err := ethdriver.New(providerUrl)
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

func MustLoad(cfg []config.ParsersFactoriesConfig, log *slog.Logger) map[string]IHistory {
	prs := map[string]IHistory{}

	for _, factoryCfg := range cfg {
		factory, err := GetParsersFactory(factoryCfg.Blockchain, factoryCfg.ProviderUrl)
		if err != nil {
			panic(err)
		}

		hCfg := factoryCfg.Parsers.History

		prs[fmt.Sprintf("%s:history", factoryCfg.ID)] =
			factory.GetHistoryParser(hCfg.BlockFrom, hCfg.BlockTo, hCfg.Workers, log)
	}

	return prs
}
