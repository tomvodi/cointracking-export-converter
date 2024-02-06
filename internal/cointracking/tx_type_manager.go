package cointracking

import (
	"cointracking-export-converter/internal/common"
	bpt "cointracking-export-converter/internal/common/blockpit_tx_type"
	ctt "cointracking-export-converter/internal/common/cointracking_tx_type"
	"cointracking-export-converter/internal/interfaces"
	"cointracking-export-converter/internal/localization/en"
	"fmt"
)

type TxTypeManagerInitializer interface {
	interfaces.TxTypeManager
	interfaces.Initializer
}

type mapper struct {
	ctDisplays map[ctt.CtTxType]common.TxDisplayName
	bpDisplays map[bpt.BpTxType]common.TxDisplayName
	mapping    map[ctt.CtTxType]bpt.BpTxType
}

func (m *mapper) BlockpitTxType(ctTxType ctt.CtTxType) (common.TxDisplayName, error) {
	bpType, found := m.mapping[ctTxType]
	if !found {
		return common.TxDisplayName{}, fmt.Errorf("no blockpit tx type for cointracking type '%s'", ctTxType.String())
	}

	bpDisplay, found := m.bpDisplays[bpType]
	if !found {
		return common.TxDisplayName{}, fmt.Errorf("no blockpit display type for '%s'", bpType.String())
	}

	return bpDisplay, nil
}

func (m *mapper) Init() (err error) {
	m.bpDisplays, err = initBlockpitDisplaysLocalized("english", en.BpTxTypeNames)
	if err != nil {
		return err
	}
	m.ctDisplays, err = initCointrackingDisplaysLocalized("english", en.CtTxTypeNames)
	if err != nil {
		return err
	}
	m.mapping = defaultCt2BpMap

	return nil
}

func initBlockpitDisplaysLocalized(
	languageName string,
	typeNames map[bpt.BpTxType]string,
) (map[bpt.BpTxType]common.TxDisplayName, error) {
	txNames := map[bpt.BpTxType]common.TxDisplayName{}
	for _, txType := range bpt.BpTxTypeValues() {
		if txType == bpt.NoBpTxType {
			continue
		}

		txName := common.TxDisplayName{
			Value: txType.String(),
		}

		if translation, found := typeNames[txType]; found {
			txName.Title = translation
		} else {
			return nil, fmt.Errorf(
				"no localization for Blockpit tx type"+
					" %s in language %s found", txName.Value, languageName)
		}
		txNames[txType] = txName
	}

	return txNames, nil
}

func initCointrackingDisplaysLocalized(
	languageName string,
	typeNames map[ctt.CtTxType]string,
) (map[ctt.CtTxType]common.TxDisplayName, error) {
	txNames := map[ctt.CtTxType]common.TxDisplayName{}
	for _, txType := range ctt.CtTxTypeValues() {
		if txType == ctt.NoCtTxType {
			continue
		}

		txName := common.TxDisplayName{
			Value: txType.String(),
		}

		if translation, found := typeNames[txType]; found {
			txName.Title = translation
		} else {
			return nil, fmt.Errorf(
				"no localization for Cointracking tx type"+
					" %s in language %s found", txName.Value, languageName)
		}
		txNames[txType] = txName
	}

	return txNames, nil
}

func (m *mapper) BlockpitTxTypes() (txNames []common.TxDisplayName, err error) {
	for _, txType := range bpt.BpTxTypeValues() {
		if txType == bpt.NoBpTxType {
			continue
		}
		txNames = append(txNames, m.bpDisplays[txType])
	}
	return txNames, nil
}

func (m *mapper) GetMapping() (mapping []common.Ct2BpTxMapping, err error) {
	for _, txType := range ctt.CtTxTypeValues() {
		if txType == ctt.NoCtTxType {
			continue
		}

		mapItem := common.Ct2BpTxMapping{
			Cointracking: common.TxDisplayName{
				Value: txType.String(),
			},
		}

		if translation, found := en.CtTxTypeNames[txType]; found {
			mapItem.Cointracking.Title = translation
		} else {
			return nil,
				fmt.Errorf("no localization for CoinTracking tx type %s found", txType.String())
		}

		var bpType bpt.BpTxType
		var found bool
		if bpType, found = defaultCt2BpMap[txType]; found {
			mapItem.Blockpit.Value = bpType.String()
		} else {
			return nil,
				fmt.Errorf("no blockpit tx type for CoinTracking tx type %s found", txType.String())
		}

		if translation, found := en.BpTxTypeNames[bpType]; found {
			mapItem.Blockpit.Title = translation
		} else {
			return nil,
				fmt.Errorf("no localization for Blockpit tx type %s found", mapItem.Blockpit.Value)
		}

		mapping = append(mapping, mapItem)
	}

	return mapping, nil
}

func (m *mapper) SetMapping(ctTxType ctt.CtTxType, bpTxType bpt.BpTxType) error {
	m.mapping[ctTxType] = bpTxType

	return fmt.Errorf("set mapping from ct %v to bp %v", ctTxType, bpTxType)
}

func NewTxTypeManagerInitializer() TxTypeManagerInitializer {
	return &mapper{}
}

var defaultCt2BpMap = map[ctt.CtTxType]bpt.BpTxType{
	ctt.NoCtTxType:               bpt.NoBpTxType,
	ctt.Trade:                    bpt.Trade,
	ctt.MarginTrade:              bpt.Trade,
	ctt.DerivativesFuturesTrade:  bpt.Trade,
	ctt.Deposit:                  bpt.Deposit,
	ctt.Income:                   bpt.Income,
	ctt.GiftTip:                  bpt.GiftReceived,
	ctt.RewardBonus:              bpt.Bounty,
	ctt.Mining:                   bpt.Mining,
	ctt.Airdrop:                  bpt.Airdrop,
	ctt.AirdropNonTaxable:        bpt.NonTaxableIn,
	ctt.Staking:                  bpt.Staking,
	ctt.Masternode:               bpt.Masternode,
	ctt.Minting:                  bpt.Trade,
	ctt.MiningCommercial:         bpt.Mining,
	ctt.DividendsIncome:          bpt.Income,
	ctt.LendingIncome:            bpt.Income,
	ctt.InterestIncome:           bpt.Income,
	ctt.DerivativesFuturesProfit: bpt.Income,
	ctt.MarginProfit:             bpt.Income,
	ctt.OtherIncome:              bpt.Income,
	ctt.IncomeNonTaxable:         bpt.NonTaxableIn,
	ctt.RemoveLiquidity:          bpt.NonTaxableIn,
	ctt.ReceiveLpToken:           bpt.NonTaxableIn,
	ctt.LpRewards:                bpt.Income,
	ctt.Withdrawal:               bpt.Withdrawal,
	ctt.Spend:                    bpt.GiftSent,
	ctt.Donation:                 bpt.GiftSent,
	ctt.Gift:                     bpt.GiftSent,
	ctt.Stolen:                   bpt.Lost,
	ctt.Lost:                     bpt.Lost,
	ctt.BorrowingFee:             bpt.Fee,
	ctt.SettlementFee:            bpt.Fee,
	ctt.MarginLoss:               bpt.Lost,
	ctt.MarginFee:                bpt.Fee,
	ctt.DerivativesFuturesLoss:   bpt.Lost,
	ctt.OtherFee:                 bpt.Fee,
	ctt.OtherExpense:             bpt.Payment,
	ctt.ProvideLiquidity:         bpt.NonTaxableOut,
	ctt.ReturnLpToken:            bpt.NonTaxableOut,
	ctt.ExpenseNonTaxable:        bpt.NonTaxableOut,
	ctt.SwapNonTaxable:           bpt.NonTaxableOut,
	ctt.ReceiveLoan:              bpt.NonTaxableIn,
	ctt.ReceiveCollateral:        bpt.NonTaxableIn,
	ctt.SendCollateral:           bpt.NonTaxableOut,
	ctt.RepayLoan:                bpt.Payment,
	ctt.Liquidation:              bpt.Trade,
}
