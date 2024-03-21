package parser

type StateID string

var (
	ReadyStateID   StateID = "ready"
	RunningStateID StateID = "running"
)

type IState interface {
	getID() StateID
	start(resp chan Ping)
}

type State struct {
	id StateID
}

func (s *State) getID() StateID {
	return s.id
}

type RunningState struct {
	State
	prsr *History
}

func (s *RunningState) start(resp chan Ping) {
	resp <- AlreadyStartedPing
}

type ReadyState struct {
	State
	prsr *History
}

func (s *ReadyState) start(resp chan Ping) {
	runningState := s.prsr.runningState

	s.prsr.executeQuery()

	s.prsr.setState(runningState)

	resp <- SuccessStartedPing
}
