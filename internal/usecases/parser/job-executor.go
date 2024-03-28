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
	mediator Mediator

	free    bool
	alive   bool
	stopped bool

	log *slog.Logger
	ctx context.Context

	blkProvider BlockProvider
	blkCaching  BlockCaching
}

func NewJobExecutor(
	mediator Mediator,
	log *slog.Logger,
	ctx context.Context,
	blkProvider BlockProvider,
	blkCaching BlockCaching,
) *JobExecutor {
	exec := &JobExecutor{
		mediator:    mediator,
		free:        true,
		alive:       true,
		stopped:     true,
		log:         log,
		ctx:         ctx,
		blkProvider: blkProvider,
		blkCaching:  blkCaching,
	}

	go exec.run()

	return exec
}

func (e *JobExecutor) run() {

	for e.alive {

		for !e.stopped {

			job := e.mediator.getJob()

			if job == nil {
				e.stop()
				e.mediator.tryHandleQueueCompletion()

				continue
			}

			e.free = false

			cache := NewCacheBlock(e.log, e.ctx, e.blkCaching)

			parse := NewParseBlock(e.log, e.ctx, e.blkProvider)
			parse.setNext(cache)

			parse.execute(job)

			e.free = true

		}

	}

}

func (e *JobExecutor) kill() {
	e.alive = false
}

func (e *JobExecutor) stop() {
	e.stopped = true
}

func (e *JobExecutor) proceed() {
	e.stopped = false
}

type ParseBlock struct {
	log *slog.Logger
	ctx context.Context

	blkProvider BlockProvider

	next JobHandler
}

func NewParseBlock(log *slog.Logger, ctx context.Context, blkProvider BlockProvider) *ParseBlock {
	return &ParseBlock{log: log, ctx: ctx, blkProvider: blkProvider}
}

func (p *ParseBlock) execute(job *Job) {
	op := "ParseBlock.execute"

	blk, err := p.blkProvider.Block(p.ctx, job.blkNumber)
	if err != nil {
		p.log.Error(fmt.Sprintf("%s-%s: %s", op, job.id, err.Error()))
		return
	}

	job.block = blk

	//p.log.Info(fmt.Sprintf("%s: block parsed", block.id))

	p.next.execute(job)
}

func (p *ParseBlock) setNext(handler JobHandler) {
	p.next = handler
}

type CacheBlock struct {
	log *slog.Logger
	ctx context.Context

	blkCaching BlockCaching

	next JobHandler
}

func (c *CacheBlock) execute(job *Job) {
	op := "CacheBlock.execute"

	err := c.blkCaching.Cache(c.ctx, job.id, job.block)
	if err != nil {
		c.log.Error(fmt.Sprintf("%s-%s: %s", op, job.id, err.Error()))
		return
	}

	c.log.Info(fmt.Sprintf("%s: block cached", job.id))

}

func (c *CacheBlock) setNext(handler JobHandler) {
	c.next = handler

}

func NewCacheBlock(log *slog.Logger, ctx context.Context, blkCaching BlockCaching) *CacheBlock {
	return &CacheBlock{log: log, ctx: ctx, blkCaching: blkCaching}
}
