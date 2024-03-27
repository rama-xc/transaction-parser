package parser

import (
	"encoding/json"
	"strconv"
	"transaction-parser/internal/entity"
)

type Job struct {
	id        string
	blkNumber int64
	block     *entity.Block
}

func NewJob(id string, blkNumber int64) *Job {
	return &Job{id: id, blkNumber: blkNumber}
}

func (j *Job) toMap() (map[string]string, error) {
	jsonBlock, err := json.Marshal(j.block)
	if err != nil {
		return nil, err
	}

	return map[string]string{"id": j.id, "blkNumber": strconv.FormatInt(j.blkNumber, 10), "block": string(jsonBlock)}, nil
}
