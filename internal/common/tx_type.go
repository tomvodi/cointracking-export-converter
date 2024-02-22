package common

import (
	ctt "cointracking-export-converter/internal/common/cointracking_tx_type"
	"fmt"
)

var trades = []string{"Trade", "Trade"}
var marginTrades = []string{"Margin Trade", "Margin Trade"}
var derivativesFuturesTrades = []string{"Derivatives / Futures Trade", "Derivate / Futures Trade"}
var deposits = []string{"Deposit", "Einzahlung"}
var incomes = []string{"Income", "Einnahme"}
var giftTips = []string{"Gift/Tip", "Geschenk"}
var rewardBonuses = []string{"Reward / Bonus", "Belohnung / Bonus"}
var minings = []string{"Mining", "Mining"}
var airdrops = []string{"Airdrop", "Airdrop"}
var airdropNonTaxables = []string{"Airdrop (non taxable)", "Airdrop (steuerfrei)"}
var stakings = []string{"Staking", "Staking"}
var masternodes = []string{"Masternode", "Masternode"}
var mintings = []string{"Minting", "Minting"}
var miningCommercials = []string{"Mining (commercial)", "Mining (kommerziell)"}
var dividendsIncomes = []string{"Dividends Income", "Dividenden Einnahme"}
var lendingIncomes = []string{"Lending Income", "Lending Einnahme"}
var interestIncomes = []string{"Interest Income", "Zinsen"}
var derivativesFuturesProfits = []string{"Derivatives / Futures Profit", "Derivate / Futures Gewinn"}
var marginProfits = []string{"Margin Profit", "Margin Gewinn"}
var otherIncomes = []string{"Other Income", "Sonstige Einnahme"}
var incomeNonTaxables = []string{"Income (non taxable)", "Einnahme (steuerfrei)"}
var removeLiquiditys = []string{"Remove Liquidity", "Remove Liquidity"}
var receiveLpTokens = []string{"Receive LP Token", "Receive LP Token"}
var lpRewards = []string{"LP Rewards", "LP Rewards"}
var withdrawals = []string{"Withdrawal", "Auszahlung"}
var spends = []string{"Spend", "Ausgabe"}
var donations = []string{"Donation", "Spende"}
var gifts = []string{"Gift", "Schenkung"}
var stolens = []string{"Stolen", "Gestohlen"}
var losts = []string{"Lost", "Verlust"}
var borrowingFees = []string{"Borrowing Fee", "Borrowing Gebühr"}
var settlementFees = []string{"Settlement Fee", "Settlement Gebühr"}
var marginLosses = []string{"Margin Loss", "Margin Verlust"}
var marginFees = []string{"Margin Fee", "Margin Gebühr"}
var derivativesFuturesLosses = []string{"Derivatives / Futures Loss", "Derivate / Futures Verlust"}
var otherFees = []string{"Other Fee", "Sonstige Gebühr"}
var otherExpenses = []string{"Other Expense", "Sonstige Ausgabe"}
var provideLiquidities = []string{"Provide Liquidity", "Provide Liquidity"}
var returnLpTokens = []string{"Return LP Token", "Return LP Token"}
var expenseNonTaxables = []string{"Expense (non taxable)", "Ausgabe (steuerfrei)"}
var swapNonTaxables = []string{"Swap (non taxable)", "Swap (steuerfrei)"}
var receiveLoans = []string{"Receive Loan", "Darlehen erhalten"}
var receiveCollaterals = []string{"Receive Collateral", "Collateral erhalten"}
var sendCollaterals = []string{"Send Collateral", "Collateral gesendet"}
var repayLoans = []string{"Repay Loan", "Darlehen zurückgezahlt"}
var liquidations = []string{"Liquidation", "Liquidation"}

var allLocalizedTradeTypes = map[ctt.CtTxType][]string{
	ctt.Deposit:                  deposits,
	ctt.Withdrawal:               withdrawals,
	ctt.Trade:                    trades,
	ctt.OtherFee:                 otherFees,
	ctt.RewardBonus:              rewardBonuses,
	ctt.Staking:                  stakings,
	ctt.LpRewards:                lpRewards,
	ctt.Mining:                   minings,
	ctt.LendingIncome:            lendingIncomes,
	ctt.BorrowingFee:             borrowingFees,
	ctt.ProvideLiquidity:         provideLiquidities,
	ctt.ExpenseNonTaxable:        expenseNonTaxables,
	ctt.SwapNonTaxable:           swapNonTaxables,
	ctt.Masternode:               masternodes,
	ctt.Income:                   incomes,
	ctt.OtherIncome:              otherIncomes,
	ctt.IncomeNonTaxable:         incomeNonTaxables,
	ctt.ReceiveLoan:              receiveLoans,
	ctt.ReceiveCollateral:        receiveCollaterals,
	ctt.SendCollateral:           sendCollaterals,
	ctt.RepayLoan:                repayLoans,
	ctt.RemoveLiquidity:          removeLiquiditys,
	ctt.InterestIncome:           interestIncomes,
	ctt.ReturnLpToken:            returnLpTokens,
	ctt.Liquidation:              liquidations,
	ctt.GiftTip:                  giftTips,
	ctt.MarginTrade:              marginTrades,
	ctt.Minting:                  mintings,
	ctt.MiningCommercial:         miningCommercials,
	ctt.AirdropNonTaxable:        airdropNonTaxables,
	ctt.Airdrop:                  airdrops,
	ctt.Gift:                     gifts,
	ctt.Stolen:                   stolens,
	ctt.Lost:                     losts,
	ctt.Spend:                    spends,
	ctt.ReceiveLpToken:           receiveLpTokens,
	ctt.OtherExpense:             otherExpenses,
	ctt.DerivativesFuturesTrade:  derivativesFuturesTrades,
	ctt.DividendsIncome:          dividendsIncomes,
	ctt.DerivativesFuturesProfit: derivativesFuturesProfits,
	ctt.MarginProfit:             marginProfits,
	ctt.Donation:                 donations,
	ctt.SettlementFee:            settlementFees,
	ctt.MarginLoss:               marginLosses,
	ctt.MarginFee:                marginFees,
	ctt.DerivativesFuturesLoss:   derivativesFuturesLosses,
}

type TxType struct {
	TxType ctt.CtTxType
}

func (tt *TxType) MarshalCSV() (string, error) {
	return tt.TxType.String(), nil
}

func (tt *TxType) UnmarshalCSV(csv string) error {
	for tradeEnum, tradeTypeTranslations := range allLocalizedTradeTypes {
		for _, translation := range tradeTypeTranslations {
			if translation == csv {
				tt.TxType = tradeEnum
				return nil
			}
		}
	}

	return fmt.Errorf("%w: no trade type found for %s", NoKnownTradeType, csv)
}
