package blockpit_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/tomvodi/cointracking-export-converter/internal/blockpit"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointrackingtxtype"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces/mocks"
	"github.com/tomvodi/cointracking-export-converter/internal/test"
	"time"
)

var _ = Describe("XmlWriter FromCointrackingTx", func() {
	var convert *blockpit.TxConverter
	var err error
	var txIn *common.CointrackingTx
	var txsOut []*interfaces.BlockpitTx
	var txsExpected []*interfaces.BlockpitTx
	var appCfg *mocks.AppConfig
	var txTypeMgr *mocks.TxTypeManager

	BeforeEach(func() {
		appCfg = mocks.NewAppConfig(GinkgoT())
		txTypeMgr = mocks.NewTxTypeManager(GinkgoT())
		convert = blockpit.NewTxConvert(
			appCfg,
			txTypeMgr,
		)
	})

	JustBeforeEach(func() {
		txsOut, err = convert.FromCointrackingTx(txIn)
	})

	Context("When there is no display type for CoinTracking Tx type", func() {
		BeforeEach(func() {
			txIn = test.RandomCtTx()
			txTypeMgr.EXPECT().BlockpitTxType(mock.Anything).
				Return(common.TxDisplayName{}, fmt.Errorf("no matching type"))
		})

		It("should return an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when transaction is a swap with fee in out-currency", func() {
		BeforeEach(func() {
			txTypeMgr.EXPECT().BlockpitTxType(mock.Anything).
				Return(common.TxDisplayName{
					Title: "Tx Type Display Name",
					Value: "type value",
				}, nil)

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

		Context("when there is no swap handling enum for the given string", func() {
			BeforeEach(func() {
				appCfg.EXPECT().SwapHandling().Return("xxx_not_existing_xxx")
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
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
				expCtTxOut := *txIn
				expCtTxOut.Type = &common.TxType{TxType: ctt.ExpenseNonTaxable}
				expCtTxOut.BuyValue = 0
				expCtTxOut.BuyCurrency = ""
				expCtTxOut.ID = "659529169254fb80"
				txExpOut := &interfaces.BlockpitTx{
					TxType: common.TxDisplayName{
						Title: "Tx Type Display Name",
						Value: "type value",
					},
					CtTx: &expCtTxOut,
				}

				expCtTxIn := *txIn
				expCtTxIn.Type = &common.TxType{TxType: ctt.IncomeNonTaxable}
				expCtTxIn.SellValue = 0
				expCtTxIn.SellCurrency = ""
				expCtTxIn.FeeValue = 0
				expCtTxIn.FeeCurrency = ""
				expCtTxIn.ID = "c1142b218ba74ed9"
				txExpIn := &interfaces.BlockpitTx{
					TxType: common.TxDisplayName{
						Title: "Tx Type Display Name",
						Value: "type value",
					},
					CtTx: &expCtTxIn,
				}

				txsExpected = []*interfaces.BlockpitTx{
					txExpOut,
					txExpIn,
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

				txOut := &interfaces.BlockpitTx{
					TxType: common.TxDisplayName{
						Title: "Tx Type Display Name",
						Value: "type value",
					},
					CtTx: &txExpOut,
				}

				txsExpected = []*interfaces.BlockpitTx{
					txOut,
				}

				Expect(txsOut).To(Equal(txsExpected))
			})
		})
	})
})
