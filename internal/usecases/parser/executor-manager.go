package parser

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type Mediator interface {
	getJob() *Job
	tryHandleQueueCompletion()
}

type ExecutorManager struct {
	jobQueue *JobQueue
	execs    []*JobExecutor

	currentState State
	readyState   State
	runningState State

	log     *slog.Logger
	gateway BlockProvider
	redis   *redis.Client
}

func NewExecutorManager(queue []*Job, execsAmount int, log *slog.Logger, gateway BlockProvider, redis *redis.Client) *ExecutorManager {
	e := &ExecutorManager{
		jobQueue: NewJobQueue(queue),
		execs:    []*JobExecutor{},
		log:      log,
		gateway:  gateway,
		redis:    redis,
	}

	readyState := &ReadyState{mgr: e}
	runningState := &RunningState{mgr: e}

	e.setState(readyState)

	e.readyState = readyState
	e.runningState = runningState

	for i := 0; i < execsAmount; i++ {
		e.createExec()
	}

	return e
}

func (m *ExecutorManager) setState(s State) {
	m.currentState = s
}

func (m *ExecutorManager) isQueueDone() bool {
	f := 0

	for _, exec := range m.execs {
		if exec.free {
			f++
		}
	}

	isAllExecsFree := len(m.execs) == f
	isQueueEmpty := m.jobQueue.Length() == 0

	return isAllExecsFree && isQueueEmpty
}

func (m *ExecutorManager) createExec() {
	exec := NewJobExecutor(
		m,
		m.log,
		context.TODO(),
		m.gateway,
		m.redis,
	)

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

func (m *ExecutorManager) proceedExec() {

	for _, exec := range m.execs {
		exec.proceed()
	}

}

func (m *ExecutorManager) RunParsing() {
	m.currentState.run()
}

func (m *ExecutorManager) getJob() *Job {

	return m.jobQueue.GetJob()

}

func (m *ExecutorManager) tryHandleQueueCompletion() {
	if !m.isQueueDone() {
		return
	}
	m.setState(m.readyState)
}
