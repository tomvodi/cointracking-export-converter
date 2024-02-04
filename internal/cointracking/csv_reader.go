package cointracking

import (
	"cointracking-export-converter/internal/common"
	"cointracking-export-converter/internal/interfaces"
	"github.com/gocarina/gocsv"
	"os"
	"path/filepath"
	"time"
)

type csvReader struct {
}

func (c *csvReader) ReadFile(absoluteFilePath string, loc *time.Location) (*common.ExportFileInfo, error) {
	exportFile, err := os.OpenFile(absoluteFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer exportFile.Close()

	decoder := NewCsvDecoder(exportFile)

	var txs []*common.CointrackingTx
	err = gocsv.UnmarshalDecoderToCallback(decoder,
		func(tx *common.CointrackingTx) {
			txs = append(txs, tx)
		})
	if err != nil {
		return nil, err
	}

	// set timestamp to correct timezone
	for i := 0; i < len(txs); i++ {
		txs[i].DateTime.Time = txs[i].DateTime.Time.In(loc)
	}

	filename := filepath.Base(absoluteFilePath)

	fileInfo := common.ExportFileInfo{
		FilePath:     absoluteFilePath,
		FileName:     filename,
		TxCount:      len(txs),
		Exchanges:    distinctExchangesFromTransactions(txs),
		Transactions: txs,
	}

	return &fileInfo, nil
}

func distinctExchangesFromTransactions(txs []*common.CointrackingTx) []string {
	var exchanges []string
	for _, tx := range txs {
		if tx.Exchange == "" {
			continue
		}

		exchangeAlreadyAdded := false
		for _, exchange := range exchanges {
			if exchange == tx.Exchange {
				exchangeAlreadyAdded = true
				break
			}
		}
		if !exchangeAlreadyAdded {
			exchanges = append(exchanges, tx.Exchange)
		}
	}
	return exchanges
}

func NewCsvReader() interfaces.CointrackingCsvReader {
	return &csvReader{}
}
