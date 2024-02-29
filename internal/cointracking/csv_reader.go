package cointracking

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
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
	var skippedTxCnt int
	var txIds []string
	err = gocsv.UnmarshalDecoderToCallback(decoder,
		func(tx *common.CointrackingTx) {
			// There are sometimes nonsense transactions that transfer no value
			// and will be rejected by blockpit
			if tx.BuyValue == 0.0 && tx.SellValue == 0.0 && tx.FeeValue == 0.0 {
				skippedTxCnt++
				return
			}

			// finally add a transaction ID
			err = common.SetIdForTransaction(tx)
			if err != nil {
				return
			}

			// Skip transactions that have already been added
			for _, id := range txIds {
				if id == tx.ID {
					return
				}
			}

			txIds = append(txIds, tx.ID)
			txs = append(txs, tx)
		})
	if err != nil {
		if errors.Is(err, common.NoKnownTradeType) {
			return nil, fmt.Errorf("could not get trade types. Maybe your file was exported with an unsupported language")
		}
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
		SkippedTxs:   skippedTxCnt,
		Exchanges:    distinctExchangesFromTransactions(txs),
		Transactions: txs,
	}

	return &fileInfo, nil
}

func distinctExchangesFromTransactions(txs []*common.CointrackingTx) []string {
	exchanges := []string{}
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
