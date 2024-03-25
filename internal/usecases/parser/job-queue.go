package parser

type JobQueue struct {
	queue []*Job

	jobs     chan *Job
	complete chan bool
}

func NewJobQueue(queue []*Job, jobs chan *Job) *JobQueue {
	return &JobQueue{queue: queue, jobs: jobs, complete: make(chan bool)}
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

func (q *JobQueue) SendJob() {
	job := q.shift()
	if job == nil {
		q.complete <- true
		return
	}

	q.jobs <- job
}

func (q *JobQueue) AddJob(job *Job) {
	q.push(job)
}
