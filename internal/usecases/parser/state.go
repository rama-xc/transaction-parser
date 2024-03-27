package parser

type State interface {
	run()
	//profile() map[string]interface{}
	//options(wrks int)
}

type RunningState struct {
	mgr *ExecutorManager
}

func (s *RunningState) run() {
	s.mgr.log.Warn("Parsing already started!")
}

//func (s *RunningState) options(execs int) {
//	if execs < 0 {
//		return
//	}
//
//	absExecs := int(math.Abs(float64(execs - s.prsr.execs)))
//
//	if execs > s.prsr.execs {
//		s.prsr.createExecs(absExecs)
//		s.prsr.activateExecs(absExecs)
//
//		return
//	}
//
//	s.prsr.destroyExecs(absExecs)
//}

//func (s *RunningState) profile() map[string]interface{} {
//	res := make(map[string]interface{})
//
//	res["state"] = "running"
//	res["executors"] = s.prsr.execs
//	res["queue_len"] = s.prsr.queue.Length()
//
//	return res
//}

type ReadyState struct {
	mgr *ExecutorManager
}

func (s *ReadyState) run() {
	s.mgr.proceedExec()

	s.mgr.setState(s.mgr.runningState)
}

//func (s *ReadyState) options(execs int) {
//	if execs < 0 {
//		return
//	}
//
//	s.prsr.setExecs(execs)
//
//}
//
//func (s *ReadyState) profile() map[string]interface{} {
//	res := make(map[string]interface{})
//
//	res["state"] = "ready to start"
//	res["executors"] = s.prsr.execs
//	res["queue_len"] = s.prsr.queue.Length()
//
//	return res
//}
