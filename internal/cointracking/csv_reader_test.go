package cointracking

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointrackingtxtype"
	"path/filepath"
	"time"
)

var _ = Describe("CsvReader", func() {
	var csvRd *CsvReader
	var err error
	var loc *time.Location
	var fileInfo *common.ExportFileInfo

	BeforeEach(func() {
		csvRd = &CsvReader{}
		loc, err = time.LoadLocation("Europe/Amsterdam")
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("ReadFile", func() {
		Context("when file is not found", func() {
			BeforeEach(func() {
				fileInfo, err = csvRd.ReadFile("non-existing-file.csv", loc, nil)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	DescribeTable("read valid file with one line",
		func(testfilePath string) {
			fileInfo, err = csvRd.ReadFile(testfilePath, loc, nil)
			Expect(err).ToNot(HaveOccurred())

			filename := filepath.Base(testfilePath)
			Expect(fileInfo.FilePath).To(Equal(testfilePath))
			Expect(fileInfo.FileName).To(Equal(filename))
			Expect(fileInfo.TxCount).To(Equal(1))
			Expect(fileInfo.Transactions).To(Equal([]*common.CointrackingTx{
				{
					Type:         &common.TxType{TxType: ctt.Withdrawal},
					BuyValue:     1111.983574,
					BuyCurrency:  "BTC",
					SellValue:    300.13506100,
					SellCurrency: "USDT",
					FeeValue:     1.456,
					FeeCurrency:  "USDC",
					Exchange:     "ETH Wallet",
					Group:        "testaddress",
					Comment:      "this is a comment",
					DateTime: &common.TxTimestamp{
						Time: time.Date(
							2024, 1, 21, 20, 56, 23, 0, loc),
					},
					ID: "12543af52fdbaff4",
				},
			}))
		},
		Entry("German file", "./testfiles/file1_de_one_line.csv"),
		Entry("English file", "./testfiles/file1_en_one_line.csv"),
	)

	DescribeTable("read valid file with one line and comma in field",
		func(testfilePath string) {
			fileInfo, err = csvRd.ReadFile(testfilePath, loc, nil)
			Expect(err).ToNot(HaveOccurred())

			filename := filepath.Base(testfilePath)
			Expect(fileInfo.FilePath).To(Equal(testfilePath))
			Expect(fileInfo.FileName).To(Equal(filename))
			Expect(fileInfo.TxCount).To(Equal(1))
			Expect(fileInfo.Transactions).To(Equal([]*common.CointrackingTx{
				{
					Type:         &common.TxType{TxType: ctt.Withdrawal},
					BuyValue:     1111.983574,
					BuyCurrency:  "BTC",
					SellValue:    300.13506100,
					SellCurrency: "USDT",
					FeeValue:     1.456,
					FeeCurrency:  "USDC",
					Exchange:     "ETH Wallet",
					Group:        "testaddress",
					Comment:      "this, is a comment",
					DateTime: &common.TxTimestamp{
						Time: time.Date(
							2024, 1, 21, 20, 56, 23, 0, loc),
					},
					ID: "69be2b239478213a",
				},
			}))
		},
		Entry("German file and comma in field", "./testfiles/file1_de_one_line_comma_in_field.csv"),
	)

	DescribeTable("read valid file with all transaction types",
		func(testfilePath string) {
			fileInfo, err = csvRd.ReadFile(testfilePath, loc, nil)
			Expect(err).ToNot(HaveOccurred())

			var txTypes []ctt.CtTxType
			for _, tx := range fileInfo.Transactions {
				txTypes = append(txTypes, tx.Type.TxType)
			}
			expectedTxTypes := ctt.CtTxTypeValues()[1:] // remove "NoCtTxType" type
			Expect(txTypes).To(Equal(expectedTxTypes))
		},
		Entry("English file", "./testfiles/CoinTracking_trades_default_en.csv"),
		Entry("German file", "./testfiles/CoinTracking_trades_default_de.csv"),
	)

	Describe("read file with duplicate transactions", func() {
		BeforeEach(func() {
			fileInfo, err = csvRd.ReadFile("./testfiles/file_with_duplicate_tx_de.csv", loc, []string{})
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return transactions without duplicates", func() {
			Expect(fileInfo.TxCount).To(Equal(1))
		})
	})

	Describe("opening the file with existing tx id", func() {
		BeforeEach(func() {
			fileInfo, err = csvRd.ReadFile(
				"./testfiles/file1_de_one_line.csv",
				loc,
				[]string{"12543af52fdbaff4"},
			)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not import the transaction", func() {
			Expect(fileInfo.TxCount).To(Equal(0))
		})
	})
})
