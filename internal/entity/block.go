package entity

type Block struct {
	Difficulty   uint64         `json:"difficulty"`
	Number       uint64         `json:"number"`
	Time         uint64         `json:"time"`
	Hash         string         `json:"hash"`
	Transactions []*Transaction `json:"transactions"`
}
