package cointracking

import (
	"cointracking-export-converter/internal/common"
	"cointracking-export-converter/internal/interfaces"
	"github.com/gocarina/gocsv"
	"os"
	"time"
)

type csvReader struct {
}

func (c *csvReader) ReadFile(filepath string, loc *time.Location) ([]*common.CointrackingTx, error) {
	exportFile, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
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

	return txs, nil
}

func NewCsvReader() interfaces.CointractingCsvReader {
	return &csvReader{}
}
