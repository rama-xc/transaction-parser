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

func (p *History) setState(s IState) {
	p.currentState = s
}

func (p *History) SendCommand(c Command) {
	p.comm <- c
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
		command := <-p.comm

		command.Execute()
	}
}
