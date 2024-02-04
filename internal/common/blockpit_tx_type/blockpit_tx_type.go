package blockpit_tx_type

//go:generate go run github.com/dmarkham/enumer -json -yaml -transform=snake -type=BpTxType

type BpTxType uint

const (
	NoBpTxType BpTxType = iota
	Deposit
	Trade
	Withdrawal
	Mining
	Staking
	GiftSent
	Airdrop
	GiftReceived
	Masternode
	Interest
	Payment
	DerivativeProfit
	DerivativeFee
	DerivativeLoss
	HardFork
	Bounty
	NonTaxableIn
	NonTaxableOut
	Lost
	Cashback
	Fee
	Income
)
