package cointrackingtxtype

//go:generate go run github.com/dmarkham/enumer -json -yaml -transform=snake -type=CtTxType

type CtTxType uint

const (
	NoCtTxType CtTxType = iota
	Trade
	MarginTrade
	DerivativesFuturesTrade
	Deposit
	Income
	GiftTip
	RewardBonus
	Mining
	Airdrop
	AirdropNonTaxable
	Staking
	Masternode
	Minting
	MiningCommercial
	DividendsIncome
	LendingIncome
	InterestIncome
	DerivativesFuturesProfit
	MarginProfit
	OtherIncome
	IncomeNonTaxable
	RemoveLiquidity
	ReceiveLpToken
	LpRewards
	Withdrawal
	Spend
	Donation
	Gift
	Stolen
	Lost
	BorrowingFee
	SettlementFee
	MarginLoss
	MarginFee
	DerivativesFuturesLoss
	OtherFee
	OtherExpense
	ProvideLiquidity
	ReturnLpToken
	ExpenseNonTaxable
	SwapNonTaxable
	ReceiveLoan
	ReceiveCollateral
	SendCollateral
	RepayLoan
	Liquidation
)
