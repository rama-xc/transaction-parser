package parser

type Mediator interface {
	pingExecFree()
}

type ExecutorManager struct {
	jobs     chan *Job
	jobQueue *JobQueue
	execs    []*JobExecutor
}

func (m *ExecutorManager) createExec() {
	exec := NewJobExecutor()
	m.execs = append(m.execs, exec)
}

func (m *ExecutorManager) destroyExec() {
	if len(m.execs) == 0 {
		return
	}

	exec := m.execs[0]
	m.execs = m.execs[1:]

	exec.kill()
}

func (m *ExecutorManager) runExec() {
	for i := 0; i < len(m.execs); i++ {
		m.jobs <- m.jobQueue.GetJob()
	}
}

func (m *ExecutorManager) pingExecFree() {
	job := m.jobQueue.GetJob()

	m.jobs <- job
}
