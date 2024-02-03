package common

import "cointracking-export-converter/internal/common/csv_language"

type ExportFileInfo struct {
	FilePath  string                   `json:"filePath"`
	Language  csv_language.CsvLanguage `json:"language"`
	TxCount   int                      `json:"txCount"`
	Exchanges []string                 `json:"exchanges"`
}
