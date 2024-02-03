package csv_language

//go:generate go run github.com/dmarkham/enumer -json -yaml -transform=lower -type=CsvLanguage

type CsvLanguage uint

const (
	En CsvLanguage = iota
	De
)
