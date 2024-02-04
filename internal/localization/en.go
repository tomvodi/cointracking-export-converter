package localization

import (
	bp "cointracking-export-converter/internal/common/blockpit_tx_type"
	ct "cointracking-export-converter/internal/common/cointracking_tx_type"
)

var CtTxTypeNames = map[ct.CtTxType]string{
	ct.Trade:                    "Trade",
	ct.MarginTrade:              "Margin Trade",
	ct.DerivativesFuturesTrade:  "Derivatives / Futures Trade",
	ct.Deposit:                  "Deposit",
	ct.Income:                   "Income",
	ct.GiftTip:                  "Gift / Tip",
	ct.RewardBonus:              "Reward / Bonus",
	ct.Mining:                   "Mining",
	ct.Airdrop:                  "Airdrop",
	ct.AirdropNonTaxable:        "Airdrop (non taxable)",
	ct.Staking:                  "Staking",
	ct.Masternode:               "Masternode",
	ct.Minting:                  "Minting",
	ct.MiningCommercial:         "Mining (commercial)",
	ct.DividendsIncome:          "Dividends Income",
	ct.LendingIncome:            "Lending Income",
	ct.InterestIncome:           "Interest Income",
	ct.DerivativesFuturesProfit: "Derivatives / Futures Profit",
	ct.MarginProfit:             "Margin Profit",
	ct.OtherIncome:              "Other Income",
	ct.IncomeNonTaxable:         "Income (non taxable)",
	ct.RemoveLiquidity:          "Remove Liquidity",
	ct.ReceiveLpToken:           "Receive LP Token",
	ct.LpRewards:                "LP Rewards",
	ct.Withdrawal:               "Withdrawal",
	ct.Spend:                    "Spend",
	ct.Donation:                 "Donation",
	ct.Gift:                     "Gift",
	ct.Stolen:                   "Stolen",
	ct.Lost:                     "Lost",
	ct.BorrowingFee:             "Borrowing Fee",
	ct.SettlementFee:            "Settlement Fee",
	ct.MarginLoss:               "Margin Loss",
	ct.MarginFee:                "Margin Fee",
	ct.DerivativesFuturesLoss:   "Derivatives / Futures Loss",
	ct.OtherFee:                 "Other Fee",
	ct.OtherExpense:             "Other Expense",
	ct.ProvideLiquidity:         "Provide Liquidity",
	ct.ReturnLpToken:            "Return LP Token",
	ct.ExpenseNonTaxable:        "Expense (non taxable)",
	ct.SwapNonTaxable:           "Swap (non taxable)",
	ct.ReceiveLoan:              "Receive Loan",
	ct.ReceiveCollateral:        "Receive Collateral",
	ct.SendCollateral:           "Send Collateral",
	ct.RepayLoan:                "Repay Loan",
	ct.Liquidation:              "Liquidation",
}

var BpTxTypeNames = map[bp.BpTxType]string{
	bp.Deposit:          "Deposit",
	bp.Trade:            "Trade",
	bp.Withdrawal:       "Withdrawal",
	bp.Mining:           "Mining",
	bp.Staking:          "Staking",
	bp.GiftSent:         "GiftSent",
	bp.Airdrop:          "Airdrop",
	bp.GiftReceived:     "Gift Received",
	bp.Masternode:       "Masternode",
	bp.Interest:         "Interest",
	bp.Payment:          "Payment",
	bp.DerivativeProfit: "Derivative Profit",
	bp.DerivativeFee:    "Derivative Fee",
	bp.DerivativeLoss:   "Derivative Loss",
	bp.HardFork:         "Hard Fork",
	bp.Bounty:           "Bounty",
	bp.NonTaxableIn:     "Non-Taxable In",
	bp.NonTaxableOut:    "Non-Taxable Out",
	bp.Lost:             "Lost",
	bp.Cashback:         "Cashback",
	bp.Fee:              "Fee",
	bp.Income:           "Income",
}
