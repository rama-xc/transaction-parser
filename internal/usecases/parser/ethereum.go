package parser

import (
	"context"
	"fmt"
	"log/slog"
)

type Ethereum struct {
	gateway BlockProvider
}

func (e *Ethereum) GetHistoryParser(
	fromBlk, toBlk int64,
	execs int,
	log *slog.Logger,
) IHistory {
	jobs := make(chan *Job)
	free := make(chan bool)
	stop := make(chan bool)
	comm := make(chan Command)

	var queue []*Job

	for blk := fromBlk; blk <= toBlk; blk++ {
		job := NewJob(fmt.Sprintf("#%d", blk), blk)
		queue = append(queue, job)
	}

	jq := NewJobQueue(
		queue, jobs,
	)

	h := &History{
		comm:     comm,
		jobs:     jobs,
		execFree: free,
		execStop: stop,
		queue:    jq,
		gateway:  e.gateway,
		log:      log,
		ctx:      context.Background(),
	}

	h.createExecs(execs)

	go h.runController()

	readyState := &ReadyState{prsr: h}
	runningState := &RunningState{prsr: h}

	h.setState(readyState)

	h.readyState = readyState
	h.runningState = runningState

	return h
}

type EthereumHistory struct {
	History
}
