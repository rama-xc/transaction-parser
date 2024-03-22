package parser

type StateID string

var (
	ReadyStateID   StateID = "ready"
	RunningStateID StateID = "running"
)

type IState interface {
	getID() StateID

	start(resp chan Ping)
	options(dto OptionDTO)
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

func (s *RunningState) options(dto OptionDTO) {
	if dto.Workers > s.prsr.wrks {
		s.prsr.createWrks(dto.Workers - s.prsr.wrks)
	} else {

	}

	s.prsr.setWrks(dto.Workers)

	dto.Resp <- SuccessOptionPing
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

func (s *ReadyState) options(dto OptionDTO) {
	s.prsr.setWrks(dto.Workers)

	dto.Resp <- SuccessOptionPing
}
