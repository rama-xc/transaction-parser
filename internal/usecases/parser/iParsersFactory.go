package parser

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	ethgateway "transaction-parser/internal/adapters/gateways/block"
	blkrepo "transaction-parser/internal/adapters/repositories/block"
	"transaction-parser/internal/app/common/config"
	ethdriver "transaction-parser/internal/app/drivers/eth"
)

type IParsersFactory interface {
	GetHistoryParser(fromBlk, toBlk int64, execs int, log *slog.Logger) IHistory
}

func GetParsersFactory(
	blockchain config.BlockchainName,
	providerUrl string,
	redisDrv *redis.Client,
) (IParsersFactory, error) {
	switch blockchain {
	case config.Ethereum:
		ethDrv, err := ethdriver.New(providerUrl)
		if err != nil {
			return nil, DriverConnectionError
		}

		blkprovider := ethgateway.New(ethDrv)
		blkcaching := blkrepo.NewRedisRepo(redisDrv)

		return &Ethereum{blkProvider: blkprovider, blkCaching: blkcaching}, nil
	default:
		return nil, BlockchainNameError
	}
}

type IParserBase interface {
	Listen()
}

func MustLoad(cfg []config.ParsersFactoriesConfig, log *slog.Logger, redisDrv *redis.Client) map[string]IHistory {
	prs := map[string]IHistory{}

	for _, factoryCfg := range cfg {
		factory, err := GetParsersFactory(factoryCfg.Blockchain, factoryCfg.ProviderUrl, redisDrv)
		if err != nil {
			panic(err)
		}

		hCfg := factoryCfg.Parsers.History

		prs[fmt.Sprintf("%s:history", factoryCfg.ID)] =
			factory.GetHistoryParser(hCfg.BlockFrom, hCfg.BlockTo, hCfg.Workers, log)
	}

	return prs
}
