package interfaces

type CoinTrackingBackend interface {
	OpenExportFile(timezone string) (string, error)
	ExportToBlockpitXlsx() error
}
