package trans

type TXInfo struct {
	TxHash     string `json:"tx_hash"`
	TxValue    uint64 `json:"tx_value"`
	TxGas      uint64 `json:"tx_gas"`
	TxGasPrice uint64 `json:"tx_gas_price"`
	TxNonce    uint64 `json:"tx_nonce"`
	TxData     []byte `json:"tx_data"`
	TxTo       string `json:"tx_to"`
	TxFrom     string `json:"tx_from"`
	TxReceipt  uint8  `json:"tx_receipt"`
}
