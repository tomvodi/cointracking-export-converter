package app_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
	"github.com/tomvodi/cointracking-export-converter/internal/app"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
)

var _ = Describe("AppContext", func() {
	var appCtx *app.Ctx
	var afs afero.Fs
	var err error

	BeforeEach(func() {
		afs = afero.NewMemMapFs()
		appCtx = app.NewAppContext(afs)
	})

	Describe("Add Export File", func() {
		var fInfo *common.ExportFileInfo

		BeforeEach(func() {
			fInfo = &common.ExportFileInfo{
				FileName:   "file.csv",
				TxCount:    12,
				SkippedTxs: 1,
				Exchanges:  []string{"Binance", "Kraken"},
				FilePath:   "/home/user/file.csv",
				Transactions: []*common.CointrackingTx{
					{
						ID: "tx1-ID",
					},
				},
			}
		})

		JustBeforeEach(func() {
			appCtx.AddExportFile(fInfo)
		})

		It("should add the file to the list of export files", func() {
			Expect(appCtx.ExportFiles()).To(ContainElement(fInfo))
		})

		It("should have added the transaction ids to the list of known transaction ids", func() {
			Expect(appCtx.AllTxIDs()).To(ContainElement("tx1-ID"))
		})
	})

	Describe("Set Last Selected File Dir From File", func() {
		Context("when the file does not exist", func() {
			BeforeEach(func() {
				appCtx.SetLastSelectedFileDirFromFile("/home/user/non-existing-file.csv")
			})

			It("should not set the last selected file dir", func() {
				Expect(appCtx.LastSelectedFileDir()).To(BeEmpty())
			})
		})

		Context("when the file exists", func() {
			BeforeEach(func() {
				err = afs.MkdirAll("/home/user", 0755)
				Expect(err).ToNot(HaveOccurred())
				_, err = afs.Create("/home/user/file.csv")
				Expect(err).ToNot(HaveOccurred())

				appCtx.SetLastSelectedFileDirFromFile("/home/user/file.csv")
			})

			It("should set the last selected file dir", func() {
				Expect(appCtx.LastSelectedFileDir()).To(Equal("/home/user"))
			})
		})
	})

	Describe("Context", func() {
		var ctx context.Context

		Context("when the context is set", func() {
			BeforeEach(func() {
				ctx = context.Background()
				appCtx.SetContext(ctx)
			})

			It("should return the context", func() {
				Expect(appCtx.Context()).To(Equal(ctx))
			})
		})
	})
})
