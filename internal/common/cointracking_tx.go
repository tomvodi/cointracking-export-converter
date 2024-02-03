package common

type CointrackingTx struct {
	Type         string      `csv:"type"`
	BuyValue     float64     `csv:"buy_value"`
	BuyCurrency  string      `csv:"buy_currency"`
	SellValue    float64     `csv:"sell"`
	SellCurrency string      `csv:"sell_currency"`
	FeeValue     float64     `csv:"fee"`
	FeeCurrency  string      `csv:"fee_currency"`
	Exchange     string      `csv:"exchange"`
	Group        string      `csv:"group"`
	Comment      string      `csv:"comment"`
	DateTime     TxTimestamp `csv:"date"`
}
