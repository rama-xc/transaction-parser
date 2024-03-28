package parser

import (
	"fmt"
	"log/slog"
)

type Ethereum struct {
	blkProvider BlockProvider
	blkCaching  BlockCaching
}

func (e *Ethereum) GetHistoryParser(
	fromBlk, toBlk int64,
	execsAmount int,
	log *slog.Logger,
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
		execManager: NewExecutorManager(queue, execsAmount, log, e.blkProvider, e.blkCaching),
	}

	go h.runController()

	return h
}

type EthereumHistory struct {
	History
}
