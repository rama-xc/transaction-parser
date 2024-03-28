package blkgtw

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"transaction-parser/internal/entity"
)

type Gateway struct {
	client *ethclient.Client
}

func New(client *ethclient.Client) *Gateway {
	return &Gateway{
		client: client,
	}
}

func (g *Gateway) Block(ctx context.Context, number int64) (*entity.Block, error) {
	block, err := g.client.BlockByNumber(
		ctx,
		big.NewInt(number),
	)
	if err != nil {
		return nil, err
	}
	chainID, err := g.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	var transactions []*entity.Transaction

	for i, tx := range block.Transactions() {
		from, _ := types.Sender(types.NewLondonSigner(chainID), tx)

		var to string

		if tx.To() != nil {
			to = tx.To().Hex()
		}

		transactions = append(transactions, &entity.Transaction{
			Hash:             tx.Hash().Hex(),
			To:               to,
			From:             from.Hex(),
			Value:            tx.Value().String(),
			Gas:              tx.Gas(),
			GasPrice:         tx.GasPrice().Uint64(),
			Nonce:            tx.Nonce(),
			TransactionIndex: uint16(i),
		})
	}

	return &entity.Block{
		Difficulty:   block.Difficulty().Uint64(),
		Number:       block.NumberU64(),
		Time:         block.Time(),
		Hash:         block.Hash().Hex(),
		Transactions: transactions,
	}, nil
}

func (g *Gateway) LastBlockNum(ctx context.Context) (int64, error) {
	header, err := g.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return -1, err
	}

	return header.Number.Int64(), nil
}
