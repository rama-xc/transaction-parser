package parser

type ProfilingDTO struct {
	Workers     int     `json:"workers"`
	QueueLength int     `json:"queue_length"`
	State       StateID `json:"state"`
	BlockFrom   int64   `json:"block_from"`
	BlockTo     int64   `json:"block_to"`
	BlockNext   int64   `json:"block_next"`
}

type OptionDTO struct {
	Workers int `json:"workers"`
	Resp    chan Ping
}
