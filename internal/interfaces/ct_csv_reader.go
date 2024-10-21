package interfaces

import (
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"time"
)

type CointrackingCsvReader interface {
	ReadFile(
		filepath string,
		loc *time.Location,
		existingTxIDs []string,
	) (*common.ExportFileInfo, error)
}
