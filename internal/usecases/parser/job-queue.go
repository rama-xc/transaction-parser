package parser

import "sync"

type JobQueue struct {
	queue []*Job
	mx    sync.Mutex
}

func NewJobQueue(queue []*Job) *JobQueue {
	return &JobQueue{queue: queue}
}

func (q *JobQueue) push(job *Job) {
	q.queue = append(q.queue, job)
}

func (q *JobQueue) shift() *Job {
	if len(q.queue) == 0 {
		return nil
	}

	s := q.queue[0]
	q.queue = q.queue[1:]

	return s
}

func (q *JobQueue) Length() int {
	return len(q.queue)
}

func (q *JobQueue) GetJob() *Job {
	q.mx.Lock()
	defer q.mx.Unlock()

	return q.shift()
}

func (q *JobQueue) AddJob(job *Job) {
	q.mx.Lock()

	q.push(job)

	q.mx.Unlock()
}
