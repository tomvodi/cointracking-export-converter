package cointracking_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/tomvodi/cointracking-export-converter/internal/cointracking"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces/mocks"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var _ = Describe("Cointracking", func() {
	var ct *cointracking.Backend
	var wrt *mocks.WailsRuntime
	var ctx *mocks.AppContext
	var csvReader *mocks.CointrackingCsvReader
	var err error

	BeforeEach(func() {
		wrt = mocks.NewWailsRuntime(GinkgoT())
		ctx = mocks.NewAppContext(GinkgoT())
		csvReader = mocks.NewCointrackingCsvReader(GinkgoT())
		ct = cointracking.New(ctx, wrt, csvReader)
	})

	Describe("GetExportFiles", func() {
		var fi *common.ExportFileInfo

		BeforeEach(func() {
			fi = &common.ExportFileInfo{
				FileName:     "testfile.csv",
				TxCount:      12,
				SkippedTxs:   1,
				Exchanges:    []string{"Binance", "Kraken"},
				FilePath:     "/path/to/testfile.csv",
				Transactions: nil,
			}
			ctx.EXPECT().ExportFiles().Return([]*common.ExportFileInfo{fi})
		})

		It("should return the export files", func() {
			var files []*common.ExportFileInfo
			files, err = ct.GetExportFiles()
			Expect(err).To(BeNil())
			Expect(files).To(Equal([]*common.ExportFileInfo{fi}))
		})
	})

	Describe("OpenExportFile", func() {
		var timezone string
		var lastSelectedFileDir string
		var exportFile *common.ExportFileInfo
		var exportFiles []*common.ExportFileInfo
		var filename string

		BeforeEach(func() {
			timezone = ""
			lastSelectedFileDir = ""
			exportFile = nil
			exportFiles = nil
			filename = ""
		})

		JustBeforeEach(func() {
			filename, err = ct.OpenExportFile(timezone)
		})

		When("the file dialog is cancelled", func() {
			BeforeEach(func() {
				ctx.EXPECT().LastSelectedFileDir().Return("/path/to")
				wrt.EXPECT().OpenFileDialog(mock.Anything).
					Return("", fmt.Errorf("cancelled"))
			})

			It("should return an empty string", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		When("the file dialog returns an empty file path", func() {
			BeforeEach(func() {
				ctx.EXPECT().LastSelectedFileDir().Return("/path/to")
				wrt.EXPECT().OpenFileDialog(mock.Anything).
					Return("", nil)
			})

			It("should return an empty string", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("the time location is faulty", func() {
			BeforeEach(func() {
				timezone = "NotATimezone"
				ctx.EXPECT().LastSelectedFileDir().Return("/path/to")
				wrt.EXPECT().OpenFileDialog(mock.Anything).
					Return("/path/to/file", nil)
			})

			It("should return an empty string", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		When("selecting an already selected file", func() {
			BeforeEach(func() {
				exportFiles = []*common.ExportFileInfo{
					{
						FilePath: "/path/to/file",
					},
				}

				ctx.EXPECT().LastSelectedFileDir().Return("/path/to")
				ctx.EXPECT().ExportFiles().Return(exportFiles)
				wrt.EXPECT().OpenFileDialog(mock.Anything).
					Return("/path/to/file", nil)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		When("successfully opening a file", func() {
			BeforeEach(func() {
				timezone = "Europe/Berlin"
				lastSelectedFileDir = "/path/to"
				exportFile = &common.ExportFileInfo{
					FileName: "/path/to/testfile.csv",
					TxCount:  2,
				}
				exportFiles = []*common.ExportFileInfo{
					{
						FileName: "/testfile1",
					},
					{
						FileName: "/testfile2",
					},
				}
				txIDs := []string{"tx1", "tx2"}

				ctx.EXPECT().LastSelectedFileDir().Return(lastSelectedFileDir)
				wrt.EXPECT().OpenFileDialog(runtime.OpenDialogOptions{
					DefaultDirectory: lastSelectedFileDir,
					Title:            "Select CoinTracking export file",
					Filters: []runtime.FileFilter{
						{
							DisplayName: "CoinTracking export files (.csv)",
							Pattern:     "*.csv",
						},
					},
				}).Return("/path/to/testfile.csv", nil)
				ctx.EXPECT().LastSelectedFileDir().Return(lastSelectedFileDir)
				ctx.EXPECT().ExportFiles().Return(exportFiles)
				ctx.EXPECT().AllTxIDs().Return(txIDs)
				ctx.EXPECT().SetLastSelectedFileDirFromFile("/path/to/testfile.csv")
				csvReader.EXPECT().ReadFile(
					"/path/to/testfile.csv",
					mock.Anything,
					txIDs,
				).Return(exportFile, nil)

				ctx.EXPECT().AddExportFile(exportFile)
				wrt.EXPECT().EventsEmit("ExportFilesChanged", exportFiles)
			})

			It("should return the filename", func() {
				Expect(err).To(BeNil())
				Expect(filename).To(Equal("/path/to/testfile.csv"))
			})
		})
	})
})
