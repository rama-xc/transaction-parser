package parser

import (
	"context"
	"log/slog"
	"transaction-parser/internal/entity"
)

type BlockProvider interface {
	Block(ctx context.Context, number int64) (*entity.Block, error)
	LastBlockNum(ctx context.Context) (int64, error)
}

type Type string

const (
	History  Type = "history"
	Support  Type = "support"
	RealTime Type = "real-time"
)

type IParser interface {
	Start()
	runQueryExecution()
	runWorkersGeneration()
}

type Parser struct {
	ctx context.Context

	jobs chan *Job
	free chan bool
	stop chan bool

	wrks int

	queue *JobQueue

	gateway BlockProvider
	log     *slog.Logger
}

func (p *Parser) Start() {

	p.runWorkersGeneration()

	p.runQueryExecution()

}

func (p *Parser) runQueryExecution() {
	for p.wrks != 0 {
		select {
		case <-p.free:
			p.queue.SendJob()
		case <-p.stop:
			p.wrks--
		}
	}
}

func (p *Parser) runWorkersGeneration() {
	for i := 0; i < p.wrks; i++ {
		wrk := NewJobExecutor(
			p.jobs,
			p.free,
			p.stop,
			p.log,
			p.ctx,
			p.gateway,
		)

		go wrk.Run()
	}
}

type HistoryParser struct {
	Parser
}

func NewHistoryParser(parser Parser) *HistoryParser {
	return &HistoryParser{Parser: parser}
}

//func NewHistoryParser(
//	gateway BlockProvider, log *slog.Logger,
//	wrks int,
//	fromBlk, toBlk int64,
//) *HistoryParser {
//	jobs := make(chan *Job)
//	free := make(chan bool)
//	stop := make(chan bool)
//
//	var queue []*Job
//
//	for blk := fromBlk; blk <= toBlk; blk++ {
//		job := NewJob(fmt.Sprintf("#%d", blk), blk)
//		queue = append(queue, job)
//	}
//
//	jq := NewJobQueue(
//		queue,
//		jobs, free, stop,
//	)
//
//	return &HistoryParser{
//		Parser: Parser{
//			jobs:    jobs,
//			free:    free,
//			stop:    stop,
//			wrks:    wrks,
//			log:     log,
//			ctx:     context.Background(),
//			queue:   jq,
//			gateway: gateway,
//		},
//	}
//
//}

type RealTimeParser struct {
	Parser
}

func NewRealTimeParser(
	gateway BlockProvider,
) IParser {

	return &RealTimeParser{
		Parser: Parser{
			gateway: gateway,
		},
	}

}

type SupportParser struct {
	Parser
}

//func NewSupportParser(
//	wrks int,
//	gateway prschain.BlockProvider,
//) *SupportParser {
//
//	return &SupportParser{
//		Parser: Parser{
//			wrks:    wrks,
//			gateway: gateway,
//		},
//	}
//
//}
