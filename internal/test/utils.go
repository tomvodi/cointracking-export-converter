package test

import (
	"github.com/jaswdr/faker/v2"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"time"
)

func RandomCtTx() *common.CointrackingTx {
	ctTx := &common.CointrackingTx{
		DateTime: &common.TxTimestamp{},
	}
	fake := faker.New()
	fake.Struct().Fill(ctTx)
	txTypeName := fake.RandomStringElement(cointracking_tx_type.CtTxTypeStrings())
	txTypeVal, err := cointracking_tx_type.CtTxTypeString(txTypeName)
	if err != nil {
		panic(err)
	}

	ctTx.Type = &common.TxType{
		TxType: txTypeVal,
	}

	fakeT := faker.Time{
		Faker: &fake,
	}
	ctTx.DateTime.Time = fakeT.Time(time.Now())

	return ctTx
}

func BpTxForCtTx(ctTx *common.CointrackingTx) []*interfaces.BlockpitTx {
	bpTx := &interfaces.BlockpitTx{
		TxType: common.TxDisplayName{},
	}
	fake := faker.New()
	fake.Struct().Fill(bpTx)
	bpTx.CtTx = ctTx

	return []*interfaces.BlockpitTx{bpTx}
}
