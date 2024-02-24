package config

import (
"github.com/tomvodi/cointracking-export-converter/internal/common"
bpt "github.com/tomvodi/cointracking-export-converter/internal/common/blockpit_tx_type"
ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
"github.com/tomvodi/cointracking-export-converter/internal/localization/en"
"fmt"
"github.com/spf13/viper"
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

	err = m.initMappingFromConfig()
	if err != nil {
		return err
	}

	return err
}

func (m *mapper) initMappingFromConfig() error {
	viper.SetDefault("tx_mapping", defaultCt2BpMap)

	config := viper.Get("tx_mapping")

	switch v := config.(type) {
	case map[ctt.CtTxType]bpt.BpTxType:
		m.mapping = v
	case map[string]interface{}:
		configMap, err := typedConfigFromGeneric(v)
		if err != nil {
			return err
		}
		m.mapping = configMap
	default:
		return fmt.Errorf("unhandled type of data in config file: %T", config)
	}
	return nil
}

func typedConfigFromGeneric(genericConfigMap map[string]interface{}) (map[ctt.CtTxType]bpt.BpTxType, error) {
	configMap := map[ctt.CtTxType]bpt.BpTxType{}
	for key, value := range genericConfigMap {
		ctType, err := ctt.CtTxTypeString(key)
		if err != nil {
			return nil, fmt.Errorf("%s is no cointracking transaction type", key)
		}
		if ctType == ctt.NoCtTxType {
			continue
		}
		bptStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("blockpit tx type for cointracking type %s is not a string", key)
		}
		bpType, err := bpt.BpTxTypeString(bptStr)
		if err != nil {
			return nil, fmt.Errorf("%s is no blockpit transaction type", bptStr)
		}
		configMap[ctType] = bpType
	}
	return configMap, nil
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
		if bpType, found = m.mapping[txType]; found {
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

	viper.Set("tx_mapping", m.mapping)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("failed saving mapping to config: %s", err.Error())
	}

	return nil
}

func NewTxTypeManagerInitializer() TxTypeManagerInitializer {
	return &mapper{}
}

var defaultCt2BpMap = map[string]interface{}{
	ctt.NoCtTxType.String():               bpt.NoBpTxType.String(),
	ctt.Trade.String():                    bpt.Trade.String(),
	ctt.MarginTrade.String():              bpt.Trade.String(),
	ctt.DerivativesFuturesTrade.String():  bpt.Trade.String(),
	ctt.Deposit.String():                  bpt.Deposit.String(),
	ctt.Income.String():                   bpt.Income.String(),
	ctt.GiftTip.String():                  bpt.GiftReceived.String(),
	ctt.RewardBonus.String():              bpt.Bounty.String(),
	ctt.Mining.String():                   bpt.Mining.String(),
	ctt.Airdrop.String():                  bpt.Airdrop.String(),
	ctt.AirdropNonTaxable.String():        bpt.NonTaxableIn.String(),
	ctt.Staking.String():                  bpt.Staking.String(),
	ctt.Masternode.String():               bpt.Masternode.String(),
	ctt.Minting.String():                  bpt.Trade.String(),
	ctt.MiningCommercial.String():         bpt.Mining.String(),
	ctt.DividendsIncome.String():          bpt.Income.String(),
	ctt.LendingIncome.String():            bpt.Income.String(),
	ctt.InterestIncome.String():           bpt.Income.String(),
	ctt.DerivativesFuturesProfit.String(): bpt.Income.String(),
	ctt.MarginProfit.String():             bpt.Income.String(),
	ctt.OtherIncome.String():              bpt.Income.String(),
	ctt.IncomeNonTaxable.String():         bpt.NonTaxableIn.String(),
	ctt.RemoveLiquidity.String():          bpt.NonTaxableIn.String(),
	ctt.ReceiveLpToken.String():           bpt.NonTaxableIn.String(),
	ctt.LpRewards.String():                bpt.Income.String(),
	ctt.Withdrawal.String():               bpt.Withdrawal.String(),
	ctt.Spend.String():                    bpt.GiftSent.String(),
	ctt.Donation.String():                 bpt.GiftSent.String(),
	ctt.Gift.String():                     bpt.GiftSent.String(),
	ctt.Stolen.String():                   bpt.Lost.String(),
	ctt.Lost.String():                     bpt.Lost.String(),
	ctt.BorrowingFee.String():             bpt.Fee.String(),
	ctt.SettlementFee.String():            bpt.Fee.String(),
	ctt.MarginLoss.String():               bpt.Lost.String(),
	ctt.MarginFee.String():                bpt.Fee.String(),
	ctt.DerivativesFuturesLoss.String():   bpt.Lost.String(),
	ctt.OtherFee.String():                 bpt.Fee.String(),
	ctt.OtherExpense.String():             bpt.Payment.String(),
	ctt.ProvideLiquidity.String():         bpt.NonTaxableOut.String(),
	ctt.ReturnLpToken.String():            bpt.NonTaxableOut.String(),
	ctt.ExpenseNonTaxable.String():        bpt.NonTaxableOut.String(),
	ctt.SwapNonTaxable.String():           bpt.NonTaxableOut.String(),
	ctt.ReceiveLoan.String():              bpt.NonTaxableIn.String(),
	ctt.ReceiveCollateral.String():        bpt.NonTaxableIn.String(),
	ctt.SendCollateral.String():           bpt.NonTaxableOut.String(),
	ctt.RepayLoan.String():                bpt.Payment.String(),
	ctt.Liquidation.String():              bpt.Trade.String(),
}
