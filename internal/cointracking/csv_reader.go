package cointracking

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"os"
	"path/filepath"
	"slices"
	"time"
)

type CsvReader struct {
}

func (c *CsvReader) ReadFile(
	absoluteFilePath string,
	loc *time.Location,
	existingTxIDs []string,
) (*common.ExportFileInfo, error) {
	exportFile, err := os.OpenFile(absoluteFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer exportFile.Close()

	txs, skippedTxCnt, err := getTransactionsFromFile(exportFile, loc, existingTxIDs)
	if err != nil {
		return nil, err
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

func getTransactionsFromFile(
	exportFile *os.File,
	loc *time.Location,
	existingTxIDs []string,
) ([]*common.CointrackingTx, int, error) {
	decoder := NewCsvDecoder(exportFile)

	var txs []*common.CointrackingTx
	var skippedTxCnt int

	err := gocsv.UnmarshalDecoder(decoder, &txs)
	if err != nil {
		if errors.Is(err, common.ErrNoKnownTradeType) {
			return nil, 0, fmt.Errorf("could not get trade types. Maybe your file was exported with an unsupported language")
		}
		return nil, 0, err
	}

	allTxCnt := len(txs)
	txs = filterOutEmptyTransactions(txs)
	skippedTxCnt = allTxCnt - len(txs)

	txs, err = setTransactionIDs(txs)
	if err != nil {
		return nil, 0, err
	}

	txs = removeDuplicateTxs(txs)
	txs = removeTxsWithIDs(txs, existingTxIDs)

	// set timestamp to correct timezone
	for i := 0; i < len(txs); i++ {
		txs[i].DateTime.Time = txs[i].DateTime.Time.In(loc)
	}

	return txs, skippedTxCnt, nil
}

func filterOutEmptyTransactions(
	txs []*common.CointrackingTx,
) []*common.CointrackingTx {
	var filteredTxs []*common.CointrackingTx

	for _, tx := range txs {
		if tx.BuyValue == 0.0 && tx.SellValue == 0.0 && tx.FeeValue == 0.0 {
			continue
		}
		filteredTxs = append(filteredTxs, tx)
	}

	return filteredTxs
}

func setTransactionIDs(
	txs []*common.CointrackingTx,
) ([]*common.CointrackingTx, error) {
	for _, tx := range txs {
		err := common.SetIDForTransaction(tx)
		if err != nil {
			return nil, err
		}
	}
	return txs, nil
}

func removeDuplicateTxs(
	txs []*common.CointrackingTx,
) []*common.CointrackingTx {
	var uniqueIDs []string
	var filteredTxs []*common.CointrackingTx

	for _, tx := range txs {
		if slices.Contains(uniqueIDs, tx.ID) {
			continue
		}
		uniqueIDs = append(uniqueIDs, tx.ID)
		filteredTxs = append(filteredTxs, tx)
	}

	return filteredTxs
}

func removeTxsWithIDs(
	txs []*common.CointrackingTx,
	txIDs []string,
) []*common.CointrackingTx {
	var filteredTxs []*common.CointrackingTx

	for _, tx := range txs {
		if slices.Contains(txIDs, tx.ID) {
			continue
		}

		filteredTxs = append(filteredTxs, tx)
	}

	return filteredTxs
}

func distinctExchangesFromTransactions(txs []*common.CointrackingTx) []string {
	var exchanges []string
	for _, tx := range txs {
		if tx.Exchange == "" {
			continue
		}

		if !slices.Contains(exchanges, tx.Exchange) {
			exchanges = append(exchanges, tx.Exchange)
		}
	}
	return exchanges
}

func NewCsvReader() *CsvReader {
	return &CsvReader{}
}
