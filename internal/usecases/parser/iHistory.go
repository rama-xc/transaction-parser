package parser

import (
	"context"
	"log/slog"
)

type IHistory interface {
	Start()
	runController()
	generateWrksAndRun()
}

type History struct {
	ctx context.Context

	jobs chan *Job
	free chan bool
	stop chan bool
	comm chan Command

	wrks  int
	queue *JobQueue

	log     *slog.Logger
	gateway BlockProvider
}

func (p *History) Start() {

	p.generateWrksAndRun()

}

func (p *History) runController() {
	for p.wrks != 0 {
		select {
		case <-p.free:
			p.queue.SendJob()
		case <-p.stop:
			p.wrks--
		case command := <-p.comm:
			command.Execute()
		}
	}
}

func (p *History) generateWrksAndRun() {
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
