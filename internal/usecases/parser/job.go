package parser

import "transaction-parser/internal/entity"

type Job struct {
	id        string
	blkNumber int64
	block     *entity.Block
}

func NewJob(id string, blkNumber int64) *Job {
	return &Job{id: id, blkNumber: blkNumber}
}
