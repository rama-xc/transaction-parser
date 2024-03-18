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

func NewJobExecutor(jobs chan *Job, free chan bool, stop chan bool, log *slog.Logger, ctx context.Context, gateway BlockProvider) *JobExecutor {
	return &JobExecutor{jobs: jobs, free: free, stop: stop, log: log, ctx: ctx, gateway: gateway}
}

func (e *JobExecutor) Run() {
	for {
		e.free <- true

		job := <-e.jobs

		if job == nil {
			e.stop <- true
			break
		}

		parse := NewParseBlock(e.log, e.ctx, e.gateway)
		parse.execute(job)
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
		p.log.Error(fmt.Sprintf("%s: %s", job.id, err.Error()))
		return
	}

	job.block = blk

	p.log.Info(fmt.Sprintf("%s: block parsed", job.id))
}

func (p *ParseBlock) setNext(handler JobHandler) {
	//TODO implement me
	panic("implement me")
}
