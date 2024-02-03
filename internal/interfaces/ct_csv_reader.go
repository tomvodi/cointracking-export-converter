package interfaces

import (
	"cointracking-export-converter/internal/common"
	"time"
)

//go:generate mockgen -source CtCsvReader.go -destination ./mocks/mock_CtCsvReader.go.go

type CointractingCsvReader interface {
	ReadFile(filepath string, loc *time.Location) ([]*common.CointrackingTx, error)
}
