package interfaces

import "github.com/tomvodi/cointracking-export-converter/internal/common"

type XmlWriter interface {
	WriteTransactionsToXmlFile(filePath string, transactions []*common.CointrackingTx) error
}
