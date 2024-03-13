package ethdriver

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func New(url string) (*ethclient.Client, error) {
	op := "eth_driver.New"

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, errors.New(
			fmt.Sprintf("%s: %e", op, err),
		)
	}

	return client, nil
}
