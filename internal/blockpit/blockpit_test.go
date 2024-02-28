package blockpit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"time"
)

var _ = Describe("Blockpit", func() {
	var err error
	DescribeTable("getBlockpitTxFeeAdapted", func(
		txIn *common.CointrackingTx,
		txExp *common.CointrackingTx,
		wantErr bool,
	) {
		var txOut *common.CointrackingTx
		txOut, err = getBlockpitTxFeeAdapted(txIn)
		if wantErr {
			Expect(err).To(HaveOccurred())
			return
		} else {
			Expect(err).ToNot(HaveOccurred())
		}
		Expect(txOut).To(Equal(txExp))
	},
		Entry("CoinTracking tx has fee in out-currency",
			txWithValues(1.0, "wBTC", 1.1, "BTC", 0.1, "BTC"),
			txWithValues(1.0, "wBTC", 1.0, "BTC", 0.1, "BTC"),
			false,
		),
		Entry("CoinTracking tx has fee in out-currency larger than sell value should not result in negative sell value",
			txWithValues(0.1, "wBTC", 0.1, "BTC", 0.15, "BTC"),
			nil,
			true,
		),
		Entry("CoinTracking tx has fee in in-currency",
			txWithValues(0.9, "wBTC", 1.0, "BTC", 0.1, "wBTC"),
			txWithValues(1.0, "wBTC", 1.0, "BTC", 0.1, "wBTC"),
			false,
		),
		Entry("CoinTracking tx has fee in other currency",
			txWithValues(1.0, "wBTC", 1.0, "BTC", 0.1, "ETH"),
			txWithValues(1.0, "wBTC", 1.0, "BTC", 0.1, "ETH"),
			false,
		),
		Entry("CoinTracking tx has no fee",
			txWithValues(1.0, "wBTC", 1.0, "BTC", 0.0, "BTC"),
			txWithValues(1.0, "wBTC", 1.0, "BTC", 0.0, "BTC"),
			false,
		),
	)
})

func txWithValues(
	buyVal float64, buyCurr string,
	sellVal float64, sellCurr string,
	feeVal float64, feeCurr string) *common.CointrackingTx {
	return &common.CointrackingTx{
		Type:         &common.TxType{TxType: ctt.Trade},
		BuyValue:     buyVal,
		BuyCurrency:  buyCurr,
		SellValue:    sellVal,
		SellCurrency: sellCurr,
		FeeValue:     feeVal,
		FeeCurrency:  feeCurr,
		Exchange:     "Test Exchange",
		Group:        "Test Group",
		Comment:      "Test Comment",
		DateTime:     &common.TxTimestamp{Time: time.Date(2024, 1, 21, 20, 56, 23, 0, time.UTC)},
		ID:           "xxx-yyyy-zzz",
	}

}
