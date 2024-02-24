package config

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	bpt "github.com/tomvodi/cointracking-export-converter/internal/common/blockpit_tx_type"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces/mocks"
)

var _ = Describe("AppConfig", func() {
	var appConf *appConfig
	var err error
	var mockAppCtx *mocks.AppContext
	var mocktxTypeMgr *mocks.TxTypeManager
	var mockWailsLog *mocks.WailsLog

	BeforeEach(func() {
		fs := afero.NewMemMapFs()
		_, err = fs.Create("./test.yaml")
		Expect(err).ToNot(HaveOccurred())
		viper.SetFs(fs)
		viper.SetConfigFile("./test.yaml")
		err = viper.ReadInConfig()
		Expect(err).ToNot(HaveOccurred())

		mockAppCtx = mocks.NewAppContext(GinkgoT())
		mocktxTypeMgr = mocks.NewTxTypeManager(GinkgoT())
		mockWailsLog = mocks.NewWailsLog(GinkgoT())
		appConf = &appConfig{
			appCtx:        mockAppCtx,
			txTypeManager: mocktxTypeMgr,
			wailsLog:      mockWailsLog,
		}
	})

	Describe("Timezone", func() {
		Context("when timezone is not set", func() {
			It("should return empty string", func() {
				Expect(appConf.Timezone()).To(Equal(""))
			})
		})
	})

	Describe("SetTimezone", func() {
		BeforeEach(func() {
			err = appConf.SetTimezone("Europe/Amsterdam")
		})

		It("should not return an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Describe("Getting the timezone", func() {
			It("should return the set timezone", func() {
				Expect(appConf.Timezone()).To(Equal("Europe/Amsterdam"))
			})
		})
	})

	Describe("AllTimezones", func() {
		It("should return all timezones", func() {
			Expect(appConf.AllTimezones()).To(Equal(AllTimezones))
		})
	})

	Describe("BlockpitTxTypes", func() {
		bpTxTypes := []common.TxDisplayName{{Value: "trade", Title: "Trade"}}
		BeforeEach(func() {
			mocktxTypeMgr.EXPECT().BlockpitTxTypes().Return(bpTxTypes, nil)
		})

		It("should return all blockpit tx types", func() {
			Expect(appConf.BlockpitTxTypes()).To(Equal(bpTxTypes))
		})
	})

	Describe("TxTypeMappings", func() {
		txMappings := []common.Ct2BpTxMapping{
			{
				Cointracking: common.TxDisplayName{Value: "bonus", Title: "Bonus"},
				Blockpit:     common.TxDisplayName{Value: "trade", Title: "Trade"},
			},
		}
		BeforeEach(func() {
			mocktxTypeMgr.EXPECT().GetMapping().Return(txMappings, nil)
		})

		It("should return all tx type mappings", func() {
			Expect(appConf.TxTypeMappings()).To(Equal(txMappings))
		})
	})

	Describe("SetCointracking2BlockpitMapping", func() {
		Context("when cointracking tx type is invalid", func() {
			BeforeEach(func() {
				err = appConf.SetCointracking2BlockpitMapping("invalid", "trade")
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when blockpit tx type is invalid", func() {
			BeforeEach(func() {
				err = appConf.SetCointracking2BlockpitMapping("trade", "invalid")
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when cointracking tx type and blockpit tx type are valid", func() {
			BeforeEach(func() {
				mockAppCtx.EXPECT().Context().Return(context.Background())
				mocktxTypeMgr.EXPECT().SetMapping(ctt.Trade, bpt.Trade).Return(nil)
				mockWailsLog.EXPECT().LogTracef(mock.Anything, mock.Anything, mock.Anything, mock.Anything)

				err = appConf.SetCointracking2BlockpitMapping("trade", "trade")
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

})
