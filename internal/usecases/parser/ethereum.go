package parser

import (
	"context"
	"fmt"
	"log/slog"
)

type Ethereum struct {
	gateway BlockProvider
}

func (e *Ethereum) getHistoryParser(
	fromBlk, toBlk int64,
	wrks int,
	log *slog.Logger,
) IHistory {
	jobs := make(chan *Job)
	free := make(chan bool)
	stop := make(chan bool)

	var queue []*Job

	for blk := fromBlk; blk <= toBlk; blk++ {
		job := NewJob(fmt.Sprintf("#%d", blk), blk)
		queue = append(queue, job)
	}

	jq := NewJobQueue(
		queue,
		jobs, free, stop,
	)

	return &History{
		jobs:    jobs,
		free:    free,
		stop:    stop,
		wrks:    wrks,
		log:     log,
		ctx:     context.Background(),
		queue:   jq,
		gateway: e.gateway,
	}
}

type EthereumHistory struct {
	History
}
