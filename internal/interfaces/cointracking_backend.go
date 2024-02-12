package interfaces

//go:generate mockgen -source CoinTrackingBackend.go -destination ./mocks/mock_CointrackingBackend.go

type CoinTrackingBackend interface {
	OpenExportFile(timezone string) (string, error)
	ExportToBlockpitXlsx() error
}
