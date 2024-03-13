package entity

type SegmentMeta struct {
	SegmentID string `json:"segmentID"`
	From      int64  `json:"from"`
	To        int64  `json:"to"`
}

type Segment struct {
	Blocks []*Block     `json:"blocks"`
	Meta   *SegmentMeta `json:"meta"`
}
