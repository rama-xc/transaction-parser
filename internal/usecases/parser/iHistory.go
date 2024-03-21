package parser

import (
	"context"
	"log/slog"
)

type ProfilingDTO struct {
	Workers     int     `json:"workers"`
	QueueLength int     `json:"queue_length"`
	State       StateID `json:"state"`
	BlockFrom   int64   `json:"block_from"`
	BlockTo     int64   `json:"block_to"`
	BlockNext   int64   `json:"block_next"`
}

type IHistory interface {
	SendCommand(c Command)
	Profile() *ProfilingDTO
	start(resp chan Ping)
	runController()
	createWrks(wrks int)
	destroyWrks(wrks int)
	executeQuery()
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

func (p *History) Profile() *ProfilingDTO {
	return &ProfilingDTO{
		Workers:     p.wrks,
		QueueLength: p.queue.Length(),
		State:       p.currentState.getID(),
		BlockFrom:   0,
		BlockTo:     0,
		BlockNext:   p.queue.ViewNext().blkNumber,
	}
}

func (p *History) start(resp chan Ping) {

	p.currentState.start(resp)

}

func (p *History) runController() {
	for {
		select {
		case <-p.free:
			p.queue.SendJob()
		case command := <-p.comm:
			command.Execute()
		}
	}

	close(p.jobs)
	close(p.free)
	close(p.stop)
	close(p.comm)
}

func (p *History) createWrks(wrks int) {
	for i := 0; i < wrks; i++ {
		wrk := NewJobExecutor(
			p.jobs,
			p.free,
			p.stop,
			p.log,
			p.ctx,
			p.gateway,
		)

		go wrk.Run()

		p.wrks++
	}
}

func (p *History) destroyWrks(wrks int) {
	for i := 0; i < wrks; i++ {
		if p.wrks == 0 {
			break
		}

		p.jobs <- nil

		p.wrks--
	}
}

func (p *History) executeQuery() {
	for i := 0; i < p.wrks; i++ {
		p.queue.SendJob()
	}
}
