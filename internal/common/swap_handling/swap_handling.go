package swap_handling

//go:generate go run github.com/dmarkham/enumer -json -yaml -transform=snake -type=SwapHandling

type SwapHandling uint

const (
	NoSwapHandling SwapHandling = iota
	SwapNonTaxable
	SwapToTrade
)
