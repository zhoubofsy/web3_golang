package block

type BlockInfo struct {
	Hash       string `json:"hash"`
	Height     uint64 `json:"height"`
	Timestamp  uint64 `json:"timestamp"`
	Difficulty uint64 `json:"difficulty"`
	Nonce      uint64 `json:"nonce"`
	Miner      string `json:"miner"`
	TransCount uint64 `json:"trans_count"`
}
