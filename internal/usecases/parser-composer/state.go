package prsrcmps

type State interface {
	run()
	stop()
}

type ReadyState struct {
	parserComposer *ParserComposer
}

func (r *ReadyState) run() {

	//log := r.parserComposer.log
	//
	//cfg := map[parser.Type]int{parser.History: 4}
	//
	//for tp, wrks := range cfg {
	//	psr, err := r.parserComposer.createParser(tp)
	//	if err != nil {
	//		log.Error(
	//			fmt.Sprintf("can't create %s", parser.History),
	//		)
	//		return
	//	}
	//
	//	psr.Start()
	//}

}

func (r *ReadyState) stop() {
	//TODO implement me
	panic("implement me")
}

type RunningState struct {
	parserComposer *ParserComposer
}

func (r *RunningState) run() {
	//TODO implement me
	panic("implement me")
}

func (r *RunningState) stop() {
	//TODO implement me
	panic("implement me")
}

type StoppedState struct {
	parserComposer *ParserComposer
}

func (s *StoppedState) run() {
	//TODO implement me
	panic("implement me")
}

func (s *StoppedState) stop() {
	//TODO implement me
	panic("implement me")
}
