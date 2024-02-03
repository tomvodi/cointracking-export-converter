package common

type ExportFileInfo struct {
	FilePath     string            `json:"filePath"`
	TxCount      int               `json:"txCount"`
	Exchanges    []string          `json:"exchanges"`
	Transactions []*CointrackingTx `json:"-"`
}
