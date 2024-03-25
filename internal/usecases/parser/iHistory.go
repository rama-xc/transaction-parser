package parser

import (
	"context"
	"log/slog"
)

type IHistory interface {
	SendCommand(c Command)
	Profile() map[string]interface{}

	run()
	options(execs int)

	runController()

	setExecs(execs int)
	createExecs(execs int)
	destroyExecs(execs int)
	activateExecs(execs int)
}

type History struct {
	ctx context.Context

	jobs     chan *Job
	execFree chan bool
	execStop chan bool
	comm     chan Command

	execs int
	queue *JobQueue

	log     *slog.Logger
	gateway BlockProvider

	readyState   IState
	runningState IState

	currentState IState
}

func (p *History) SendCommand(c Command) {
	p.comm <- c
}

func (p *History) setState(s IState) {
	p.currentState = s
}

func (p *History) Profile() map[string]interface{} {
	return p.currentState.profile()
}

func (p *History) run() {

	p.currentState.run()

}

func (p *History) options(execs int) {

	p.currentState.options(execs)

}

func (p *History) runController() {
	for {
		select {
		case <-p.execFree:
			p.queue.SendJob()
		case command := <-p.comm:
			command.Execute()
		}
	}
}

func (p *History) createExecs(execs int) {
	for i := 0; i < execs; i++ {
		exec := NewJobExecutor(
			p.jobs,
			p.execFree,
			p.execStop,
			p.log,
			p.ctx,
			p.gateway,
		)

		go exec.Run()

		p.execs++
	}
}

func (p *History) destroyExecs(execs int) {
	for i := 0; i < execs; i++ {
		if p.execs == 0 {
			break
		}

		p.execStop <- true

		p.execs--
	}
}

func (p *History) setExecs(execs int) {

	p.destroyExecs(p.execs)
	p.createExecs(execs)

}

func (p *History) activateExecs(execs int) {
	for i := 0; i < execs; i++ {
		p.queue.SendJob()
	}
}

func (p *History) handleQueryCompletion() {
	for i := 0; i < p.execs; i++ {
		<-p.queue.complete
	}

	p.log.Info("Parsing Completed.")

	p.setState(
		p.readyState,
	)
}
