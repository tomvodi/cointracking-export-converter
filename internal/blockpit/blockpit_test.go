package blockpit

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointrackingtxtype"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces/mocks"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

var _ = Describe("Blockpit", func() {
	var err error
	var backend *Backend
	var wrt *mocks.WailsRuntime
	var ctx *mocks.AppContext
	var fileWriter *mocks.TransactionsFileWriter

	Context("ExportToBlockpitXlsx", func() {
		BeforeEach(func() {
			wrt = mocks.NewWailsRuntime(GinkgoT())
			ctx = mocks.NewAppContext(GinkgoT())
			fileWriter = mocks.NewTransactionsFileWriter(GinkgoT())
			backend = New(ctx, wrt, fileWriter)
		})

		JustBeforeEach(func() {
			err = backend.ExportToBlockpitXlsx()
		})

		Context("failed saving file dialog", func() {
			BeforeEach(func() {
				ctx.EXPECT().LastSelectedFileDir().Return("/path/to/dir")
				wrt.EXPECT().SaveFileDialog(runtime.SaveDialogOptions{
					DefaultDirectory: "/path/to/dir",
					DefaultFilename:  "blockpit-import.xlsx",
					Title:            "Save Blockpit manual import file",
					Filters: []runtime.FileFilter{
						{
							DisplayName: "Blockpit import files (.xlsx)",
							Pattern:     "*.xlsx",
						},
					},
				}).Return("", fmt.Errorf("error"))
			})

			It("should return error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("empty file returned", func() {
			BeforeEach(func() {
				ctx.EXPECT().LastSelectedFileDir().Return("/path/to/dir")
				wrt.EXPECT().SaveFileDialog(runtime.SaveDialogOptions{
					DefaultDirectory: "/path/to/dir",
					DefaultFilename:  "blockpit-import.xlsx",
					Title:            "Save Blockpit manual import file",
					Filters: []runtime.FileFilter{
						{
							DisplayName: "Blockpit import files (.xlsx)",
							Pattern:     "*.xlsx",
						},
					},
				}).Return("", nil)
			})

			It("should return nil", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("writing transactions successfully", func() {
			BeforeEach(func() {
				ctx.EXPECT().LastSelectedFileDir().Return("/path/to/dir")
				wrt.EXPECT().SaveFileDialog(runtime.SaveDialogOptions{
					DefaultDirectory: "/path/to/dir",
					DefaultFilename:  "blockpit-import.xlsx",
					Title:            "Save Blockpit manual import file",
					Filters: []runtime.FileFilter{
						{
							DisplayName: "Blockpit import files (.xlsx)",
							Pattern:     "*.xlsx",
						},
					},
				}).Return("/path/to/save/file", nil)
				ctx.EXPECT().ExportFiles().Return([]*common.ExportFileInfo{
					{
						FileName: "file1",
						Transactions: []*common.CointrackingTx{
							{
								Type: &common.TxType{TxType: ctt.Airdrop},
								ID:   "tx1",
							},
						},
					},
					{
						FileName: "file2",
						Transactions: []*common.CointrackingTx{
							{
								Type: &common.TxType{TxType: ctt.Airdrop},
								ID:   "tx2",
							},
						},
					},
				})
				fileWriter.EXPECT().WriteTransactionsToFile(
					"/path/to/save/file",
					[]*common.CointrackingTx{
						{
							Type: &common.TxType{TxType: ctt.Airdrop},
							ID:   "tx1",
						},
						{
							Type: &common.TxType{TxType: ctt.Airdrop},
							ID:   "tx2",
						},
					}).Return(nil)
			})

			It("should return nil", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

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
		}
		Expect(err).ToNot(HaveOccurred())

		Expect(txOut).To(Equal(txExp))
	},
		Entry("CoinTracking tx has fee in out-currency",
			txWithValues(TxValues{
				BuyVal:   1.0,
				BuyCurr:  "wBTC",
				SellVal:  1.1,
				SellCurr: "BTC",
				FeeVal:   0.1,
				FeeCurr:  "BTC",
			}),
			txWithValues(TxValues{
				BuyVal:   1.0,
				BuyCurr:  "wBTC",
				SellVal:  1.0,
				SellCurr: "BTC",
				FeeVal:   0.1,
				FeeCurr:  "BTC",
			}),
			false,
		),
		Entry("CoinTracking tx has fee in out-currency larger than sell value should not result in negative sell value",

			txWithValues(TxValues{
				BuyVal:   0.1,
				BuyCurr:  "wBTC",
				SellVal:  0.1,
				SellCurr: "BTC",
				FeeVal:   0.15,
				FeeCurr:  "BTC",
			}),
			nil,
			true,
		),

		Entry("CoinTracking tx has fee in in-currency",
			txWithValues(TxValues{
				BuyVal:   0.9,
				BuyCurr:  "wBTC",
				SellVal:  1.0,
				SellCurr: "BTC",
				FeeVal:   0.1,
				FeeCurr:  "wBTC",
			}),
			txWithValues(TxValues{
				BuyVal:   1.0,
				BuyCurr:  "wBTC",
				SellVal:  1.0,
				SellCurr: "BTC",
				FeeVal:   0.1,
				FeeCurr:  "wBTC",
			}),
			false,
		),
		Entry("CoinTracking tx has fee in other currency",

			txWithValues(TxValues{
				BuyVal:   1.0,
				BuyCurr:  "wBTC",
				SellVal:  1.0,
				SellCurr: "BTC",
				FeeVal:   0.1,
				FeeCurr:  "ETH",
			}),
			txWithValues(TxValues{
				BuyVal:   1.0,
				BuyCurr:  "wBTC",
				SellVal:  1.0,
				SellCurr: "BTC",
				FeeVal:   0.1,
				FeeCurr:  "ETH",
			}),
			false,
		),
		Entry("CoinTracking tx has no fee",

			txWithValues(TxValues{
				BuyVal:   1.0,
				BuyCurr:  "wBTC",
				SellVal:  1.0,
				SellCurr: "BTC",
				FeeVal:   0.0,
				FeeCurr:  "BTC",
			}),
			txWithValues(TxValues{
				BuyVal:   1.0,
				BuyCurr:  "wBTC",
				SellVal:  1.0,
				SellCurr: "BTC",
				FeeVal:   0.0,
				FeeCurr:  "BTC",
			}),
			false,
		),
	)

	Describe("adaptTxTypeForTradesWith0Income", func() {
		var txIn *common.CointrackingTx

		Context("when trade has 0 sell value", func() {
			BeforeEach(func() {
				txIn = txWithValues(TxValues{
					BuyVal:   0.00001,
					BuyCurr:  "DFI",
					SellVal:  0.0,
					SellCurr: "BTC",
					FeeVal:   0.0,
					FeeCurr:  "",
				})
				Expect(txIn.Type.TxType).To(Equal(ctt.Trade))

				adaptTxTypeForTradesWith0Income(txIn)
			})

			It("should change tx type to other income", func() {
				Expect(txIn.Type.TxType).To(Equal(ctt.OtherIncome))
			})
		})

		Context("when trade has 0 buy value", func() {
			BeforeEach(func() {
				txIn = txWithValues(TxValues{
					BuyVal:   0.0,
					BuyCurr:  "BTC",
					SellVal:  0.000001,
					SellCurr: "DFI",
					FeeVal:   0.0,
					FeeCurr:  "",
				})
				Expect(txIn.Type.TxType).To(Equal(ctt.Trade))

				adaptTxTypeForTradesWith0Income(txIn)
			})

			It("should change tx type to other expense", func() {
				Expect(txIn.Type.TxType).To(Equal(ctt.OtherExpense))
			})
		})
	})
})

type TxValues struct {
	BuyVal   float64
	BuyCurr  string
	SellVal  float64
	SellCurr string
	FeeVal   float64
	FeeCurr  string
}

func txWithValues(
	v TxValues,
) *common.CointrackingTx {
	return &common.CointrackingTx{
		Type:         &common.TxType{TxType: ctt.Trade},
		BuyValue:     v.BuyVal,
		BuyCurrency:  v.BuyCurr,
		SellValue:    v.SellVal,
		SellCurrency: v.SellCurr,
		FeeValue:     v.FeeVal,
		FeeCurrency:  v.FeeCurr,
		Exchange:     "Test Exchange",
		Group:        "Test Group",
		Comment:      "Test Comment",
		DateTime:     &common.TxTimestamp{Time: time.Date(2024, 1, 21, 20, 56, 23, 0, time.UTC)},
		ID:           "xxx-yyyy-zzz",
	}
}
