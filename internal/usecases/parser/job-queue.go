package parser

type JobQueue struct {
	queue []*Job

	jobs chan *Job
	free chan bool
	stop chan bool
}

func NewJobQueue(queue []*Job, jobs chan *Job, free chan bool, stop chan bool) *JobQueue {
	return &JobQueue{queue: queue, jobs: jobs, free: free, stop: stop}
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

func (q *JobQueue) SendJob() {
	q.jobs <- q.shift()
}

func (q *JobQueue) AddJob(job *Job) {
	q.push(job)
}
