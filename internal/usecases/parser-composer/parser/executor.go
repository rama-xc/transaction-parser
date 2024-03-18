package parser

import (
	"context"
	"fmt"
	"log/slog"
)

type JobHandler interface {
	execute(job *Job)
	setNext(handler JobHandler)
}

type JobExecutor struct {
	jobs chan *Job
	free chan bool
	stop chan bool

	log *slog.Logger
	ctx context.Context

	gateway BlockProvider
}

func (e *JobExecutor) Run() {
jobHandelLoop:
	for {
		select {
		case job := <-e.jobs:

			parse := NewParseBlock(e.log, e.ctx, e.gateway)
			parse.execute(job)

			e.free <- true
		case <-e.stop:
			break jobHandelLoop
		}
	}
}

type ParseBlock struct {
	log *slog.Logger
	ctx context.Context

	gateway BlockProvider

	next JobHandler
}

func NewParseBlock(log *slog.Logger, ctx context.Context, gateway BlockProvider) *ParseBlock {
	return &ParseBlock{log: log, ctx: ctx, gateway: gateway}
}

func (p *ParseBlock) execute(job *Job) {
	blk, err := p.gateway.Block(p.ctx, job.blkNumber)
	if err != nil {
		p.log.Error(fmt.Sprintf("#%d: %s", job.blkNumber, err.Error()))
		return
	}

	job.block = blk
}

func (p *ParseBlock) setNext(handler JobHandler) {
	//TODO implement me
	panic("implement me")
}
