package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"os"
)

var _ = Describe("AppConfigFileHandler", func() {
	var configDir = "./test_config_dir"
	var fH *fileHandler
	var err error

	BeforeEach(func() {
		fH = &fileHandler{
			configDir: configDir,
		}
	})

	AfterEach(func() {
		err = os.RemoveAll(configDir)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Init", func() {
		Context("when config file does not exist", func() {
			It("should create the file", func() {
				err = fH.Init()
				Expect(err).ToNot(HaveOccurred())

				_, err = os.Stat(configDir + "/" + configFileName + ".yaml")
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when config file exists", func() {
			Context("and has a config value in it", func() {
				BeforeEach(func() {
					viper.Set("test", "test")
				})

				Context("and the file is read in", func() {
					BeforeEach(func() {
						err = fH.Init()
						Expect(err).ToNot(HaveOccurred())
					})

					It("should read in the file", func() {
						Expect(viper.GetString("test")).To(Equal("test"))
					})
				})
			})
		})
	})
})
