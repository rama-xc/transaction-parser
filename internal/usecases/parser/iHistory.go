package parser

import (
	"context"
	"log/slog"
)

type IHistory interface {
	Start()
	runQueryExecution()
	runWorkersGeneration()
}

type History struct {
	ctx context.Context

	jobs chan *Job
	free chan bool
	stop chan bool

	wrks int

	queue *JobQueue

	gateway BlockProvider
	log     *slog.Logger
}

func (p *History) Start() {

	p.runWorkersGeneration()

	p.runQueryExecution()

}

func (p *History) runQueryExecution() {
	for p.wrks != 0 {
		select {
		case <-p.free:
			p.queue.SendJob()
		case <-p.stop:
			p.wrks--
		}
	}
}

func (p *History) runWorkersGeneration() {
	for i := 0; i < p.wrks; i++ {
		wrk := NewJobExecutor(
			p.jobs,
			p.free,
			p.stop,
			p.log,
			p.ctx,
			p.gateway,
		)

		go wrk.Run()
	}
}
