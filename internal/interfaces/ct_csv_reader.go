package interfaces

import (
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"time"
)

//go:generate mockgen -source CtCsvReader.go -destination ./mocks/mock_CtCsvReader.go.go

type CointrackingCsvReader interface {
	ReadFile(filepath string, loc *time.Location) (*common.ExportFileInfo, error)
}
