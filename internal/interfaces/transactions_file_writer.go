package interfaces

import "github.com/tomvodi/cointracking-export-converter/internal/common"

type TransactionsFileWriter interface {
	WriteTransactionsToFile(filePath string, transactions []*common.CointrackingTx) error
}
