package cointracking

import (
	"cointracking-export-converter/internal/common"
	"cointracking-export-converter/internal/common/cointracking_tx_type"
	"reflect"
	"testing"
	"time"
)

func Test_csvReader_ReadFile(t *testing.T) {
	type args struct {
		filepath string
		loc      *time.Location
		want     []*common.CointrackingTx
		wantErr  bool
	}
	tests := []struct {
		name    string
		prepare func(args *args)
	}{
		{
			name: "test simple english",
			prepare: func(args *args) {
				args.filepath = "./testfiles/file1_en_one_line.csv"
				args.wantErr = false
				args.loc, _ = time.LoadLocation("Europe/Amsterdam")

				args.want = []*common.CointrackingTx{
					{
						Type:         &common.TxType{TxType: cointracking_tx_type.Withdrawal},
						BuyValue:     1111.983574,
						BuyCurrency:  "BTC",
						SellValue:    300.13506100,
						SellCurrency: "USDT",
						FeeValue:     1.456,
						FeeCurrency:  "USDC",
						Exchange:     "ETH Wallet",
						Group:        "testadress",
						Comment:      "this is a comment",
						DateTime: &common.TxTimestamp{
							Time: time.Date(2024, 01, 21, 21, 56, 23, 0, args.loc),
						},
					},
				}
			},
		},
		{
			name: "test simple deutsch",
			prepare: func(args *args) {
				args.filepath = "./testfiles/file1_de_one_line.csv"
				args.wantErr = false
				args.loc, _ = time.LoadLocation("Europe/Amsterdam")

				args.want = []*common.CointrackingTx{
					{
						Type:         &common.TxType{TxType: cointracking_tx_type.Withdrawal},
						BuyValue:     1111.983574,
						BuyCurrency:  "BTC",
						SellValue:    300.13506100,
						SellCurrency: "USDT",
						FeeValue:     1.456,
						FeeCurrency:  "USDC",
						Exchange:     "ETH Wallet",
						Group:        "testadress",
						Comment:      "this is a comment",
						DateTime: &common.TxTimestamp{
							Time: time.Date(2024, 01, 21, 21, 56, 23, 0, args.loc),
						},
					},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &csvReader{}

			a := &args{}

			if tt.prepare != nil {
				tt.prepare(a)
			}

			got, err := c.ReadFile(a.filepath, a.loc)
			if (err != nil) != a.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, a.wantErr)
				return
			}
			if !reflect.DeepEqual(got, a.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, a.want)
			}
		})
	}
}
