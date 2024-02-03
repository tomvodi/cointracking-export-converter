package interfaces

import (
	"cointracking-export-converter/internal/common/cointracking_tx_type"
	"cointracking-export-converter/internal/common/csv_language"
)

//go:generate mockgen -source cointracking_csv_translator.go -destination ./mocks/mock_cointracking_csv_translator.go

type CointrackingCsvTranslator interface {
	GetLanguage() (csv_language.CsvLanguage, error)
	TranslateTxType(txType string) cointracking_tx_type.CtTxType
}
