package blockpit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces/mocks"
	"time"
)

var _ = Describe("XmlWriter convertCtTxToBlockpitTx", func() {
	var writer *xmlWriter
	var err error
	var txIn *common.CointrackingTx
	var txsOut []*common.CointrackingTx
	var txsExpected []*common.CointrackingTx
	var appCfg *mocks.AppConfig

	BeforeEach(func() {
		appCfg = mocks.NewAppConfig(GinkgoT())
		writer = &xmlWriter{
			appCfg: appCfg,
		}
	})

	JustBeforeEach(func() {
		txsOut, err = writer.convertCtTxToBlockpitTx(txIn)
	})

	Context("when transaction is a swap with fee in out-currency", func() {
		BeforeEach(func() {
			txIn = &common.CointrackingTx{
				Type:         &common.TxType{TxType: ctt.SwapNonTaxable},
				BuyValue:     1.1,
				BuyCurrency:  "wBTC",
				SellValue:    1.0,
				SellCurrency: "BTC",
				FeeValue:     0.001,
				FeeCurrency:  "BTC",
				Exchange:     "Binance",
				Group:        "test group",
				Comment:      "just a swap",
				DateTime:     &common.TxTimestamp{Time: time.Date(2024, 1, 21, 20, 56, 23, 0, time.UTC)},
				ID:           "xxx-yyy-zzz",
			}
		})

		Context("when swap handling is 'swap non taxable'", func() {
			BeforeEach(func() {
				appCfg.EXPECT().SwapHandling().Return("swap_non_taxable")
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return two transactions", func() {
				Expect(txsOut).To(HaveLen(2))
			})

			It("should return the correct transactions", func() {
				txExpOut := *txIn
				txExpOut.Type = &common.TxType{TxType: ctt.ExpenseNonTaxable}
				txExpOut.BuyValue = 0
				txExpOut.BuyCurrency = ""
				txExpOut.ID = "659529169254fb80"
				txExpIn := *txIn
				txExpIn.Type = &common.TxType{TxType: ctt.IncomeNonTaxable}
				txExpIn.SellValue = 0
				txExpIn.SellCurrency = ""
				txExpIn.FeeValue = 0
				txExpIn.FeeCurrency = ""
				txExpIn.ID = "c1142b218ba74ed9"

				txsExpected = []*common.CointrackingTx{
					&txExpOut,
					&txExpIn,
				}

				Expect(txsOut).To(Equal(txsExpected))
			})
		})

		Context("when swap handling is 'swap to trade'", func() {
			BeforeEach(func() {
				appCfg.EXPECT().SwapHandling().Return("swap_to_trade")
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return one transaction", func() {
				Expect(txsOut).To(HaveLen(1))
			})

			It("should return the correct transaction", func() {
				txExpOut := *txIn
				txExpOut.Type.TxType = ctt.Trade

				txsExpected = []*common.CointrackingTx{
					&txExpOut,
				}

				Expect(txsOut).To(Equal(txsExpected))
			})
		})
	})
})
