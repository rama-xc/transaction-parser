package parser

import (
	"log/slog"
)

type IHistory interface {
	SendCommand(c Command)

	runController()
	run()
}

type History struct {
	comm        chan Command
	execManager *ExecutorManager
	log         *slog.Logger
}

func (p *History) SendCommand(c Command) {
	p.comm <- c
}

func (p *History) run() {
	p.execManager.RunParsing()
}

func (p *History) runController() {
	for {
		command := <-p.comm

		command.Execute()
	}
}
