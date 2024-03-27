package parser

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type Ethereum struct {
	gateway BlockProvider
}

func (e *Ethereum) GetHistoryParser(
	fromBlk, toBlk int64,
	execsAmount int,
	log *slog.Logger,
	redis *redis.Client,
) IHistory {
	comm := make(chan Command)

	var queue []*Job

	for blk := fromBlk; blk <= toBlk; blk++ {
		job := NewJob(fmt.Sprintf("#%d", blk), blk)
		queue = append(queue, job)
	}

	h := &History{
		comm:        comm,
		log:         log,
		execManager: NewExecutorManager(queue, execsAmount, log, e.gateway, redis),
	}

	go h.runController()

	return h
}

type EthereumHistory struct {
	History
}
