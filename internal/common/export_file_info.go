package common

type ExportFileInfo struct {
	FileName     string            `json:"fileName"`
	TxCount      int               `json:"txCount"`
	Exchanges    []string          `json:"exchanges"`
	FilePath     string            `json:"-"`
	Transactions []*CointrackingTx `json:"-"`
}
