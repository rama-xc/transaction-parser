package entity

type Transaction struct {
	Hash             string `json:"hash"`
	To               string `json:"to"`
	From             string `json:"from"`
	TransactionIndex uint16 `json:"transaction_index"`
	Value            string `json:"value"`
	Gas              uint64 `json:"gas"`
	GasPrice         uint64 `json:"gas_price"`
	Nonce            uint64 `json:"nonce"`
}
