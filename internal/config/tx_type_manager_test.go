package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	bpt "github.com/tomvodi/cointracking-export-converter/internal/common/blockpit_tx_type"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"os"
)

var _ = Describe("TxTypeManager", func() {
	typeMgr := &mapper{}
	var mapping []common.Ct2BpTxMapping
	var err error
	var tempConfig *os.File

	BeforeEach(func() {
		tempConfig, err = os.CreateTemp("", "*.yaml")
		Expect(err).ToNot(HaveOccurred())
		viper.SetConfigFile(tempConfig.Name())
		err = viper.ReadInConfig()
		Expect(err).ToNot(HaveOccurred())

		err := typeMgr.Init()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		err = os.Remove(tempConfig.Name())
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("initial mapping", func() {
		BeforeEach(func() {
			mapping, err = typeMgr.GetMapping()
		})

		It("should have a mapping", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(mapping).ToNot(BeNil())
		})

		DescribeTable("should have correct values for blockpit tx type",
			func(ctTxType ctt.CtTxType, bpTitle string, bpValue string) {
				mapping, err := typeMgr.BlockpitTxType(ctTxType)
				Expect(err).ToNot(HaveOccurred())

				Expect(mapping.Title).To(Equal(bpTitle))
				Expect(mapping.Value).To(Equal(bpValue))
			},
			Entry("should have correct values for trade", ctt.Trade, "Trade", "trade"),
			Entry("should have correct values for marginTrade", ctt.MarginTrade, "Trade", "trade"),
			Entry("should have correct values for derivativesFuturesTrade", ctt.DerivativesFuturesTrade, "Trade", "trade"),
			Entry("should have correct values for deposit", ctt.Deposit, "Deposit", "deposit"),
			Entry("should have correct values for income", ctt.Income, "Income", "income"),
			Entry("should have correct values for giftTip", ctt.GiftTip, "Gift Received", "gift_received"),
			Entry("should have correct values for rewardBonus", ctt.RewardBonus, "Bounty", "bounty"),
			Entry("should have correct values for mining", ctt.Mining, "Mining", "mining"),
			Entry("should have correct values for airdrop", ctt.Airdrop, "Airdrop", "airdrop"),
			Entry("should have correct values for airdropNonTaxable", ctt.AirdropNonTaxable, "Non-Taxable In", "non_taxable_in"),
			Entry("should have correct values for staking", ctt.Staking, "Staking", "staking"),
			Entry("should have correct values for masternode", ctt.Masternode, "Masternode", "masternode"),
			Entry("should have correct values for minting", ctt.Minting, "Trade", "trade"),
			Entry("should have correct values for miningCommercial", ctt.MiningCommercial, "Mining", "mining"),
			Entry("should have correct values for dividendsIncome", ctt.DividendsIncome, "Income", "income"),
			Entry("should have correct values for lendingIncome", ctt.LendingIncome, "Income", "income"),
			Entry("should have correct values for interestIncome", ctt.InterestIncome, "Income", "income"),
			Entry("should have correct values for derivativesFuturesProfit", ctt.DerivativesFuturesProfit, "Income", "income"),
			Entry("should have correct values for marginProfit", ctt.MarginProfit, "Income", "income"),
			Entry("should have correct values for otherIncome", ctt.OtherIncome, "Income", "income"),
			Entry("should have correct values for incomeNonTaxable", ctt.IncomeNonTaxable, "Non-Taxable In", "non_taxable_in"),
			Entry("should have correct values for removeLiquidity", ctt.RemoveLiquidity, "Non-Taxable In", "non_taxable_in"),
			Entry("should have correct values for receiveLpToken", ctt.ReceiveLpToken, "Non-Taxable In", "non_taxable_in"),
			Entry("should have correct values for lpRewards", ctt.LpRewards, "Income", "income"),
			Entry("should have correct values for withdrawal", ctt.Withdrawal, "Withdrawal", "withdrawal"),
			Entry("should have correct values for spend", ctt.Spend, "Gift Sent", "gift_sent"),
			Entry("should have correct values for donation", ctt.Donation, "Gift Sent", "gift_sent"),
			Entry("should have correct values for gift", ctt.Gift, "Gift Sent", "gift_sent"),
			Entry("should have correct values for stolen", ctt.Stolen, "Lost", "lost"),
			Entry("should have correct values for lost", ctt.Lost, "Lost", "lost"),
			Entry("should have correct values for borrowingFee", ctt.BorrowingFee, "Fee", "fee"),
			Entry("should have correct values for settlementFee", ctt.SettlementFee, "Fee", "fee"),
			Entry("should have correct values for marginLoss", ctt.MarginLoss, "Lost", "lost"),
			Entry("should have correct values for marginFee", ctt.MarginFee, "Fee", "fee"),
			Entry("should have correct values for derivativesFuturesLoss", ctt.DerivativesFuturesLoss, "Lost", "lost"),
			Entry("should have correct values for otherFee", ctt.OtherFee, "Fee", "fee"),
			Entry("should have correct values for otherExpense", ctt.OtherExpense, "Payment", "payment"),
			Entry("should have correct values for provideLiquidity", ctt.ProvideLiquidity, "Non-Taxable Out", "non_taxable_out"),
			Entry("should have correct values for returnLpToken", ctt.ReturnLpToken, "Non-Taxable Out", "non_taxable_out"),
			Entry("should have correct values for expenseNonTaxable", ctt.ExpenseNonTaxable, "Non-Taxable Out", "non_taxable_out"),
			Entry("should have correct values for receiveLoan", ctt.ReceiveLoan, "Non-Taxable In", "non_taxable_in"),
			Entry("should have correct values for receiveCollateral", ctt.ReceiveCollateral, "Non-Taxable In", "non_taxable_in"),
			Entry("should have correct values for sendCollateral", ctt.SendCollateral, "Non-Taxable Out", "non_taxable_out"),
			Entry("should have correct values for repayLoan", ctt.RepayLoan, "Payment", "payment"),
			Entry("should have correct values for liquidation", ctt.Liquidation, "Trade", "trade"),
		)

		Describe("changing a mapping", func() {
			BeforeEach(func() {
				err = typeMgr.SetMapping(ctt.Trade, bpt.Airdrop)
			})

			It("should not have an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the new mapping", func() {
				bpTxType, err := typeMgr.BlockpitTxType(ctt.Trade)
				Expect(err).ToNot(HaveOccurred())

				Expect(bpTxType.Title).To(Equal("Airdrop"))
				Expect(bpTxType.Value).To(Equal("airdrop"))
			})
		})
	})
})
