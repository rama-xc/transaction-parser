package parser

type JobQueue struct {
	queue []*Job
	jobs  chan Job
}
